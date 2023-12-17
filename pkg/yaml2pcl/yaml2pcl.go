// package yaml2pcl provides a method to convert k8s yaml to
// PCL (pulumi schema)
package yaml2pcl

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/goccy/go-yaml/ast"
	"github.com/goccy/go-yaml/parser"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// ConvertFile returns a string conversion of the input YAML
// in a file into PCL: sample below
// Output: resource foo "kubernetes:core/v1:Namespace" {
// apiVersion = "v1"
// kind = "Namespace"
// metadata = {
// name = "foo"
// }
// }
func ConvertFile(filename string) (string, hcl.Diagnostics, error) {
	testFiles, err := parser.ParseFile(filename, parser.ParseComments)
	if err != nil {
		return "", hcl.Diagnostics{}, fmt.Errorf("failed to parse: input file does not contain valid yaml: %s", filename)
	}
	return convert(*testFiles)
}

// ConvertDirectory returns a string conversion of the input directory with
// YAML manifests into PCL: sample below
// Output: resource foo "kubernetes:core/v1:Namespace" {
// apiVersion = "v1"
// kind = "Namespace"
// metadata = {
// name = "foo"
// }
// }
func ConvertDirectory(dirName string) (string, hcl.Diagnostics, error) {
	var buff bytes.Buffer
	diagnostics := hcl.Diagnostics{}

	files, err := os.ReadDir(dirName)
	if err != nil {
		return "", diagnostics, err
	}

	yamlFiles := 0
	for _, file := range files {
		if strings.Contains(file.Name(), ".yaml") || strings.Contains(file.Name(), ".yml") {
			yamlFiles++
			pcl, diags, err := ConvertFile(filepath.Join(dirName, file.Name()))
			// append all diags from file into new diagnostics object
			diagnostics = diagnostics.Extend(diags)
			if err != nil {
				return "", diagnostics, err
			}
			_, err = fmt.Fprintf(&buff, "%s\n", pcl)
			if err != nil {
				return "", diagnostics, err
			}
		}
	}
	if yamlFiles == 0 {
		return "", diagnostics, fmt.Errorf("unable to find any YAML files in directory: %s", dirName)
	}
	return buff.String(), diagnostics, err
}

func convert(testFiles ast.File) (string, hcl.Diagnostics, error) {
	var v Visitor
	var buff bytes.Buffer
	var err error
	diagnostics := hcl.Diagnostics{}

	for _, doc := range testFiles.Docs {
		if doc == nil || doc.Body == nil {
			continue
		}

		baseNodes, ok := doc.Body.(*ast.MappingNode)
		if !ok {
			return "", hcl.Diagnostics{}, fmt.Errorf("unable to parse yaml ast")
		}
		header, diag := getHeader(baseNodes.Values)
		// check diagnostics here and break at malformed resource then continue for other resources defined
		if diag.Severity == hcl.DiagnosticSeverity(1) {
			diagnostics = diagnostics.Append(&diag)
			break
		}
		_, err = fmt.Fprint(&buff, header)
		if err != nil {
			return "", diagnostics, err
		}
		err = walkToPCL(v, doc.Body, &buff, "")
		if err != nil {
			return "", diagnostics, err
		}
	}
	return buff.String(), diagnostics, err
}

// this will remove any incorrect quotes around the apiVersion of a yaml file
// apiVersion: "apps/v1" is just as valid as apiVersion: apps/v1
// so we should accept both
func trimQuotes(s string) string {
	if len(s) > 0 && s[0] == '"' {
		s = s[1:]
	}
	if len(s) > 0 && s[len(s)-1] == '"' {
		s = s[:len(s)-1]
	}
	return s
}

// resource <metadata/name> “kubernetes:<apiVersion>:<kind>”
func getHeader(nodes []*ast.MappingValueNode) (string, hcl.Diagnostic) {
	var apiVersion string
	for _, node := range nodes {
		if node.GetPath() == "$.apiVersion" {
			if valStr, ok := getStringMapNodeVal(node); ok {
				apiVersion = trimQuotes(valStr)
				break
			}
		}
	}
	// missing apiVersion
	if apiVersion == "" {
		return "", hcl.Diagnostic{
			Severity: hcl.DiagnosticSeverity(1),
			Summary:  "malformed yaml: apiVersion not specified",
			Detail:   "apiVersion field for the resource is not specified and is required",
		}
	}
	if !strings.Contains(apiVersion, "/") {
		apiVersion = fmt.Sprintf("core/%s", apiVersion)
	}

	name, kind := resourceName(nodes)
	// missing kind
	if kind == "" {
		return "", hcl.Diagnostic{
			Severity: hcl.DiagnosticSeverity(1),
			Summary:  "malformed yaml: resource kind not specified",
			Detail:   "kind field for the resource is not specified and is required",
		}
	}
	// kind is a CRD and will break the program
	if kind == "CustomResourceDefinition" {
		return "", hcl.Diagnostic{
			Severity: hcl.DiagnosticSeverity(1),
			Summary:  "contains CRD",
			Detail: "custom resource definitions cannot not be converted, please refer to \n" +
				"https://github.com/pulumi/crd2pulumi in order to convert your CRD",
		}
	}
	// missing name
	if name == kind {
		return "", hcl.Diagnostic{
			Severity: hcl.DiagnosticSeverity(1),
			Summary:  "malformed yaml: resource name not specified",
			Detail:   "name field within the metadata for the resource is not specified and is required",
		}
	}

	header := fmt.Sprintf(`resource %q "kubernetes:%s:%s" `, name, apiVersion, kind)
	return header, hcl.Diagnostic{}
}

// getResourceKNN iteratively passes through the Yaml ast nodes to obtain the
// resource's kind, name and namespace using DFS.
func getResourceKNN(nodes []*ast.MappingValueNode) (kind, name, ns string) {
	// Clone nodes slice so we can traverse with DFS without modifying object in-place.
	// We need to do this because we need to walk to the .metadata.{name,namespace} values
	// without modifying the original nodes slice and removing it from future traversals.
	stack := make([]*ast.MappingValueNode, len(nodes))
	copy(stack, nodes)

	// DFS search the required properties.
	for len(stack) != 0 {
		currNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		switch currNode.GetPath() {
		case "$.kind":
			if valStr, ok := getStringMapNodeVal(currNode); ok {
				kind = valStr
			}

		case "$.metadata.name":
			if valStr, ok := getStringMapNodeVal(currNode); ok {
				name = valStr
			}

		case "$.metadata.namespace":
			if valStr, ok := getStringMapNodeVal(currNode); ok {
				ns = valStr
			}

		case "$.metadata":
			// Add nested metadata nodes to our search stack, so we can walk to
			// .metadata.{name,namespace} values.
			// Add all child mapping values if the metadata block is a map of multiple map nodes, eg:
			// .metadata.{.name,.namespace.labels}.
			if nestedMappingNode, ok := currNode.Value.(*ast.MappingNode); ok {
				stack = append(stack, nestedMappingNode.Values...)
				continue
			}

			// The child node is a singular mapping value node if the metadata block only contains
			// one map item, eg. .metadata.name.
			if nestedMappingValueNode, ok := currNode.Value.(*ast.MappingValueNode); ok {
				stack = append(stack, nestedMappingValueNode)
			}
		}

		// Return early if we have all 3 information.
		if kind != "" && name != "" && ns != "" {
			return
		}
	}

	// Note: we can also have just kind and name, as an empty namespace means the default ns.
	return
}

// getStringMapNodeVal gets the string value of a map of string if the map key given
// matches the node map key.
func getStringMapNodeVal(mapValNode *ast.MappingValueNode) (string, bool) {
	nodeVal, ok := mapValNode.Value.(*ast.StringNode)
	if !ok {
		return "", false // Not a string value, skip.
	}

	return nodeVal.Value, true
}

// resourceName computes a name for the resource based on the namespace, name, and kind.
func resourceName(nodes []*ast.MappingValueNode) (string, string) {
	kind, name, ns := getResourceKNN(nodes)
	name = strings.ReplaceAll(strings.ToLower(name), "-", "_")
	ns = strings.ReplaceAll(strings.ToLower(ns), "-", "_")

	if len(ns) > 0 {
		// Capitalize the object's name if namespace is defined.
		name = cases.Title(language.English, cases.Compact).String(name)
		return fmt.Sprintf("%s%s%s", ns, name, kind), kind
	}

	return strings.ReplaceAll(fmt.Sprintf("%s%s", name, kind), ".", "_"), kind
}

// walkToPCL traverses an AST in depth-first order and converts the corresponding PCL code
func walkToPCL(v Visitor, node ast.Node, totalPCL io.Writer, suffix string) error {
	var err error

	tk := node.GetToken()
	switch n := node.(type) {
	case *ast.LiteralNode:
		multLine := n.Value.Value

		// note: PCL heredoc strings must have a trialing newline
		multLine = strings.TrimSpace(multLine)
		// Escape any ${} interpolation sequences.
		multLine = strings.ReplaceAll(multLine, "${", "$${")
		_, err = fmt.Fprintf(totalPCL, "<<EOF\n%s\nEOF\n", multLine)
		if err != nil {
			return err
		}

		comment := addComment(node)
		if comment != "" {
			_, err = fmt.Fprintf(totalPCL, "%s", comment)
			if err != nil {
				return err
			}
		}
	case *ast.NullNode:
		_, err = fmt.Fprintf(totalPCL, "null\n")
		if err != nil {
			return err
		}
		comment := addComment(node)
		if comment != "" {
			_, err = fmt.Fprintf(totalPCL, "%s", comment)
			if err != nil {
				return err
			}
		}
	case *ast.IntegerNode:
		_, err = fmt.Fprintf(totalPCL, "%v\n", n.Value)
		if err != nil {
			return err
		}
		comment := addComment(node)
		if comment != "" {
			_, err = fmt.Fprintf(totalPCL, "%s", comment)
			if err != nil {
				return err
			}
		}
	case *ast.FloatNode:
		_, err = fmt.Fprintf(totalPCL, "%v\n", n.Value)
		if err != nil {
			return err
		}
		comment := addComment(node)
		if comment != "" {
			_, err = fmt.Fprintf(totalPCL, "%s", comment)
			if err != nil {
				return err
			}
		}
	case *ast.StringNode:
		if tk.Next == nil || tk.Next.Value != ":" {
			s := n.Value

			// Escape any ${} interpolation sequences.
			s = strings.ReplaceAll(s, "${", "$${")
			strVal := fmt.Sprintf("%q%s", s, suffix)
			_, err = fmt.Fprintf(totalPCL, "%s\n", strVal)
			if err != nil {
				return err
			}
			// add comments on the same line
			comment := addComment(node)
			if comment != "" {
				_, err = fmt.Fprintf(totalPCL, "%s", comment)
				if err != nil {
					return err
				}
			}
		}
	case *ast.MergeKeyNode:
	case *ast.BoolNode:
		_, err = fmt.Fprintf(totalPCL, "%v\n", n.Value)
		if err != nil {
			return err
		}
		comment := addComment(node)
		if comment != "" {
			_, err = fmt.Fprintf(totalPCL, "%s", comment)
			if err != nil {
				return err
			}
		}
	case *ast.InfinityNode:
	case *ast.NanNode:
	case *ast.TagNode:
		err = walkToPCL(v, n.Value, totalPCL, "")
		if err != nil {
			return err
		}
	case *ast.DocumentNode:
		comment := addComment(node)
		if comment != "" {
			_, err = fmt.Fprintf(totalPCL, "%s", comment)
			if err != nil {
				return err
			}
		}
		err = walkToPCL(v, n.Body, totalPCL, "")
		if err != nil {
			return err
		}
	case *ast.MappingNode:
		_, err = fmt.Fprintf(totalPCL, "%s\n", "{")
		comment := addComment(node)
		if comment != "" {
			_, err = fmt.Fprintf(totalPCL, "%s", comment)
			if err != nil {
				return err
			}
		}
		if err != nil {
			return err
		}
		for _, value := range n.Values {
			err = walkToPCL(v, value, totalPCL, "")
			if err != nil {
				return err
			}
		}
		_, err = fmt.Fprintf(totalPCL, "%s%s\n", "}", suffix)
		if err != nil {
			return err
		}
	case *ast.MappingKeyNode:
		comment := addComment(node)
		if comment != "" {
			_, err = fmt.Fprintf(totalPCL, "%s", comment)
			if err != nil {
				return err
			}
		}
		err = walkToPCL(v, n.Value, totalPCL, "")
		if err != nil {
			return err
		}
	case *ast.MappingValueNode:
		comment := addComment(node)
		if comment != "" {
			_, err = fmt.Fprintf(totalPCL, "%s", comment)
			if err != nil {
				return err
			}
		}
		var key string
		switch nk := n.Key.(type) {
		case ast.ScalarNode:
			key = fmt.Sprintf("%v", nk.GetValue())
		default:
			return fmt.Errorf(fmt.Sprintf("unexpected key type: %T\n Please file an issue with the YAML input so we"+
				"can take a look: https://github.com/pulumi/kube2pulumi/issues/new", n.Key))
		}

		if !hclsyntax.ValidIdentifier(key) {
			_, err = fmt.Fprintf(totalPCL, "%q = ", key)
		} else {
			_, err = fmt.Fprintf(totalPCL, "%s = ", key)
		}
		if err != nil {
			return err
		}
		if n.Value.Type() == ast.MappingValueType {
			_, err = fmt.Fprintf(totalPCL, "%s\n", "{")
			if err != nil {
				return err
			}
		} else if n.Value.Type() == ast.SequenceType {
			_, err = fmt.Fprintf(totalPCL, "%s\n", "[")
			if err != nil {
				return err
			}
		}

		err = walkToPCL(v, n.Key, totalPCL, "")
		if err != nil {
			return err
		}
		err = walkToPCL(v, n.Value, totalPCL, "")
		if err != nil {
			return err
		}

		if n.Value.Type() == ast.MappingValueType {
			_, err = fmt.Fprintf(totalPCL, "%s%s\n", "}", suffix)
			if err != nil {
				return err
			}
		}
	case *ast.SequenceNode:
		suffix := ""
		for i, value := range n.Values {
			if len(n.Values) > 1 && i < len(n.Values)-1 {
				suffix = ","
			} else {
				suffix = ""
			}

			if value.Type() == ast.MappingValueType {
				_, err = fmt.Fprintf(totalPCL, "%s%s\n", "{", "")
				if err != nil {
					return err
				}
			}
			comment := addComment(node)
			if comment != "" {
				_, err = fmt.Fprintf(totalPCL, "%s", comment)
				if err != nil {
					return err
				}
			}
			err = walkToPCL(v, value, totalPCL, suffix)
			if err != nil {
				return err
			}
			if value.Type() == ast.MappingValueType {
				_, err = fmt.Fprintf(totalPCL, "%s%s\n", "}", suffix)
				if err != nil {
					return err
				}
			}
		}
		_, err = fmt.Fprintf(totalPCL, "%s\n", "]")
		if err != nil {
			return err
		}
	case *ast.AnchorNode:
		err = walkToPCL(v, n.Name, totalPCL, "")
		if err != nil {
			return err
		}
		err = walkToPCL(v, n.Value, totalPCL, "")
		if err != nil {
			return err
		}
	case *ast.AliasNode:
		err = walkToPCL(v, n.Value, totalPCL, "")
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf(fmt.Sprintf("unexpected node type: %s\n Please file an issue with the YAML input so we"+
			"can take a look: https://github.com/pulumi/kube2pulumi/issues/new", n.Type().String()))
	}

	return nil
}

func addComment(node ast.Node) string {
	/**
	check for comments here in order to add to the PCL string
	*/
	if comments := node.GetComment(); comments != nil {
		commentVal := ""
		for _, line := range comments.Comments {
			commentVal = fmt.Sprintf("#%s", line.Token.Value)
			// TODO handle multi-line comments
			break
		}
		return fmt.Sprintf("%s\n", commentVal)
	}
	return ""
}

type Visitor struct {
}

func (v *Visitor) Visit(node ast.Node) ast.Visitor {
	tk := node.GetToken()

	tk.Prev = nil
	return v
}

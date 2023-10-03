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
		baseNodes := ast.Filter(ast.MappingValueType, doc.Body)
		header, diag := getHeader(baseNodes)
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
func getHeader(nodes []ast.Node) (string, hcl.Diagnostic) {
	var apiVersion string
	for _, node := range nodes {
		if mapValNode, ok := node.(*ast.MappingValueNode); ok {
			if mapValNode.Key.String() == "apiVersion" {
				apiVersion = trimQuotes(mapValNode.Value.String())
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

// resourceName computes a name for the resource based on the namespace, name, and kind.
func resourceName(nodes []ast.Node) (string, string) {
	var kind, name, ns string
	for _, node := range nodes {
		if mapValNode, ok := node.(*ast.MappingValueNode); ok {
			if len(kind) == 0 && mapValNode.Key.String() == "kind" {
				kind = mapValNode.Value.String()
				continue
			}
			if mapValNode.Key.String() == "metadata" {
				if mapValNode.Value.Type() == ast.StringType {
					name = mapValNode.Value.String()
					continue
				} else {
					for _, inner := range ast.Filter(ast.MappingValueType, node) {
						if innerMvNode, ok := inner.(*ast.MappingValueNode); ok {
							if innerMvNode.Key.String() == "name" {
								name = innerMvNode.Value.String()
							} else if innerMvNode.Key.String() == "namespace" {
								ns = innerMvNode.Value.String()
							}
						}
					}
				}
			}
		}
		if len(ns) > 0 && len(name) > 0 {
			break
		}
	}

	name = strings.ReplaceAll(strings.ToLower(name), "-", "_")
	ns = strings.ReplaceAll(strings.ToLower(ns), "-", "_")

	if len(ns) > 0 {
		name = strings.ToUpper(string(name[0])) + name[1:]
		return fmt.Sprintf("%s%s%s", ns, name, kind), kind
	}

	return strings.ReplaceAll(fmt.Sprintf("%s%s", name, kind), ".", "_"), kind
}

// walkToPCL traverses an AST in depth-first order and converts the corresponding PCL code
func walkToPCL(v Visitor, node ast.Node, totalPCL io.Writer, suffix string) error {
	if v := v.Visit(node); v == nil {
		return nil
	}

	var err error

	tk := node.GetToken()
	switch n := node.(type) {
	case *ast.LiteralNode:
		multLine := node.String()
		multLine = strings.TrimPrefix(multLine, "|")
		multLine = strings.TrimSpace(multLine)
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
		_, err = fmt.Fprintf(totalPCL, "%s\n", node)
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
		_, err = fmt.Fprintf(totalPCL, "%s\n", node)
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
		_, err = fmt.Fprintf(totalPCL, "%s\n", node)
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
			s := n.String()
			// Remove leading quote if present.
			if len(s) > 0 && (s[0] == '"' || s[0] == '\'') {
				s = s[1:]
			}
			// Remove trailing quote if present unless it is escaped.
			if len(s) > 0 && (s[len(s)-1] == '"' || s[len(s)-1] == '\'') {
				if len(s) == 1 {
					s = ""
				}
				if len(s) > 1 && s[len(s)-2] != '\\' {
					s = s[:len(s)-1]
				}
			}
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
		_, err = fmt.Fprintf(totalPCL, "%s\n", node)
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
		_, err = fmt.Fprintf(totalPCL, "%s\n", node)
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
		key := n.Key.String()
		// trim surrounding quotations if there
		if len(key) >= 2 {
			if key[0] == '"' && key[len(key)-1] == '"' {
				key = key[1 : len(key)-1]
			}
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
	if comment := node.GetComment(); comment != nil {
		commentVal := strings.TrimSpace(comment.Value)
		if !strings.HasPrefix(commentVal, "#") {
			commentVal = fmt.Sprintf("# %s", commentVal)
		}
		return fmt.Sprintf("%s\n", commentVal)
	}
	return ""
}

type Visitor struct {
}

func (v *Visitor) Visit(node ast.Node) ast.Visitor {
	tk := node.GetToken()

	if comment := node.GetComment(); comment != nil {
		comment.Prev = nil
		comment.Next = nil
	}

	tk.Prev = nil
	return v
}

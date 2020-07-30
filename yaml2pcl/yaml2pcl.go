// package yaml2pcl provides a method to convert k8s yaml to
// PCL (pulumi schema)
package yaml2pcl

import (
	"github.com/goccy/go-yaml/ast"
	"github.com/goccy/go-yaml/parser"
	"strings"
)

// Convert returns a string conversion of the input YAML
// as a byte array into PCL: sample below
// Output: resource foo "kubernetes:core/v1:Namespace" {
// apiVersion = "v1"
// kind = "Namespace"
// metadata = {
// name = "foo"
// }
// }
func Convert(input []byte) (string, error) {
	testFiles, err := parser.ParseBytes(input, parser.ParseComments)
	if err != nil {
		return "", err
	}
	var pcl string
	var v Visitor
	for _, doc := range testFiles.Docs {
		baseNodes := ast.Filter(ast.MappingValueType, doc.Body)
		pcl += getHeader(baseNodes)
		pcl = walkToPCL(v, doc.Body, pcl)
	}
	return pcl, err
}

// resource <metadata/name> “kubernetes : <apiVersion>: <kind>”
func getHeader(nodes []ast.Node) string {
	var apiVersion string
	for _, node := range nodes {
		if node.(*ast.MappingValueNode).Key.String() == "apiVersion" {
			apiVersion = node.(*ast.MappingValueNode).Value.String()
			break
		}
	}
	if !strings.Contains(apiVersion, "/") {
		apiVersion = "core/" + apiVersion
	}

	metaName := getMetaName(nodes)

	var kind string
	for _, node := range nodes {
		if node.(*ast.MappingValueNode).Key.String() == "kind" {
			kind = node.(*ast.MappingValueNode).Value.String()
			break
		}
	}

	header := strings.Join([]string{"resource ", metaName, " \"kubernetes:", apiVersion, ":", kind, "\" "}, "")
	return header
}

// returns <metadata/name> field as a string from AST
func getMetaName(nodes []ast.Node) string {
	for _, node := range nodes {
		if node.(*ast.MappingValueNode).Key.String() == "metadata" {
			if node.(*ast.MappingValueNode).Value.Type() == ast.StringType {
				return node.(*ast.MappingValueNode).Value.String()
			} else {
				for _, inner := range ast.Filter(ast.MappingValueType, node) {
					if inner.(*ast.MappingValueNode).Key.String() == "name" {
						return inner.(*ast.MappingValueNode).Value.String()
					}
				}
			}
		}
	}
	return ""
}

// walkToPCL traverses an AST in depth-first order and converts the corresponding PCL code
func walkToPCL(v Visitor, node ast.Node, totalPCL string) string {
	if v := v.Visit(node); v == nil {
		return ""
	}

	tk := node.GetToken()
	/**
	check for comments here in order to add to the PCL string
	*/
	if comment := node.GetComment(); comment != nil {
		totalPCL += comment.Value
	}

	switch n := node.(type) {
	case *ast.NullNode:
	case *ast.IntegerNode:
		totalPCL += node.String() + "\n"
	case *ast.FloatNode:
		totalPCL += node.String() + "\n"
	case *ast.StringNode:
		if tk.Next == nil || tk.Next.Value != ":" {
			totalPCL += strings.Join([]string{"\"", n.String(), "\"", "\n"}, "")
		}
	case *ast.MergeKeyNode:
	case *ast.BoolNode:
		totalPCL += node.String() + "\n"
	case *ast.InfinityNode:
	case *ast.NanNode:
	case *ast.TagNode:
		totalPCL = walkToPCL(v, n.Value, totalPCL)
	case *ast.DocumentNode:
		totalPCL = walkToPCL(v, n.Body, totalPCL)
	case *ast.MappingNode:
		totalPCL += "{\n"
		for _, value := range n.Values {
			totalPCL = walkToPCL(v, value, totalPCL)
		}
		totalPCL += "}\n"
	case *ast.MappingKeyNode:
		totalPCL = walkToPCL(v, n.Value, totalPCL)
	case *ast.MappingValueNode:
		totalPCL += n.Key.String() + " = "
		if n.Value.Type() == ast.MappingValueType {
			totalPCL += "{\n"
		} else if n.Value.Type() == ast.SequenceType {
			totalPCL += "[\n"
		}

		totalPCL = walkToPCL(v, n.Key, totalPCL)
		totalPCL = walkToPCL(v, n.Value, totalPCL)

		if n.Value.Type() == ast.MappingValueType {
			totalPCL += "}\n"
		}
	case *ast.SequenceNode:
		for _, value := range n.Values {
			totalPCL = walkToPCL(v, value, totalPCL)
		}
		totalPCL += "]\n"
	case *ast.AnchorNode:
		totalPCL = walkToPCL(v, n.Name, totalPCL)
		totalPCL = walkToPCL(v, n.Value, totalPCL)
	case *ast.AliasNode:
		totalPCL = walkToPCL(v, n.Value, totalPCL)
	}
	return totalPCL
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

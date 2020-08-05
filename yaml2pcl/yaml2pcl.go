// package yaml2pcl provides a method to convert k8s yaml to
// PCL (pulumi schema)
package yaml2pcl

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"strings"

	"github.com/goccy/go-yaml/ast"
	"github.com/goccy/go-yaml/parser"
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
	return convert(*testFiles)
}

// ConvertFile returns a string conversion of the input YAML
// in a file into PCL: sample below
// Output: resource foo "kubernetes:core/v1:Namespace" {
// apiVersion = "v1"
// kind = "Namespace"
// metadata = {
// name = "foo"
// }
// }
func ConvertFile(filename string) (string, error) {
	testFiles, err := parser.ParseFile(filename, parser.ParseComments)
	if err != nil {
		return "", err
	}
	return convert(*testFiles)
}

func ConvertDirectory(dirName string) (string, error) {
	var buff bytes.Buffer
	files, _ := ioutil.ReadDir(dirName)
	for _, file := range files {
		if strings.Contains(file.Name(), ".yaml") {
			pcl, err := ConvertFile(dirName + file.Name())
			if err != nil {
				return "", err
			}
			_, err = fmt.Fprintf(&buff, "%s\n", pcl)
			if err != nil {
				return "", err
			}
		}
	}
	return buff.String(), nil
}

func convert(testFiles ast.File) (string, error) {
	var v Visitor
	var buff bytes.Buffer
	var err error

	for _, doc := range testFiles.Docs {
		baseNodes := ast.Filter(ast.MappingValueType, doc.Body)
		_, err = fmt.Fprint(&buff, getHeader(baseNodes))
		if err != nil {
			return "", err
		}
		err = walkToPCL(v, doc.Body, &buff, "")
		if err != nil {
			return "", err
		}
	}
	return buff.String(), err
}

// resource <metadata/name> “kubernetes:<apiVersion>:<kind>”
func getHeader(nodes []ast.Node) string {
	var apiVersion string
	for _, node := range nodes {
		if mapValNode, ok := node.(*ast.MappingValueNode); ok {
			if mapValNode.Key.String() == "apiVersion" {
				apiVersion = mapValNode.Value.String()
				break
			}
		}
	}
	if !strings.Contains(apiVersion, "/") {
		apiVersion = fmt.Sprintf("core/%s", apiVersion)
	}

	metaName := getMetaName(nodes)

	var kind string
	for _, node := range nodes {
		if mapValNode, ok := node.(*ast.MappingValueNode); ok {
			if mapValNode.Key.String() == "kind" {
				kind = mapValNode.Value.String()
				break
			}
		}
	}
	header := fmt.Sprintf(`resource %s "kubernetes:%s:%s" `, metaName, apiVersion, kind)
	return header
}

// returns <metadata/name> field as a string from AST
func getMetaName(nodes []ast.Node) string {
	for _, node := range nodes {
		if mapValNode, ok := node.(*ast.MappingValueNode); ok {
			if mapValNode.Key.String() == "metadata" {
				if mapValNode.Value.Type() == ast.StringType {
					return mapValNode.Value.String()
				} else {
					for _, inner := range ast.Filter(ast.MappingValueType, node) {
						if innerMvNode, ok := inner.(*ast.MappingValueNode); ok {
							if innerMvNode.Key.String() == "name" {
								return innerMvNode.Value.String()
							}
						}
					}
				}
			}
		}
	}
	return ""
}

// walkToPCL traverses an AST in depth-first order and converts the corresponding PCL code
func walkToPCL(v Visitor, node ast.Node, totalPCL io.Writer, suffix string) error {
	if v := v.Visit(node); v == nil {
		return nil
	}

	var err error
	tk := node.GetToken()
	/**
	check for comments here in order to add to the PCL string
	*/
	if comment := node.GetComment(); comment != nil {
		_, err = fmt.Fprintf(totalPCL, "%s", comment.Value)
		if err != nil {
			return err
		}
	}
	switch n := node.(type) {
	case *ast.NullNode:
	case *ast.IntegerNode:
		_, err = fmt.Fprintf(totalPCL, "%s\n", node)
		if err != nil {
			return err
		}
	case *ast.FloatNode:
		_, err = fmt.Fprintf(totalPCL, "%s\n", node)
		if err != nil {
			return err
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
			strVal := fmt.Sprintf("%q%s\n", s, suffix)
			_, err = fmt.Fprintf(totalPCL, "%s", strVal)
			if err != nil {
				return err
			}
		}
	case *ast.MergeKeyNode:
	case *ast.BoolNode:
		_, err = fmt.Fprintf(totalPCL, "%s\n", node)
		if err != nil {
			return err
		}
	case *ast.InfinityNode:
	case *ast.NanNode:
	case *ast.TagNode:
		err = walkToPCL(v, n.Value, totalPCL, "")
		if err != nil {
			return err
		}
	case *ast.DocumentNode:
		err = walkToPCL(v, n.Body, totalPCL, "")
		if err != nil {
			return err
		}
	case *ast.MappingNode:
		_, err = fmt.Fprintf(totalPCL, "%s\n", "{")
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
		err = walkToPCL(v, n.Value, totalPCL, "")
		if err != nil {
			return err
		}
	case *ast.MappingValueNode:
		_, err = fmt.Fprintf(totalPCL, "%s = ", n.Key)
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
			err = walkToPCL(v, value, totalPCL, suffix)
			if err != nil {
				return err
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
		return errors.New("unexpected node type: " + n.Type().String())
	}
	return nil
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

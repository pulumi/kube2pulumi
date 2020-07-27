package main

import (
	"fmt"
	"github.com/goccy/go-yaml/ast"
	"github.com/goccy/go-yaml/parser"
	"github.com/goccy/go-yaml/token"
	"strings"
)

var testData = `

apiVersion: v1
kind: Pod
metadata:
  namespace: foo
  name: bar
spec:
  containers:
    - name: nginx
      image: nginx:1.14-alpine
      resources:
        limits:
          memory: 20Mi
          cpu: 0.2

`

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

// resource <metadata/name> “kubernetes : <apiVersion>: <kind>” {
func printHeader(nodes []ast.Node) {
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

	header := strings.Join([]string{"resource ", metaName, " \"kubernetes:", apiVersion, ":", kind, "\"", " {"}, "")
	fmt.Println(header)
}

func main() {
	testFiles, err := parser.ParseBytes([]byte(testData), parser.ParseComments)
	if err != nil {
		fmt.Println(err)
	}
	var v Visitor
	for _, doc := range testFiles.Docs {
		baseNodes := ast.Filter(ast.MappingValueType, doc.Body)
		printHeader(baseNodes)
		walkToPCL(v, doc.Body)
	}
	fmt.Println("}")
}

func nodeToPCL(node ast.Node, tk *token.Token) {
	if node.Type() == ast.MappingValueType {
		mapNode := node.(*ast.MappingValueNode)
		fmt.Print(mapNode.Key, " = ")
		if mapNode.Value.Type() == ast.MappingType || mapNode.Value.Type() == ast.MappingValueType {
			fmt.Println("{")
		}
		if mapNode.Value.Type() == ast.SequenceType {
			fmt.Println("[\n{")
		}
	}
	if node.Type() == ast.StringType && (tk.Next == nil || tk.Next.Value != ":") {
		value := strings.Join([]string{"\"", node.(*ast.StringNode).String(), "\""}, "")
		fmt.Println(value)
	}
	if node.Type() == ast.FloatType {
		fmt.Println(node.(*ast.FloatNode).String())
	}
	if node.Type() == ast.IntegerType {
		fmt.Println(node.(*ast.IntegerNode).String())
	}
	if node.Type() == ast.BoolType {
		fmt.Println(node.(*ast.BoolNode).String())
	}
}

// walktoPCL traverses an AST in depth-first order and prints out the corresponding PCL code:
// Starts by calling v.Visit(node); node must not be nil.
// If the visitor w returned by v.Visit(node) is not nil,
// walktoPCL is invoked recursively with visitor w for each of the non-nil children of node,
// followed by a call of w.Visit(nil).
func walkToPCL(v Visitor, node ast.Node) {
	if v := v.Visit(node); v == nil {
		return
	}

	switch n := node.(type) {
	case *ast.NullNode:
	case *ast.IntegerNode:
	case *ast.FloatNode:
	case *ast.StringNode:
	case *ast.MergeKeyNode:
	case *ast.BoolNode:
	case *ast.InfinityNode:
	case *ast.NanNode:
	case *ast.TagNode:
		walkToPCL(v, n.Value)
	case *ast.DocumentNode:
		walkToPCL(v, n.Body)
	case *ast.MappingNode:
		for _, value := range n.Values {
			walkToPCL(v, value)
		}
		fmt.Println("}")
	case *ast.MappingKeyNode:
		walkToPCL(v, n.Value)
	case *ast.MappingValueNode:
		walkToPCL(v, n.Key)
		walkToPCL(v, n.Value)
	case *ast.SequenceNode:
		for _, value := range n.Values {
			walkToPCL(v, value)
		}
		fmt.Println("}\n]")
	case *ast.AnchorNode:
		walkToPCL(v, n.Name)
		walkToPCL(v, n.Value)
	case *ast.AliasNode:
		walkToPCL(v, n.Value)
	}
}

type Visitor struct {
}

func (v *Visitor) Visit(node ast.Node) ast.Visitor {
	tk := node.GetToken()

	// comments must be printed out first to maintain original ordering
	if comment := node.GetComment(); comment != nil {
		comment.Prev = nil
		comment.Next = nil
		fmt.Print(comment.Value)
	}

	// conversion
	nodeToPCL(node, tk)

	tk.Prev = nil
	tk.Next = nil

	return v
}

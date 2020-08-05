package yaml2pcl

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestNamespace(t *testing.T) {
	assertion := assert.New(t)

	expected := `resource foo "kubernetes:core/v1:Namespace" {
apiVersion = "v1"
kind = "Namespace"
metadata = {
name = "foo"
}
}
`
	result, err := ConvertFile("../testdata/Namespace.yaml")
	if err != nil {
		assertion.NoError(err)
	} else {
		assertion.Equal(expected, result, "Single resource conversion was incorrect")
	}
}

func TestNamespaceComments(t *testing.T) {
	assertion := assert.New(t)

	expected := `resource foo "kubernetes:core/v1:Namespace" {
apiVersion = "v1"
kind = "Namespace"
# this is a codegentest comment
metadata = {
name = "foo"
}
}
`
	result, err := ConvertFile("../testdata/NamespaceWithComments.yaml")
	if err != nil {
		assertion.NoError(err)
	} else {
		assertion.Equal(expected, result, "Comments are converted incorrectly")
	}
}

func Test1PodArray(t *testing.T) {
	assertion := assert.New(t)

	expected := `resource bar "kubernetes:core/v1:Pod" {
apiVersion = "v1"
kind = "Pod"
metadata = {
namespace = "foo"
name = "bar"
}
spec = {
containers = [
{
name = "nginx"
image = "nginx:1.14-alpine"
resources = {
limits = {
memory = "20Mi"
cpu = 0.2
}
}
}
]
}
}
`
	result, err := ConvertFile("../testdata/OnePodArray.yaml")
	if err != nil {
		assertion.NoError(err)
	} else {
		assertion.Equal(expected, result, "Nested array is converted incorrectly")
	}
}

func TestRole(t *testing.T) {
	assertion := assert.New(t)

	b, err := ioutil.ReadFile(filepath.Join("..", "testdata", "Role.pp"))
	if err != nil {
		assertion.NoError(err)
	}
	expected := string(b)

	result, err := ConvertFile(filepath.Join("..", "testdata", "Role.yaml"))
	if err != nil {
		assertion.NoError(err)
	} else {
		assertion.Equal(expected, result, "Role is converted incorrectly")
	}
}

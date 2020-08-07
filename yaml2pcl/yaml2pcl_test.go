package yaml2pcl

import (
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNamespace(t *testing.T) {
	assertion := assert.New(t)

	expected := `resource fooNamespace "kubernetes:core/v1:Namespace" {
apiVersion = "v1"
kind = "Namespace"
metadata = {
name = "foo"
}
}
`
	result, err := ConvertFile("../testdata/Namespace.yaml")
	assertion.NoError(err)
	assertion.Equal(expected, result, "Single resource conversion was incorrect")
}

func TestNamespaceComments(t *testing.T) {
	assertion := assert.New(t)

	expected := `resource fooNamespace "kubernetes:core/v1:Namespace" {
apiVersion = "v1"
kind = "Namespace"
# this is a codegentest comment
metadata = {
name = "foo"
}
}
`
	result, err := ConvertFile("../testdata/NamespaceWithComments.yaml")
	assertion.NoError(err)
	assertion.Equal(expected, result, "Comments are converted incorrectly")
}

func Test1PodArray(t *testing.T) {
	assertion := assert.New(t)

	expected := `resource fooBarPod "kubernetes:core/v1:Pod" {
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
	assertion.NoError(err)
	assertion.Equal(expected, result, "Nested array is converted incorrectly")
}

func TestRole(t *testing.T) {
	assertion := assert.New(t)

	b, err := ioutil.ReadFile(filepath.Join("..", "testdata", "Role.pp"))
	assertion.NoError(err)
	expected := string(b)

	result, err := ConvertFile(filepath.Join("..", "testdata", "Role.yaml"))
	assertion.NoError(err)
	assertion.Equal(expected, result, "Role is converted incorrectly")
}

func TestDirk8sOperator(t *testing.T) {
	assertion := assert.New(t)

	b, err := ioutil.ReadFile(filepath.Join("..", "testdata", "expK8sOperator.pp"))
	assertion.NoError(err)
	expected := string(b)

	result, err := ConvertDirectory("../testdata/k8sOperator/")
	assertion.NoError(err)
	assertion.Equal(expected, result, "Directory is converted incorrectly")
}

func TestNamespaceTrailingComments(t *testing.T) {
	assertion := assert.New(t)

	expected := `resource fooNamespace "kubernetes:core/v1:Namespace" {
apiVersion = "v1"
kind = "Namespace"
# this is a trailing comment
metadata = {
name = "foo"
}
}
`
	result, err := ConvertFile("../testdata/NamespaceWithTrailingComment.yaml")
	assertion.NoError(err)
	assertion.Equal(expected, result, "Comments are converted incorrectly")
}

func TestIncorrectPath(t *testing.T) {
	assertion := assert.New(t)
	fakePath := "fakePath"
	_, err := ConvertFile(fakePath)
	assertion.Error(err)
	_, err = ConvertDirectory(fakePath)
	assertion.Error(err)
}

func TestMalformedHeaderYaml(t *testing.T) {
	assertion := assert.New(t)
	_, err := ConvertFile("../testdata/MalformedYaml.yaml")
	assertion.Error(err)
}

func TestMultipleResourceGen(t *testing.T) {
	assertion := assert.New(t)

	b, err := ioutil.ReadFile(filepath.Join("..", "testdata", "MultipleResources.pcl"))
	assertion.NoError(err)
	expected := string(b)

	result, _ := ConvertFile("../testdata/MultipleResources.yml")
	assertion.NoError(err)
	assertion.Equal(expected, result, "File with multiple resources is converted incorrectly")
}

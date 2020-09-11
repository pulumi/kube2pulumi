package yaml2pcl

import (
	"fmt"
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
	result, err := ConvertFile("../../testdata/Namespace.yaml")
	fmt.Println(result)
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
	result, err := ConvertFile("../../testdata/NamespaceWithComments.yaml")
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
	result, err := ConvertFile("../../testdata/OnePodArray.yaml")
	assertion.NoError(err)
	assertion.Equal(expected, result, "Nested array is converted incorrectly")
}

func TestRole(t *testing.T) {
	assertion := assert.New(t)

	b, err := ioutil.ReadFile(filepath.Join("../..", "testdata", "Role.pp"))
	assertion.NoError(err)
	expected := string(b)

	result, err := ConvertFile(filepath.Join("../..", "testdata", "Role.yaml"))
	assertion.NoError(err)
	assertion.Equal(expected, result, "Role is converted incorrectly")
}

func TestDirk8sOperator(t *testing.T) {
	assertion := assert.New(t)

	b, err := ioutil.ReadFile(filepath.Join("../..", "testdata", "expK8sOperator.pp"))
	assertion.NoError(err)
	expected := string(b)

	result, err := ConvertDirectory("../../testdata/k8sOperator/")
	assertion.NoError(err)
	assertion.Equal(expected, result, "Directory is converted incorrectly")
}

func TestNamespaceTrailingComments(t *testing.T) {
	assertion := assert.New(t)

	expected := `resource fooNamespace "kubernetes:core/v1:Namespace" {
apiVersion = "v1"
# this is a trailing comment
kind = "Namespace"
metadata = {
name = "foo"
}
}
`
	result, err := ConvertFile("../../testdata/NamespaceWithTrailingComment.yaml")
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
	result, err := ConvertFile("../../testdata/MalformedYaml.yaml")
	fmt.Println(result)
	assertion.NoError(err)
}

func TestMultipleResourceGen(t *testing.T) {
	assertion := assert.New(t)

	b, err := ioutil.ReadFile(filepath.Join("../..", "testdata", "MultipleResources.pcl"))
	assertion.NoError(err)
	expected := string(b)

	result, err := ConvertFile("../../testdata/MultipleResources.yml")
	assertion.NoError(err)
	assertion.Equal(expected, result, "File with multiple resources is converted incorrectly")
}

func TestEmptyDir(t *testing.T) {
	assertion := assert.New(t)
	_, err := ConvertDirectory("../../testdata/empty/")
	assertion.Error(err)
	assertion.Contains(err.Error(), "unable to find any YAML files")
}

func TestAnnotationsDeployment(t *testing.T) {
	assertion := assert.New(t)

	b, err := ioutil.ReadFile(filepath.Join("../..", "testdata", "testDep.pp"))
	assertion.NoError(err)
	expected := string(b)

	result, err := ConvertFile("../../testdata/testDep.yaml")
	assertion.Equal(expected, result, "pcl is incorrect")
}

func TestNoDoubleQuotes(t *testing.T) {
	assertion := assert.New(t)

	b, err := ioutil.ReadFile(filepath.Join("../..", "testdata", "doubleQuotes.pp"))
	assertion.NoError(err)
	expected := string(b)

	result, err := ConvertFile("../../testdata/doubleQuotes.yaml")
	assertion.NoError(err)
	assertion.Equal(expected, result, "double quotes inserted")
}

func TestSpecialChar(t *testing.T) {
	assertion := assert.New(t)

	b, err := ioutil.ReadFile(filepath.Join("../..", "testdata", "specialChar.pp"))
	assertion.NoError(err)
	expected := string(b)

	result, err := ConvertFile("../../testdata/specialChar.yaml")
	assertion.NoError(err)
	assertion.Equal(expected, result, "double quotes inserted")
}

func TestMultiLineString(t *testing.T) {
	assertion := assert.New(t)

	b, err := ioutil.ReadFile(filepath.Join("../..", "testdata", "MultilineString.pp"))
	assertion.NoError(err)
	expected := string(b)

	result, err := ConvertFile("../../testdata/MultilineString.yaml")
	assertion.NoError(err)
	assertion.Equal(expected, result, "incorrectly parses multiline strings")
}

func TestCRD(t *testing.T) {
	assertion := assert.New(t)
	_, err := ConvertFile("../../testdata/customResourceDef.yaml")
	assertion.NoError(err)
}

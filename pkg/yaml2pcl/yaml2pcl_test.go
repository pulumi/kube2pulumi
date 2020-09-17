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
	result, diags, err := ConvertFile(filepath.Join("..", "..", "testdata", "Namespace.yaml"))
	if diags != nil {
		assertion.False(diags.HasErrors(), "diagnostics incorrectly displayed for proper yaml")
	}
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
	result, diags, err := ConvertFile(filepath.Join("..", "..", "testdata", "NamespaceWithComments.yaml"))
	if diags != nil {
		assertion.False(diags.HasErrors(), "diagnostics incorrectly displayed for proper yaml")
	}
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
	result, diags, err := ConvertFile(filepath.Join("..", "..", "testdata", "OnePodArray.yaml"))
	if diags != nil {
		assertion.False(diags.HasErrors(), "diagnostics incorrectly displayed for proper yaml")
	}
	assertion.NoError(err)
	assertion.Equal(expected, result, "Nested array is converted incorrectly")
}

func TestRole(t *testing.T) {
	assertion := assert.New(t)

	b, err := ioutil.ReadFile(filepath.Join("..", "..", "testdata", "Role.pp"))
	assertion.NoError(err)
	expected := string(b)

	result, diags, err := ConvertFile(filepath.Join("..", "..", "testdata", "Role.yaml"))
	if diags != nil {
		assertion.False(diags.HasErrors(), "diagnostics incorrectly displayed for proper yaml")
	}
	assertion.NoError(err)
	assertion.Equal(expected, result, "Role is converted incorrectly")
}

func TestDirk8sOperator(t *testing.T) {
	assertion := assert.New(t)

	b, err := ioutil.ReadFile(filepath.Join("..", "..", "testdata", "expK8sOperator.pp"))
	assertion.NoError(err)
	expected := string(b)

	result, diags, err := ConvertDirectory(filepath.Join("..", "..", "testdata", "k8sOperator/"))
	if diags != nil {
		assertion.False(diags.HasErrors(), "diagnostics incorrectly displayed for proper yaml")
	}
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
	result, diags, err := ConvertFile(filepath.Join("..", "..", "testdata", "NamespaceWithTrailingComment.yaml"))
	if diags != nil {
		assertion.False(diags.HasErrors(), "diagnostics incorrectly displayed for proper yaml")
	}
	assertion.NoError(err)
	assertion.Equal(expected, result, "Comments are converted incorrectly")
}

func TestIncorrectPath(t *testing.T) {
	assertion := assert.New(t)
	fakePath := "fakePath"
	_, _, err := ConvertFile(fakePath)
	assertion.Error(err)
	_, _, err = ConvertDirectory(fakePath)
	assertion.Error(err)
}

func TestMalformedHeaderYaml(t *testing.T) {
	assertion := assert.New(t)
	_, diags, err := ConvertFile(filepath.Join("..", "..", "testdata", "MalformedYaml.yaml"))
	if diags != nil {
		assertion.True(diags.HasErrors(), "diagnostics incorrectly displayed for wrongly formatted yaml")
	}
	assertion.NoError(err)
}

func TestMultipleResourceGen(t *testing.T) {
	assertion := assert.New(t)

	b, err := ioutil.ReadFile(filepath.Join("..", "..", "testdata", "MultipleResources.pcl"))
	assertion.NoError(err)
	expected := string(b)

	result, diags, err := ConvertFile(filepath.Join("..", "..", "testdata", "MultipleResources.yml"))
	if diags != nil {
		assertion.False(diags.HasErrors(), "diagnostics incorrectly displayed for proper yaml")
	}
	assertion.NoError(err)
	assertion.Equal(expected, result, "File with multiple resources is converted incorrectly")
}

func TestEmptyDir(t *testing.T) {
	assertion := assert.New(t)
	_, diags, err := ConvertDirectory(filepath.Join("..", "..", "testdata", "empty/"))
	if diags != nil {
		assertion.False(diags.HasErrors(), "diagnostics incorrectly displayed for proper yaml")
	}
	assertion.Error(err)
	assertion.Contains(err.Error(), "unable to find any YAML files")
}

func TestAnnotationsDeployment(t *testing.T) {
	assertion := assert.New(t)

	b, err := ioutil.ReadFile(filepath.Join("..", "..", "testdata", "testDep.pp"))
	assertion.NoError(err)
	expected := string(b)

	result, diags, err := ConvertFile(filepath.Join("..", "..", "testdata", "testDep.yaml"))
	if diags != nil {
		assertion.False(diags.HasErrors(), "diagnostics incorrectly displayed for proper yaml")
	}
	assertion.Equal(expected, result, "pcl is incorrect")
}

func TestNoDoubleQuotes(t *testing.T) {
	assertion := assert.New(t)

	b, err := ioutil.ReadFile(filepath.Join("..", "..", "testdata", "doubleQuotes.pp"))
	assertion.NoError(err)
	expected := string(b)

	result, diags, err := ConvertFile(filepath.Join("..", "..", "testdata", "doubleQuotes.yaml"))
	if diags != nil {
		assertion.False(diags.HasErrors(), "diagnostics incorrectly displayed for proper yaml")
	}
	assertion.NoError(err)
	assertion.Equal(expected, result, "double quotes inserted")
}

func TestSpecialChar(t *testing.T) {
	assertion := assert.New(t)

	b, err := ioutil.ReadFile(filepath.Join("..", "..", "testdata", "specialChar.pp"))
	assertion.NoError(err)
	expected := string(b)

	result, diags, err := ConvertFile(filepath.Join("..", "..", "testdata", "specialChar.yaml"))
	if diags != nil {
		assertion.False(diags.HasErrors(), "diagnostics incorrectly displayed for proper yaml")
	}
	assertion.NoError(err)
	assertion.Equal(expected, result, "double quotes inserted")
}

func TestMultiLineString(t *testing.T) {
	assertion := assert.New(t)

	b, err := ioutil.ReadFile(filepath.Join("..", "..", "testdata", "MultilineString.pp"))
	assertion.NoError(err)
	expected := string(b)

	result, diags, err := ConvertFile(filepath.Join("..", "..", "testdata", "MultilineString.yaml"))
	if diags != nil {
		assertion.False(diags.HasErrors(), "diagnostics incorrectly displayed for proper yaml")
	}
	assertion.NoError(err)
	assertion.Equal(expected, result, "incorrectly parses multiline strings")
}

func TestCRD(t *testing.T) {
	assertion := assert.New(t)
	_, diags, err := ConvertFile(filepath.Join("..", "..", "testdata", "customResourceDef.yaml"))
	if diags != nil {
		assertion.True(diags.HasErrors(), "diagnostics not detecting CRD")
	}
	assertion.NoError(err)
}

func TestNotYaml(t *testing.T) {
	assertion := assert.New(t)
	_, _, err := ConvertFile(filepath.Join("..", "..", "testdata", "empty", "notYAML.txt"))
	assertion.Error(err)
}

package yaml2pcl

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/pulumi/kube2pulumi/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

func TestNamespace(t *testing.T) {
	assertion := assert.New(t)

	expected := `resource "fooNamespace" "kubernetes:core/v1:Namespace" {
apiVersion = "v1"
kind = "Namespace"
metadata = {
name = "foo"
}
}
`
	testDir := testutil.MakeTestDir(t, filepath.Join("..", "..", "testdata", "Namespace"))
	result, diags, err := ConvertFile(filepath.Join(testDir, "Namespace.yaml"))
	if diags != nil {
		assertion.False(diags.HasErrors(), "diagnostics incorrectly displayed for proper yaml")
	}
	assertion.NoError(err)
	assertion.Equal(expected, result, "Single resource conversion was incorrect")
}

func TestNamespaceComments(t *testing.T) {
	assertion := assert.New(t)

	expected := `resource "fooNamespace" "kubernetes:core/v1:Namespace" {
apiVersion = "v1"
kind = "Namespace"
# this is a codegentest comment
metadata = {
name = "foo"
}
}
`
	testDir := testutil.MakeTestDir(t, filepath.Join("..", "..", "testdata", "Namespace"))
	result, diags, err := ConvertFile(filepath.Join(testDir, "NamespaceWithComments.yaml"))
	if diags != nil {
		assertion.False(diags.HasErrors(), "diagnostics incorrectly displayed for proper yaml")
	}
	assertion.NoError(err)
	assertion.Equal(expected, result, "Comments are converted incorrectly")
}

func Test1PodArray(t *testing.T) {
	assertion := assert.New(t)

	expected := `resource "fooBarPod" "kubernetes:core/v1:Pod" {
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
	testDir := testutil.MakeTestDir(t, filepath.Join("..", "..", "testdata", "OnePodArray"))
	result, diags, err := ConvertFile(filepath.Join(testDir, "OnePodArray.yaml"))
	if diags != nil {
		assertion.False(diags.HasErrors(), "diagnostics incorrectly displayed for proper yaml")
	}
	assertion.NoError(err)
	assertion.Equal(expected, result, "Nested array is converted incorrectly")
}

func TestRole(t *testing.T) {
	assertion := assert.New(t)

	testDir := testutil.MakeTestDir(t, filepath.Join("..", "..", "testdata", "Role"))
	b, err := os.ReadFile(filepath.Join(testDir, "Role.pp"))
	assertion.NoError(err)
	expected := string(b)

	result, diags, err := ConvertFile(filepath.Join(testDir, "Role.yaml"))
	if diags != nil {
		assertion.False(diags.HasErrors(), "diagnostics incorrectly displayed for proper yaml")
	}
	assertion.NoError(err)
	assertion.Equal(expected, result, "Role is converted incorrectly")
}

func TestDirk8sOperator(t *testing.T) {
	assertion := assert.New(t)

	testDir := testutil.MakeTestDir(t, filepath.Join("..", "..", "testdata", "k8sOperator"))
	b, err := os.ReadFile(filepath.Join(testDir, "expected", "expectedK8sOperator.pp"))
	assertion.NoError(err)
	expected := string(b)

	result, diags, err := ConvertDirectory(testDir)
	if diags != nil {
		assertion.False(diags.HasErrors(), "diagnostics incorrectly displayed for proper yaml")
	}
	assertion.NoError(err)
	assertion.Equal(expected, result, "Directory is converted incorrectly")
}

func TestNamespaceTrailingComments(t *testing.T) {
	assertion := assert.New(t)

	expected := `resource "fooNamespace" "kubernetes:core/v1:Namespace" {
apiVersion = "v1"
# this is a trailing comment
kind = "Namespace"
metadata = {
name = "foo"
}
}
`
	testDir := testutil.MakeTestDir(t, filepath.Join("..", "..", "testdata", "Namespace"))
	result, diags, err := ConvertFile(filepath.Join(testDir, "NamespaceWithTrailingComment.yaml"))
	if diags != nil {
		assertion.False(diags.HasErrors(), "diagnostics incorrectly displayed for proper yaml")
	}
	assertion.NoError(err)
	assertion.Equal(expected, result, "Comments are converted incorrectly")
}

func TestIncorrectPath(t *testing.T) {
	assertion := assert.New(t)
	fakePath := "fakePath"
	// Assert fakePath does not exist
	_, err := os.Stat(fakePath)
	assertion.True(os.IsNotExist(err))
	_, _, err = ConvertFile(fakePath)
	assertion.Error(err)
	_, _, err = ConvertDirectory(fakePath)
	assertion.Error(err)
}

func TestMalformedHeaderYaml(t *testing.T) {
	assertion := assert.New(t)
	testDir := testutil.MakeTestDir(t, filepath.Join("..", "..", "testdata", "MalformedYaml"))
	_, diags, err := ConvertFile(filepath.Join(testDir, "MalformedYaml.yaml"))
	if diags != nil {
		assertion.True(diags.HasErrors(), "diagnostics incorrectly displayed for wrongly formatted yaml")
	}
	assertion.NoError(err)
}

func TestMultipleResourceGen(t *testing.T) {
	assertion := assert.New(t)

	testDir := testutil.MakeTestDir(t, filepath.Join("..", "..", "testdata", "MultipleResources"))
	b, err := os.ReadFile(filepath.Join(testDir, "MultipleResources.pcl"))
	assertion.NoError(err)
	expected := string(b)

	result, diags, err := ConvertFile(filepath.Join(testDir, "MultipleResources.yml"))
	if diags != nil {
		assertion.False(diags.HasErrors(), "diagnostics incorrectly displayed for proper yaml")
	}
	assertion.NoError(err)
	assertion.Equal(expected, result, "File with multiple resources is converted incorrectly")
}

func TestEmptyDir(t *testing.T) {
	assertion := assert.New(t)
	testDir := testutil.MakeTestDir(t, filepath.Join("..", "..", "testdata", "empty"))
	_, diags, err := ConvertDirectory(testDir)
	if diags != nil {
		assertion.False(diags.HasErrors(), "diagnostics incorrectly displayed for proper yaml")
	}
	assertion.Error(err)
	assertion.Contains(err.Error(), "unable to find any YAML files")
}

func TestAnnotationsDeployment(t *testing.T) {
	assertion := assert.New(t)

	testDir := testutil.MakeTestDir(t, filepath.Join("..", "..", "testdata", "testDep"))
	b, err := os.ReadFile(filepath.Join(testDir, "testDep.pp"))
	assertion.NoError(err)
	expected := string(b)

	result, diags, err := ConvertFile(filepath.Join(testDir, "testDep.yaml"))
	if diags != nil {
		assertion.False(diags.HasErrors(), "diagnostics incorrectly displayed for proper yaml")
	}
	assertion.Equal(expected, result, "pcl is incorrect")
}

func TestNoDoubleQuotes(t *testing.T) {
	assertion := assert.New(t)

	testDir := testutil.MakeTestDir(t, filepath.Join("..", "..", "testdata", "doubleQuotes"))
	b, err := os.ReadFile(filepath.Join(testDir, "doubleQuotes.pp"))
	assertion.NoError(err)
	expected := string(b)

	result, diags, err := ConvertFile(filepath.Join(testDir, "doubleQuotes.yaml"))
	if diags != nil {
		assertion.False(diags.HasErrors(), "diagnostics incorrectly displayed for proper yaml")
	}
	assertion.NoError(err)
	assertion.Equal(expected, result, "double quotes inserted")
}

func TestSpecialChar(t *testing.T) {
	assertion := assert.New(t)

	testDir := testutil.MakeTestDir(t, filepath.Join("..", "..", "testdata", "specialChar"))
	b, err := os.ReadFile(filepath.Join(testDir, "specialChar.pp"))
	assertion.NoError(err)
	expected := string(b)

	result, diags, err := ConvertFile(filepath.Join(testDir, "specialChar.yaml"))
	if diags != nil {
		assertion.False(diags.HasErrors(), "diagnostics incorrectly displayed for proper yaml")
	}
	assertion.NoError(err)
	assertion.Equal(expected, result, "double quotes inserted")
}

func TestMultiLineString(t *testing.T) {
	assertion := assert.New(t)

	testDir := testutil.MakeTestDir(t, filepath.Join("..", "..", "testdata", "MultilineString"))
	b, err := os.ReadFile(filepath.Join(testDir, "MultilineString.pp"))
	assertion.NoError(err)
	expected := string(b)

	result, diags, err := ConvertFile(filepath.Join(testDir, "MultilineString.yaml"))
	if diags != nil {
		assertion.False(diags.HasErrors(), "diagnostics incorrectly displayed for proper yaml")
	}
	assertion.NoError(err)
	assertion.Equal(expected, result, "incorrectly parses multiline strings")
}

func TestCRD(t *testing.T) {
	assertion := assert.New(t)
	testDir := testutil.MakeTestDir(t, filepath.Join("..", "..", "testdata", "crd"))
	_, diags, err := ConvertFile(filepath.Join(testDir, "customResourceDef.yaml"))
	if diags != nil {
		assertion.True(diags.HasErrors(), "diagnostics not detecting CRD")
	}
	assertion.NoError(err)
}

func TestNotYaml(t *testing.T) {
	assertion := assert.New(t)
	testDir := testutil.MakeTestDir(t, filepath.Join("..", "..", "testdata", "empty"))
	_, _, err := ConvertFile(filepath.Join(testDir, "notYAML.txt"))
	assertion.Error(err)
}

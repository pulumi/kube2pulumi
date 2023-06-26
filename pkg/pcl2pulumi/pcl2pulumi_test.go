package pcl2pulumi

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/pulumi/kube2pulumi/pkg/testutil"
	"github.com/stretchr/testify/assert"
)

// GENERALIZED TESTS

func TestDoubleQuotes(t *testing.T) {
	for language := range testutil.Languages() {
		language := language
		t.Run(language, func(t *testing.T) {
			t.Parallel()
			assertion := assert.New(t)
			testDir := testutil.MakeTestDir(t, filepath.Join("..", "..", "testdata", "doubleQuotes"))
			pcl, err := os.ReadFile(filepath.Join(testDir, "doubleQuotes.pp"))
			assertion.NoError(err)

			_, err = Pcl2Pulumi(string(pcl), filepath.Join(testDir, "doubleQuotes"), language)
			assertion.NoError(err)
		})
	}
}

func TestSpecialChar(t *testing.T) {
	for language := range testutil.Languages() {
		language := language
		t.Run(language, func(t *testing.T) {
			t.Parallel()
			assertion := assert.New(t)
			testDir := testutil.MakeTestDir(t, filepath.Join("..", "..", "testdata", "specialChar"))
			pcl, err := os.ReadFile(filepath.Join(testDir, "specialChar.pp"))
			assertion.NoError(err)

			_, err = Pcl2Pulumi(string(pcl), filepath.Join(testDir, "specialChar"), language)
			assertion.NoError(err)
		})
	}
}

func TestAnnotations(t *testing.T) {
	for language := range testutil.Languages() {
		language := language
		t.Run(language, func(t *testing.T) {
			t.Parallel()
			assertion := assert.New(t)
			testDir := testutil.MakeTestDir(t, filepath.Join("..", "..", "testdata", "testDep"))
			pcl, err := os.ReadFile(filepath.Join(testDir, "testDep.pp"))
			assertion.NoError(err)

			_, err = Pcl2Pulumi(string(pcl), filepath.Join(testDir, "testDep"), language)
			assertion.NoError(err)
		})
	}
}

func TestNamespace(t *testing.T) {
	for language, ext := range testutil.Languages() {
		language, ext := language, ext
		t.Run(language, func(t *testing.T) {
			t.Parallel()
			assertion := assert.New(t)
			testDir := testutil.MakeTestDir(t, filepath.Join("..", "..", "testdata", "Namespace"))
			expected, err := os.ReadFile(filepath.Join(testDir, fmt.Sprintf("expectedNamespace%s", ext)))
			assertion.NoError(err)

			pcl, err := os.ReadFile(filepath.Join(testDir, "Namespace.pp"))
			assertion.NoError(err)

			outPath, err := Pcl2Pulumi(string(pcl), filepath.Join(testDir, "Namespace"), language)
			assertion.NoError(err)

			generated, err := os.ReadFile(outPath)
			assertion.NoError(err)

			assertion.Equal(string(expected), string(generated), fmt.Sprintf("%s codegen is incorrect", language))
		})
	}
}

func TestOperator(t *testing.T) {
	for language, ext := range testutil.Languages() {
		language, ext := language, ext
		t.Run(language, func(t *testing.T) {
			t.Parallel()
			assertion := assert.New(t)
			testDir := testutil.MakeTestDir(t, filepath.Join("..", "..", "testdata", "k8sOperator"))
			expected, err := os.ReadFile(filepath.Join(testDir, fmt.Sprintf("expectedMain%s", ext)))
			assertion.NoError(err)

			pcl, err := os.ReadFile(filepath.Join(testDir, "expK8sOperator.pp"))
			assertion.NoError(err)

			outPath, err := Pcl2Pulumi(string(pcl), filepath.Join(testDir, "main"), language)
			assertion.NoError(err)

			generated, err := os.ReadFile(outPath)
			assertion.NoError(err)

			assertion.Equal(string(expected), string(generated), fmt.Sprintf("%s codegen is incorrect", language))
		})
	}
}

func TestMultiLineString(t *testing.T) {
	for language, ext := range testutil.Languages() {
		language, ext := language, ext
		t.Run(language, func(t *testing.T) {
			t.Parallel()
			assertion := assert.New(t)
			testDir := testutil.MakeTestDir(t, filepath.Join("..", "..", "testdata", "MultilineString"))
			expected, err := os.ReadFile(filepath.Join(testDir, fmt.Sprintf("expectedMultilineString%s", ext)))
			assertion.NoError(err)

			pcl, err := os.ReadFile(filepath.Join(testDir, "MultilineString.pp"))
			assertion.NoError(err)

			outPath, err := Pcl2Pulumi(string(pcl), testDir, language)
			assertion.NoError(err)

			generated, err := os.ReadFile(outPath)
			assertion.NoError(err)

			assertion.Equal(string(expected), string(generated), fmt.Sprintf("%s codegen is incorrect", language))
		})
	}
}

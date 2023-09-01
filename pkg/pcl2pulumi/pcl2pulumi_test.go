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
			testdataDir := filepath.Join("..", "..", "testdata", "Namespace")
			testDir := testutil.MakeTestDir(t, testdataDir)
			expected := filepath.Join(testdataDir, "expected", fmt.Sprintf("expectedNamespace%s", ext))

			pcl, err := os.ReadFile(filepath.Join(testDir, "Namespace.pp"))
			assertion.NoError(err)

			outPath, err := Pcl2Pulumi(string(pcl), filepath.Join(testDir, "Namespace"), language)
			assertion.NoError(err)

			testutil.AssertFilesEqual(t, expected, outPath)
		})
	}
}

func TestOperator(t *testing.T) {
	for language, ext := range testutil.Languages() {
		language, ext := language, ext
		t.Run(language, func(t *testing.T) {
			t.Parallel()
			assertion := assert.New(t)
			testdataDir := filepath.Join("..", "..", "testdata", "k8sOperator")
			testDir := testutil.MakeTestDir(t, testdataDir)
			expected := filepath.Join(testdataDir, "expected", fmt.Sprintf("expectedMain%s", ext))

			pcl, err := os.ReadFile(filepath.Join(testDir, "expected", "expectedK8sOperator.pp"))
			assertion.NoError(err)

			outPath, err := Pcl2Pulumi(string(pcl), filepath.Join(testDir, "main"), language)
			assertion.NoError(err)

			testutil.AssertFilesEqual(t, expected, outPath)
		})
	}
}

func TestMultiLineString(t *testing.T) {
	for language, ext := range testutil.Languages() {
		language, ext := language, ext
		t.Run(language, func(t *testing.T) {
			t.Parallel()
			assertion := assert.New(t)
			testdataDir := filepath.Join("..", "..", "testdata", "MultilineString")
			testDir := testutil.MakeTestDir(t, testdataDir)
			expectedPath := filepath.Join(testdataDir, "expected",
				fmt.Sprintf("expectedMultilineString%s", ext))

			pcl, err := os.ReadFile(filepath.Join(testDir, "MultilineString.pp"))
			assertion.NoError(err)

			outPath, err := Pcl2Pulumi(string(pcl), testDir, language)
			assertion.NoError(err)

			testutil.AssertFilesEqual(t, expectedPath, outPath)
		})
	}
}

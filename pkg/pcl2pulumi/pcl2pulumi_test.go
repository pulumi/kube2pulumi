package pcl2pulumi

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getLangs() map[string]string {
	return map[string]string{
		"python":     ".py",
		"typescript": ".ts",
		"csharp":     ".cs",
		"go":         ".go",
	}
}

// GENERALIZED TESTS

func TestDoubleQuotes(t *testing.T) {
	assertion := assert.New(t)
	langs := getLangs()

	for language := range langs {
		pcl, err := ioutil.ReadFile(filepath.Join("..", "..", "testdata", "doubleQuotes.pp"))
		assertion.NoError(err)

		_, err = Pcl2Pulumi(string(pcl), filepath.Join("..", "..", "testdata", "doubleQuotes"), language)
		assertion.NoError(err)
	}
}

func TestSpecialChar(t *testing.T) {
	assertion := assert.New(t)
	langs := getLangs()
	for language := range langs {
		pcl, err := ioutil.ReadFile(filepath.Join("..", "..", "testdata", "specialChar.pp"))
		assertion.NoError(err)

		_, err = Pcl2Pulumi(string(pcl), filepath.Join("..", "..", "testdata", "specialChar"), language)
		assertion.NoError(err)
	}
}

func TestAnnotations(t *testing.T) {
	assertion := assert.New(t)
	langs := getLangs()
	for language := range langs {
		pcl, err := ioutil.ReadFile(filepath.Join("..", "..", "testdata", "testDep.pp"))
		assertion.NoError(err)

		_, err = Pcl2Pulumi(string(pcl), filepath.Join("..", "..", "testdata", "testDep"), language)
		assertion.NoError(err)
	}
}

func TestNamespace(t *testing.T) {
	assertion := assert.New(t)
	langs := getLangs()

	for language, ext := range langs {
		expected, err := ioutil.ReadFile(filepath.Join("..", "..", "testdata",
			"expNamespace", fmt.Sprintf("expectedNamespace%s", ext)))
		assertion.NoError(err)

		pcl, err := ioutil.ReadFile(filepath.Join("..", "..", "testdata", "Namespace.pp"))
		assertion.NoError(err)

		outPath, err := Pcl2Pulumi(string(pcl), filepath.Join("..", "..", "testdata", "Namespace"), language)
		assertion.NoError(err)

		generated, err := ioutil.ReadFile(outPath)
		assertion.NoError(err)

		assertion.Equal(string(expected), string(generated), fmt.Sprintf("%s codegen is incorrect", language))
	}
}

func TestOperator(t *testing.T) {
	assertion := assert.New(t)
	langs := getLangs()

	for language, ext := range langs {
		expected, err := ioutil.ReadFile(filepath.Join("..", "..", "testdata",
			"k8sOperator", fmt.Sprintf("expectedMain%s", ext)))
		assertion.NoError(err)

		pcl, err := ioutil.ReadFile(filepath.Join("..", "..", "testdata", "expK8sOperator.pp"))
		assertion.NoError(err)

		outPath, err := Pcl2Pulumi(string(pcl), filepath.Join("..", "..", "testdata", "k8sOperator", "main"), language)
		assertion.NoError(err)

		generated, err := ioutil.ReadFile(outPath)
		assertion.NoError(err)

		assertion.Equal(string(expected), string(generated), fmt.Sprintf("%s codegen is incorrect", language))
	}
}

func TestMultiLineString(t *testing.T) {
	assertion := assert.New(t)
	langs := getLangs()

	for language, ext := range langs {
		expected, err := ioutil.ReadFile(filepath.Join("..", "..", "testdata",
			fmt.Sprintf("expectedMultilineString%s", ext)))
		assertion.NoError(err)

		pcl, err := ioutil.ReadFile(filepath.Join("..", "..", "testdata", "MultilineString.pp"))
		assertion.NoError(err)

		outPath, err := Pcl2Pulumi(string(pcl), filepath.Join("..", "..", "testdata", "MultilineString"), language)
		assertion.NoError(err)

		generated, err := ioutil.ReadFile(outPath)
		assertion.NoError(err)

		assertion.Equal(string(expected), string(generated), fmt.Sprintf("%s codegen is incorrect", language))
	}
}

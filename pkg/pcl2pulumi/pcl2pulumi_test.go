package pcl2pulumi

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getLangs() map[string]string {
	return map[string]string{
		"python":     ".py",
		"typescript": ".ts",
		"csharp":     ".cs",
		"java":       ".java",
		"go":         ".go",
	}
}

// GENERALIZED TESTS

func TestDoubleQuotes(t *testing.T) {
	assertion := assert.New(t)
	langs := getLangs()

	for language := range langs {
		pcl, err := os.ReadFile(filepath.Join("..", "..", "testdata", "doubleQuotes", "doubleQuotes.pp"))
		assertion.NoError(err)

		_, err = Pcl2Pulumi(string(pcl), filepath.Join("..", "..", "testdata", "doubleQuotes", "doubleQuotes"), language)
		assertion.NoError(err)
	}
}

func TestSpecialChar(t *testing.T) {
	assertion := assert.New(t)
	langs := getLangs()
	for language := range langs {
		if language == "go" {
			// will be able to run in tests when https://github.com/pulumi/pulumi/issues/8940 is complete
			continue
		}
		pcl, err := os.ReadFile(filepath.Join("..", "..", "testdata", "specialChar", "specialChar.pp"))
		assertion.NoError(err)

		_, err = Pcl2Pulumi(string(pcl), filepath.Join("..", "..", "testdata", "specialChar", "specialChar"), language)
		assertion.NoError(err)
	}
}

func TestAnnotations(t *testing.T) {
	assertion := assert.New(t)
	langs := getLangs()
	for language := range langs {
		pcl, err := os.ReadFile(filepath.Join("..", "..", "testdata", "testDep", "testDep.pp"))
		assertion.NoError(err)

		_, err = Pcl2Pulumi(string(pcl), filepath.Join("..", "..", "testdata", "testDep", "testDep"), language)
		assertion.NoError(err)
	}
}

func TestNamespace(t *testing.T) {
	assertion := assert.New(t)
	langs := getLangs()

	for language, ext := range langs {
		expected, err := os.ReadFile(filepath.Join("..", "..", "testdata",
			"Namespace", fmt.Sprintf("expectedNamespace%s", ext)))
		assertion.NoError(err)

		pcl, err := os.ReadFile(filepath.Join("..", "..", "testdata", "Namespace", "Namespace.pp"))
		assertion.NoError(err)

		outPath, err := Pcl2Pulumi(string(pcl), filepath.Join("..", "..", "testdata", "Namespace", "Namespace"), language)
		assertion.NoError(err)

		generated, err := os.ReadFile(outPath)
		assertion.NoError(err)

		assertion.Equal(string(expected), string(generated), fmt.Sprintf("%s codegen is incorrect", language))
	}
}

func TestOperator(t *testing.T) {
	assertion := assert.New(t)
	langs := getLangs()

	for language, ext := range langs {
		expected, err := os.ReadFile(filepath.Join("..", "..", "testdata",
			"k8sOperator", fmt.Sprintf("expectedMain%s", ext)))
		assertion.NoError(err)

		pcl, err := os.ReadFile(filepath.Join("..", "..", "testdata", "k8sOperator", "expK8sOperator.pp"))
		assertion.NoError(err)

		outPath, err := Pcl2Pulumi(string(pcl), filepath.Join("..", "..", "testdata", "k8sOperator", "main"), language)
		assertion.NoError(err)

		generated, err := os.ReadFile(outPath)
		assertion.NoError(err)

		assertion.Equal(string(expected), string(generated), fmt.Sprintf("%s codegen is incorrect", language))
	}
}

func TestMultiLineString(t *testing.T) {
	assertion := assert.New(t)
	langs := getLangs()

	for language, ext := range langs {
		expected, err := os.ReadFile(filepath.Join("..", "..", "testdata", "MultilineString",
			fmt.Sprintf("expectedMultilineString%s", ext)))
		assertion.NoError(err)

		pcl, err := os.ReadFile(filepath.Join("..", "..", "testdata", "MultilineString", "MultilineString.pp"))
		assertion.NoError(err)

		outPath, err := Pcl2Pulumi(string(pcl), filepath.Join("..", "..", "testdata", "MultilineString"), language)
		assertion.NoError(err)

		generated, err := os.ReadFile(outPath)
		assertion.NoError(err)

		assertion.Equal(string(expected), string(generated), fmt.Sprintf("%s codegen is incorrect", language))
	}
}

package testutil

import (
	"os"
	"testing"

	"github.com/pulumi/pulumi/sdk/v3/go/common/util/cmdutil"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/fsutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var pulumiAccept = cmdutil.IsTruthy(os.Getenv("PULUMI_ACCEPT"))

// Languages returns a map of supported languages for conversion, and their default extensions.
func Languages() map[string]string {
	return map[string]string{
		"python":     ".py",
		"typescript": ".ts",
		"csharp":     ".cs",
		"java":       ".java",
		"go":         ".go",
	}
}

// MakeTestDir creates a temporary test folder and copies the testdata into it.
// This ensures that generated test files do not pollute the working directory, and cause
// developers to commit erronerous files by mistake. We can also rely on the temporary folder
// being cleaned up by the Golang test framework.
func MakeTestDir(t *testing.T, sourceDir string) string {
	t.Helper()
	tmpDir := t.TempDir()
	err := fsutil.CopyFile(tmpDir, sourceDir, nil)
	if err != nil {
		t.Fatalf("Unable to copy testdata %q to temporary directory %q", sourceDir, tmpDir)
	}

	return tmpDir
}

// Assert that the contents of two files are equal.
//
// To update the expected file, run the test with PULUMI_ACCEPT=true.
func AssertFilesEqual(t *testing.T, expectedPath, generatedPath string) {
	generated, err := os.ReadFile(generatedPath)
	require.NoError(t, err)

	if pulumiAccept {
		err := os.WriteFile(expectedPath, generated, 0700)
		assert.NoError(t, err)
	} else {
		expected, err := os.ReadFile(expectedPath)
		require.NoError(t, err)
		assert.Equal(t, string(expected), string(generated), "codegen is incorrect")
	}

}

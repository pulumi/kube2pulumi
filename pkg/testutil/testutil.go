package testutil

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

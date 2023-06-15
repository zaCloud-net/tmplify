package tmplify

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func findImportStatements(template string) []string {
	importPattern := regexp.MustCompile(`{{import\s*\*\s*as\s*\w+\s*from\s*"([^"]+)"}}`)
	return importPattern.FindAllString(template, -1)
}

func extractImportPath(importStmt string) (string, string) {
	importPattern := regexp.MustCompile(`{{import\s*\*\s*as\s*(\w+)\s*from\s*"([^"]+)"}}`)
	matches := importPattern.FindStringSubmatch(importStmt)
	if len(matches) >= 3 {
		moduleName := matches[1]
		importPath := matches[2]
		return moduleName, importPath
	}
	return "", ""
}

func removeImportLines(input string) string {
	importPattern := regexp.MustCompile(`(?m).*{{import.*}}.*\n?`)
	result := importPattern.ReplaceAllString(input, "")
	result = strings.TrimSpace(result)
	return result
}

func importFile(importPath string) (string, error) {
	absPath, err := resolveImportPath(importPath)
	if err != nil {
		return "", err
	}

	data, err := os.ReadFile(absPath)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func resolveImportPath(importPath string) (string, error) {
	baseDir, _ := os.Getwd()
	absPath := filepath.Join(baseDir, importPath)

	return absPath, nil
}

func handleImports(template string, data map[string]interface{}) (string, map[string]interface{}, error) {
	imports := findImportStatements(template)
	for _, importStmt := range imports {
		moduleName, importPath := extractImportPath(importStmt)
		importedData, err := importFile(importPath)
		if err != nil {
			return "", data, err
		}
		data[moduleName] = importedData
	}

	template = removeImportLines(template)

	return template, data, nil
}

package tmplify

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
)

func getHelpers(t *template.Template) template.FuncMap {
	return template.FuncMap{
		"equal": func(a, b interface{}) bool {
			return a == b
		},
		"notEqual": func(a, b interface{}) bool {
			return a != b
		},
		"add": func(a, b int) int {
			return a + b
		},
		"subtract": func(a, b int) int {
			return a - b
		},
		"multiply": func(a, b int) int {
			return a * b
		},
		"divide": func(a, b int) int {
			return a / b
		},
		"uppercase": func(s string) string {
			return strings.ToUpper(s)
		},
		"lowercase": func(s string) string {
			return strings.ToLower(s)
		},
		"lessThan": func(a, b interface{}) bool {
			switch a.(type) {
			case int:
				return a.(int) < b.(int)
			case float64:
				return a.(float64) < b.(float64)
			}
			return false
		},
		"greaterThan": func(a, b interface{}) bool {
			switch a.(type) {
			case int:
				return a.(int) > b.(int)
			case float64:
				return a.(float64) > b.(float64)
			}
			return false
		},
		"contains": func(arr []interface{}, item interface{}) bool {
			for _, value := range arr {
				if value == item {
					return true
				}
			}
			return false
		},
		"substring": func(s string, start, length int) string {
			if start < 0 || start >= len(s) || length <= 0 || start+length > len(s) {
				return ""
			}
			return s[start : start+length]
		},
		"indent": func(spaces int, v string) string {
			return indent(spaces, v)
		},
		"include": func(name string) (string, error) {
			return include(name)
		},
		"includeI": func(name string, spaces int) (string, error) {
			str, err := include(name)
			return indent(spaces, str), err
		},
	}
}

func indent(spaces int, v string) string {
	arr := strings.Split(v, "\n")
	for i, str := range arr {
		if i > 0 {
			pad := strings.Repeat(" ", spaces)
			arr[i] = pad + strings.Replace(str, "\n", "\n"+pad, -1)
		}
	}
	return strings.Join(arr, "\n")
}

func handleIncludeName(name string) (string, string) {
	arr := strings.Split(name, "/")
	if len(arr) == 1 {
		return name, ""
	} else {
		path, _ := resolveImportPath(name)
		fmt.Println(path)
		return arr[len(arr)-1], path
	}
}

func include(name string) (string, error) {
	buf := bytes.NewBuffer(nil)
	name, path := handleIncludeName(name)
	if path != "" {
		t.ParseFiles(path)
	}
	if err := t.ExecuteTemplate(buf, name, nil); err != nil {
		return "", err
	}
	return buf.String(), nil
}
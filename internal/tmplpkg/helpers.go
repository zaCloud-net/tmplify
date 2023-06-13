package tmplpkg

import (
	"strings"
	"text/template"
)

var FuncMap = template.FuncMap{
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
}

package tmplpkg

import (
	"strings"
	"text/template"
)

func TemplateString(templateString string, values map[string]interface{}) (string, error) {
	tmpl, err := template.New("template").Funcs(FuncMap).Parse(templateString)
	if err != nil {
		return "", err
	}

	var result string
	buffer := &strings.Builder{}

	err = tmpl.Execute(buffer, values)
	if err != nil {
		return "", err
	}

	result = buffer.String()
	return result, nil
}

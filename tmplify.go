package tmplify

import (
	"bytes"
	"text/template"
)

var t = template.New("base")

func TemplateString(templateString string, values map[string]interface{}) (string, error) {
	tmplString, data, err := handleImports(templateString, values)
	if err != nil {
		return "", err
	}

	tmpl, err := t.Funcs(getHelpers(t, values)).Parse(tmplString)
	if err != nil {
		return "", err
	}

	var result string
	buffer := bytes.NewBuffer(nil)

	err = tmpl.Execute(buffer, data)
	if err != nil {
		return "", err
	}

	result = removeEmptyLines(buffer.String())

	return result, nil
}

func ParseTemplates(glob string) error {
	_, err := t.ParseGlob(glob)
	return err
}

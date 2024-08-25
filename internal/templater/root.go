package templater

import (
	"bytes"
	_ "embed"
	"text/template"

	"github.com/Masterminds/sprig"
)

const (
	serviceTemplateName = "ServiceTemplate"
	fileTemplateName    = "FileTemplate"
)

//go:embed templates/service.template
var serviceTemplate string

//go:embed templates/file.template
var fileTemplate string

func GenerateFile(data File) ([]byte, error) {
	return generateContent(data,
		additionalTemplate{templateName: fileTemplateName, templateContent: fileTemplate},
		additionalTemplate{templateName: serviceTemplateName, templateContent: serviceTemplate},
	)
}

func generateContent(data interface{}, templates ...additionalTemplate) ([]byte, error) {
	tp := template.New(templates[0].templateName)
	tp.Funcs(sprig.FuncMap()).Funcs(newIncludeFuncMap(tp))
	_, err := tp.Parse(templates[0].templateContent)
	if err != nil {
		return nil, err
	}
	for _, optTemplate := range templates[1:] {
		_, err := tp.New(optTemplate.templateName).Parse(optTemplate.templateContent)
		if err != nil {
			return nil, err
		}
	}

	buf := &bytes.Buffer{}
	err = tp.Execute(buf, data)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

type additionalTemplate struct {
	templateName    string
	templateContent string
}

func newIncludeFuncMap(tp *template.Template) template.FuncMap {
	return template.FuncMap{
		"include": newIncludeFn(tp),
	}
}

func newIncludeFn(tp *template.Template) func(name string, data interface{}) (string, error) {
	return func(name string, data interface{}) (string, error) {
		buf := bytes.NewBuffer(nil)
		if err := tp.ExecuteTemplate(buf, name, data); err != nil {
			return "", err
		}
		return buf.String(), nil
	}
}

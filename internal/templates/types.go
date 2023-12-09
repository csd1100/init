package templates

import (
	"log"
	"os"
	"text/template"
)

type Project interface {
	Init() error
	ParseTemplates() error
}

type TemplateFile struct {
	Src string
	Dst string
}

type Template struct {
	Name          string
	TemplateFiles []TemplateFile
	TemplateData  map[string]string
}

func (template Template) Init() error {
	return nil
}

func (tmpl Template) ParseTemplates() error {
	for _, templateFile := range tmpl.TemplateFiles {
		parsedTemplate, err := template.ParseFiles(templateFile.Src)
		if err != nil {
			return err
		}

		file, err := os.Create(templateFile.Dst)
		if err != nil {
			return err
		}

		parsedTemplate.Execute(file, tmpl.TemplateData)
		parsedTemplate.Execute(log.Writer(), tmpl.TemplateData)

	}

	return nil
}

func GetTemplate(name string, projectName string) (Project, error) {
	switch name {
	case "go":
		return generateGoTemplate(projectName), nil
	case "js":
		return generateJSTemplate(projectName), nil
	default:
		return nil, ErrInvalidArgTemplate
	}
}

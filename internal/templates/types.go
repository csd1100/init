package templates

import (
	"os"
	"text/template"

	"github.com/csd1100/init/internal/cli"
)

type Project interface {
	Sync(map[string]string) error
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
	BuildTool     cli.BuildTool
}

func (template Template) Sync(data map[string]string) error {
	return template.BuildTool.Sync(data)
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

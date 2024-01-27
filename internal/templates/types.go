package templates

import (
	"os"
	"strings"
	"text/template"

	"github.com/csd1100/init/internal/cli"
	"github.com/csd1100/init/internal/helpers"
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

func GetTemplate(templateName string, projectName string, stringOptions string) (Project, error) {
	templateOptions := make(map[string]string)

	if stringOptions != "" {
		templateOptions = getTemplateOptions(stringOptions)
	}

	templateOptions[helpers.PROJECT_NAME] = projectName

	switch templateName {
	case "go":
		return generateGoTemplate(templateOptions), nil
	case "js":
		return generateJSTemplate(templateOptions), nil
	default:
		return nil, helpers.ErrInvalidArgTemplate
	}
}

func getTemplateOptions(templateOptions string) map[string]string {
	templOptions := make(map[string]string)
	for _, option := range strings.Split(templateOptions, ",") {
		split := strings.Split(option, "=")
		templOptions[split[0]] = split[1]
	}
	return templOptions
}

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
	Template string `json:"template"`
	Real     string `json:"real"`
}

type TemplateFiles struct {
	Files []TemplateFile `json:"template-files"`
}

type Template struct {
	Name          string
	TemplateFiles TemplateFiles
	TemplateData  map[string]string
	BuildTool     cli.BuildTool
}

func (template *Template) Sync(data map[string]string) error {
	return template.BuildTool.Sync(data)
}

func (tmpl *Template) ParseTemplates() error {
	for _, templateFile := range tmpl.TemplateFiles.Files {
		parsedTemplate, err := template.ParseFiles(templateFile.Template)
		if err != nil {
			return err
		}

		file, err := os.Create(templateFile.Real)
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
		if templateOptions[helpers.GO_PACKAGE_NAME] == "" {
			templateOptions[helpers.GO_PACKAGE_NAME] = "project"
		}
		return &Template{
			Name:         "go",
			TemplateData: templateOptions,
			BuildTool:    cli.Go,
		}, nil
	case "js":
		return &Template{
			Name:         "js",
			TemplateData: templateOptions,
			BuildTool:    cli.Npm,
		}, nil
	case "rust":
		templateOptions[helpers.PROJECT_NAME_WITH_REPLACED_HYPHENS] = strings.ReplaceAll(projectName, "-", "_")
		return &Template{
			Name:         "rust",
			TemplateData: templateOptions,
			BuildTool:    cli.Cargo,
		}, nil
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

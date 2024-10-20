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

func (tmpl *Template) Sync(data map[string]string) error {
	return tmpl.BuildTool.Sync(data)
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

		err = parsedTemplate.Execute(file, tmpl.TemplateData)
		if err != nil {
			return err
		}
	}

	return nil
}

func GetTemplate(templateName string, projectName string, stringOptions string) (Project, error) {
	templateOptions := make(map[string]string)

	if stringOptions != "" {
		templateOptions = getTemplateOptions(stringOptions)
	}

	templateOptions[helpers.ProjectName] = projectName

	switch templateName {
	case "go":
		if templateOptions[helpers.GoPackageName] == "" {
			templateOptions[helpers.GoPackageName] = "project"
		}
		return &Template{
			Name:         "go",
			TemplateData: templateOptions,
			BuildTool:    cli.Go,
		}, nil
	case "js", "ts", "electron-react-ts":
		return &Template{
			Name:         templateName,
			TemplateData: templateOptions,
			BuildTool:    cli.Npm,
		}, nil
	case "rust":
		templateOptions[helpers.ProjectNameWithReplacedHyphens] = strings.ReplaceAll(projectName, "-", "_")
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
	tmplOptions := make(map[string]string)
	for _, option := range strings.Split(templateOptions, ",") {
		split := strings.Split(option, "=")
		tmplOptions[split[0]] = split[1]
	}
	return tmplOptions
}

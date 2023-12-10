package templates

import (
	"github.com/csd1100/init/internal/cli"
	"github.com/csd1100/init/internal/helpers"
)

func generateGoTemplate(templateOptions map[string]string) Project {

	var goTemplateFiles = []TemplateFile{
		{
			Src: "./templates/main.go.tmpl",
			Dst: "./main.go",
		},
	}

	if templateOptions[helpers.GO_PACKAGE_NAME] != "" {
		templateOptions[helpers.PROJECT_NAME] = templateOptions[helpers.GO_PACKAGE_NAME] + "/" +
			templateOptions[helpers.PROJECT_NAME]

	}

	return Template{
		Name:          "go",
		TemplateFiles: goTemplateFiles,
		TemplateData:  templateOptions,
		BuildTool:     cli.Go,
	}
}

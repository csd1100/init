package templates

import (
	"github.com/csd1100/init/internal/cli"
)

func generateJSTemplate(templateOptions map[string]string) Project {
	var jSTemplateFiles = []TemplateFile{
		{
			Src: "./templates/package.json.tmpl",
			Dst: "./package.json",
		},
	}

	return Template{
		Name:          "js",
		TemplateFiles: jSTemplateFiles,
		TemplateData:  templateOptions,
		BuildTool:     cli.Npm,
	}
}

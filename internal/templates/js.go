package templates

import (
	"github.com/csd1100/init/internal/cli"
)

func generateJSTemplate(projectName string) Project {
	var jSTemplateFiles = []TemplateFile{
		{
			Src: "./templates/package.json.tmpl",
			Dst: "./package.json",
		},
	}

	var templateData = make(map[string]string)
	templateData["projectName"] = projectName

	return Template{
		Name:          "js",
		TemplateFiles: jSTemplateFiles,
		TemplateData:  templateData,
		BuildTool:     cli.Npm,
	}
}

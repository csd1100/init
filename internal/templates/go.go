package templates

import "github.com/csd1100/init/internal/cli"

func generateGoTemplate(projectName string) Project {

	var goTemplateFiles = []TemplateFile{
		{
			Src: "./templates/go.mod.tmpl",
			Dst: "./go.mod",
		},
		{
			Src: "./templates/main.go.tmpl",
			Dst: "./main.go",
		},
	}

	var templateData = make(map[string]string)
	templateData["projectName"] = projectName

	return Template{
		Name:          "go",
		TemplateFiles: goTemplateFiles,
		TemplateData:  templateData,
		BuildTool:     cli.Go,
	}
}

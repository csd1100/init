package templates

type goTemplate struct {
	Template
}

func (goTemp goTemplate) Init() error {
	return nil
}

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
	}
}

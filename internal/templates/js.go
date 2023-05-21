package templates

type jsTemplate struct {
	Template
}

func (jsTemp jsTemplate) Init() error {
	return nil
}

func generateJSTemplate(projectName string) Project {
	var jSTemplateFiles = []TemplateFile{
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
		Name:          "js",
		TemplateFiles: jSTemplateFiles,
		TemplateData:  templateData,
	}
}

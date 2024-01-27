package templates

import (
	"github.com/csd1100/init/internal/cli"
)

func generateRustTemplate(templateOptions map[string]string) Project {
	var rustTemplateFiles = []TemplateFile{
		{
			Src: "./templates/Cargo.toml.tmpl",
			Dst: "./Cargo.toml",
		},
		{
			Src: "./templates/src/main.rs.tmpl",
			Dst: "./src/main.rs",
		},
		{
			Src: "./templates/src/bin/math.rs.tmpl",
			Dst: "./src/bin/math.rs",
		},
	}

	return Template{
		Name:          "rust",
		TemplateFiles: rustTemplateFiles,
		TemplateData:  templateOptions,
		BuildTool:     cli.Cargo,
	}
}

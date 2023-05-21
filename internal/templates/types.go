package templates

import (
	"os"
	"text/template"
)

type Project interface {
	Init() error
	ParseTemplates() error
}

type TemplateFile struct {
	Src string
	Dst string
}

type TemplateParser struct {
	accumulator []byte
}

func (tp *TemplateParser) Write(p []byte) (int, error) {
	tp.accumulator = append(tp.accumulator, p...)
	return len(p), nil
}

type Template struct {
	Name          string
	TemplateFiles []TemplateFile
	TemplateData  map[string]string
}

func (template Template) Init() error {
	return nil
}

func writeToFile(file string, data []byte) error {
	dstFile, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY, 0644)
	defer dstFile.Close()
	if err != nil {
		return err
	}

	_, err = dstFile.Write(data)
	if err != nil {
		return err
	}

	return nil
}

func (tmpl Template) ParseTemplates() error {
	for _, templateFile := range tmpl.TemplateFiles {
		parsedTemplate, err := template.ParseFiles(templateFile.Src)
		if err != nil {
			return err
		}

		parser := &TemplateParser{}
		parsedTemplate.Execute(parser, tmpl.TemplateData)

		err = writeToFile(templateFile.Dst, parser.accumulator)
		if err != nil {
			return err
		}
	}

	return nil
}

func GetTemplate(name string, projectName string) (Project, error) {
	switch name {
	case "go":
		return generateGoTemplate(projectName), nil
	case "js":
		return generateJSTemplate(projectName), nil
	default:
		return nil, ErrInvalidArgTemplate
	}
}

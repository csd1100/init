package templates

type Project interface {
	Init() error
	ParseTemplates() error
}

type Template struct {
	Name          string
	TemplateFiles []string
}

func (template Template) Init() error {
	return nil
}

func (template Template) ParseTemplates() error {
	return nil
}

func GetTemplate(name string) (Project, error) {
	switch name {
	case "go":
		return Go, nil
	case "js":
		return JS, nil
	default:
		return nil, ErrInvalidArgTemplate
	}
}

func newTemplate(name string, templateFiles []string) Template {
	return Template{
		Name:          name,
		TemplateFiles: templateFiles,
	}
}

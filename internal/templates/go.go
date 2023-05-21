package templates

type goTemplate struct {
	Template
}

func (goTemp goTemplate) Init() error {
	return nil
}

func (goTemp goTemplate) ParseTemplate() error {
	return nil
}

var Go = newTemplate("go", []string{})

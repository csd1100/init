package templates

type jsTemplate struct {
	Template
}

func (jsTemp jsTemplate) Init() error {
	return nil
}

func (jsTemp jsTemplate) ParseTemplate() error {
	return nil
}

var JS = newTemplate("js", []string{})

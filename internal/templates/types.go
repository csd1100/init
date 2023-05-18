package templates

import (
	"fmt"
)

func stringIsInArray(arr []string, toFind string) bool {
	for _, value := range arr {
		if value == toFind {
			return true
		}
	}
	return false
}

type Template struct {
	Name string
}

var supportedTemplates = []string{
	"go",
	"js",
}

func NewTemplate(name string) (*Template, error) {
	if stringIsInArray(supportedTemplates, name) {
		return &Template{
			Name: name,
		}, nil
	}
	return nil, fmt.Errorf("invalid template: %s", name)
}

package utils

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/csd1100/init/internal/templates"
)

func TestParse(t *testing.T) {
	cases := []struct {
		name           string
		args           []string
		expected_value *Options
		expected_error error
	}{
		{
			name:           "returns error if empty args",
			args:           []string{},
			expected_value: nil,
			expected_error: fmt.Errorf("expected at least 2 arguments: name and template"),
		},
		{
			name:           "returns error if name not included",
			args:           []string{"-t", "go"},
			expected_value: nil,
			expected_error: fmt.Errorf("expected argument: name"),
		},
		{
			name:           "returns error if template not included",
			args:           []string{"-n", "test"},
			expected_value: nil,
			expected_error: fmt.Errorf("expected argument: template"),
		},
		{
			name:           "returns error if invalid option",
			args:           []string{"-x", "test"},
			expected_value: nil,
			expected_error: fmt.Errorf("invalid argument: -x"),
		},
		{
			name:           "returns error if invalid template",
			args:           []string{"-n", "test", "-t", "test"},
			expected_value: nil,
			expected_error: fmt.Errorf("invalid template: test"),
		},
		{
			name: "returns options if only name and template",
			args: []string{"-n", "test", "-t", "go"},
			expected_value: &Options{
				Name:     "test",
				Template: templates.Template{Name: "go", TemplateFiles: []string{}},
			},
			expected_error: nil,
		},
		{
			name: "returns options if valid arguments",
			args: []string{"-n", "test", "-t", "go", "-G", "-S", "-p", "tmp/", "-h"},
			expected_value: &Options{
				Name:     "test",
				Template: templates.Template{Name: "go", TemplateFiles: []string{}},
				NoGit:    true,
				NoSync:   true,
				Path:     "tmp/",
				Help:     true,
			},
			expected_error: nil,
		},
		{
			name: "returns options if valid arguments with long version",
			args: []string{"--name", "test", "--template", "go", "--no-git", "--no-sync", "--path", "tmp/", "--help"},
			expected_value: &Options{
				Name:     "test",
				Template: templates.Template{Name: "go", TemplateFiles: []string{}},
				NoGit:    true,
				NoSync:   true,
				Path:     "tmp/",
				Help:     true,
			},
			expected_error: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := ParseArgs(tc.args)
			if err != nil {
				if err.Error() != tc.expected_error.Error() {
					t.Errorf(FAILURE_MESSAGE, tc.name, ERROR, tc.expected_error, err)
				}
				if actual != nil {
					t.Errorf(FAILURE_MESSAGE, tc.name, VALUE, tc.expected_value, actual)
				}
			} else {
				if !reflect.DeepEqual(*actual, *tc.expected_value) {
					t.Errorf(FAILURE_MESSAGE, tc.name, VALUE, *tc.expected_value, *actual)
				}
				if err != nil {
					t.Errorf(FAILURE_MESSAGE, tc.name, ERROR, tc.expected_error, err)
				}
			}
		})
	}
}

package utils

import (
	"fmt"
	"testing"

	"github.com/csd1100/init/internal/templates"
)

func TestParse(t *testing.T) {
	cases := []TestCase{
		{
			Name: "returns error if empty args",
			args: []any{
				[]string{},
			},
			expected_value: nil,
			expected_error: fmt.Errorf("expected at least 2 arguments: name and template"),
		},
		{
			Name: "returns error if name not included",
			args: []any{
				[]string{"-t", "go"},
			},
			expected_value: nil,
			expected_error: fmt.Errorf("expected argument: name"),
		},
		{
			Name: "returns error if template not included",
			args: []any{
				[]string{"-n", "test"},
			},
			expected_value: nil,
			expected_error: fmt.Errorf("expected argument: template"),
		},
		{
			Name: "returns error if invalid option",
			args: []any{
				[]string{"-x", "test"},
			},
			expected_value: nil,
			expected_error: fmt.Errorf("invalid argument: -x"),
		},
		{
			Name: "returns options if only name and template",
			args: []any{
				[]string{"-n", "test", "-t", "go"},
			},
			expected_value: Options{
				Name:     "test",
				Template: templates.Template{Name: "go"},
			},
			expected_error: nil,
		},
		{
			Name: "returns error if invalid template",
			args: []any{
				[]string{"-n", "test", "-t", "test"},
			},
			expected_value: nil,
			expected_error: fmt.Errorf("invalid template: test"),
		},
		{
			Name: "returns options if valid arguments",
			args: []any{
				[]string{"-n", "test", "-t", "go", "-G", "-S", "-p", "tmp/", "-h"},
			},
			expected_value: Options{
				Name:     "test",
				Template: templates.Template{Name: "go"},
				NoGit:    true,
				NoSync:   true,
				Path:     "tmp/",
				Help:     true,
			},
			expected_error: nil,
		},
		{
			Name: "returns options if valid arguments with long version",
			args: []any{
				[]string{"--name", "test", "--template", "go", "--no-git", "--no-sync", "--path", "tmp/", "--help"},
			},
			expected_value: Options{
				Name:     "test",
				Template: templates.Template{Name: "go"},
				NoGit:    true,
				NoSync:   true,
				Path:     "tmp/",
				Help:     true,
			},
			expected_error: nil,
		},
	}

	for _, tc := range cases {
		arg := tc.args[0].([]string)

		actual, err := ParseArgs(arg)
		if err != nil {
			if err.Error() != tc.expected_error.Error() {
				t.Errorf(FAILURE_MESSAGE, tc.Name, tc.expected_error, err)
			}
			if actual != nil {
				t.Errorf(FAILURE_MESSAGE, tc.Name, tc.expected_value, actual)
			}
		} else {
			if *actual != tc.expected_value {
				t.Errorf(FAILURE_MESSAGE, tc.Name, tc.expected_value, actual)
			}
			if err != nil {
				t.Errorf(FAILURE_MESSAGE, tc.Name, tc.expected_error, err)
			}
		}
	}
}

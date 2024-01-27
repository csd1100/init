package templates_test

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/csd1100/init/internal/helpers"
	"github.com/csd1100/init/internal/templates"
)

func TestParseTemplate(t *testing.T) {
	cases := []struct {
		name           string
		templ          templates.Template
		filename       string
		expected_error error
		expected_value string
	}{
		{
			name: "ParseTemplates generates parsed file",
			templ: templates.Template{
				Name: "test",
				TemplateFiles: []templates.TemplateFile{
					{
						Src: "./testdata/test_data.json.tmpl",
					},
				},
				TemplateData: map[string]string{
					"key":   "key",
					"value": "value",
				},
			},
			expected_value: fmt.Sprintln(`{
    "key": "value"
}`),
		},
		{
			name: "ParseTemplates replaces existing file",
			templ: templates.Template{
				Name: "test",
				TemplateFiles: []templates.TemplateFile{
					{
						Src: "./testdata/test_data.json.tmpl",
					},
				},
				TemplateData: map[string]string{
					"key":   "new_key",
					"value": "new_value",
				},
			},
			filename: "exisiting_file.json",
			expected_value: fmt.Sprintln(`{
    "new_key": "new_value"
}`),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			dir := t.TempDir()
			tc.templ.TemplateFiles[0].Dst = fmt.Sprintf("%v/test_data.json", dir)

			err := tc.templ.ParseTemplates()

			if err != nil {
				if !errors.Is(err, tc.expected_error) {
					t.Errorf(helpers.FAILURE_MESSAGE, tc.name, helpers.ERROR, tc.expected_error, err)
				}
			} else {
				actualFile, readErr := os.ReadFile(tc.templ.TemplateFiles[0].Dst)
				if readErr != nil {
					t.Errorf("unable to read actual file")
				}

				if strings.Compare(string(actualFile), tc.expected_value) != 0 {
					t.Errorf(helpers.FAILURE_MESSAGE, tc.name, helpers.VALUE, tc.expected_value, string(actualFile))
				}

				if err != nil {
					t.Errorf(helpers.FAILURE_MESSAGE, tc.name, helpers.ERROR, tc.expected_error, err)
				}
			}
		})
	}
}

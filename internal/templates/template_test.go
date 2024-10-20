package templates_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/csd1100/init/internal/helpers"
	"github.com/csd1100/init/internal/templates"
)

func TestParseTemplate(t *testing.T) {
	cases := []struct {
		name          string
		tmpl          templates.Template
		filename      string
		expectedError error
		expectedValue string
	}{
		{
			name: "ParseTemplates generates parsed file",
			tmpl: templates.Template{
				Name: "test",
				TemplateFiles: templates.TemplateFiles{
					Files: []templates.TemplateFile{
						{
							Template: "./testdata/test_data.json.tmpl",
						},
					},
				},
				TemplateData: map[string]string{
					"key":   "key",
					"value": "value",
				},
			},
			expectedValue: fmt.Sprintln(`{
    "key": "value"
}`),
		},
		{
			name: "ParseTemplates replaces existing file",
			tmpl: templates.Template{
				Name: "test",
				TemplateFiles: templates.TemplateFiles{
					Files: []templates.TemplateFile{
						{
							Template: "./testdata/test_data.json.tmpl",
						},
					},
				},
				TemplateData: map[string]string{
					"key":   "new_key",
					"value": "new_value",
				},
			},
			filename: "existing_file.json",
			expectedValue: fmt.Sprintln(`{
    "new_key": "new_value"
}`),
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			dir := t.TempDir()
			tc.tmpl.TemplateFiles.Files[0].Real = fmt.Sprintf("%v/test_data.json", dir)

			err := tc.tmpl.ParseTemplates()
			actualFile, readErr := os.ReadFile(tc.tmpl.TemplateFiles.Files[0].Real)

			helpers.ValidateExpectations(t, tc.name, string(actualFile), tc.expectedValue, err, tc.expectedError,
				func(actual any, expected any) error {
					if readErr != nil {
						return fmt.Errorf("unable to read actual file")
					}

					if strings.Compare(string(actualFile), tc.expectedValue) != 0 {
						return fmt.Errorf("actual file does not match expected value")
					}

					return nil
				})
		})
	}
}

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

			if err != nil {
				if !errors.Is(err, tc.expectedError) {
					t.Errorf(helpers.FailureMessage, tc.name, helpers.ERROR, tc.expectedError, err)
				}
			} else {
				actualFile, readErr := os.ReadFile(tc.tmpl.TemplateFiles.Files[0].Real)
				if readErr != nil {
					t.Errorf("unable to read actual file")
				}

				if strings.Compare(string(actualFile), tc.expectedValue) != 0 {
					t.Errorf(helpers.FailureMessage, tc.name, helpers.VALUE, tc.expectedValue, string(actualFile))
				}

				if err != nil {
					t.Errorf(helpers.FailureMessage, tc.name, helpers.ERROR, tc.expectedError, err)
				}
			}
		})
	}
}

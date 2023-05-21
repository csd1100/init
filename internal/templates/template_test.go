package templates_test

import (
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/csd1100/init/internal/templates"
	"github.com/csd1100/init/internal/utils"
)

// TODO: fix test arguments passed to the ParseArgs
func TestParseTemplate(t *testing.T) {
	cases := []struct {
		name           string
		templ          templates.Template
		expected_value []byte
		expected_error error
	}{
		{
			name: "ParseTemplates happy path",
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
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			dir := t.TempDir()
			tc.templ.TemplateFiles[0].Dst = fmt.Sprintf("%v/test_data.json", dir)

			err := tc.templ.ParseTemplates()

			if err != nil {
				if !errors.Is(err, tc.expected_error) {
					t.Errorf(utils.FAILURE_MESSAGE, tc.name, utils.ERROR, tc.expected_error, err)
				}
			} else {
				cwd, readErr := os.Getwd()
				if readErr != nil {
					t.Errorf("unable to get working directory")
				}
				fmt.Println(cwd)

				expectedFilePath := fmt.Sprint("./testdata/expected_test_data.json")
				expectedFile, readErr := os.ReadFile(expectedFilePath)
				if readErr != nil {
					t.Errorf("unable to read expected file")
				}

				actualFile, readErr := os.ReadFile(tc.templ.TemplateFiles[0].Dst)
				if readErr != nil {
					t.Errorf("unable to read actual file")
				}

				if string(actualFile) != string(expectedFile) {
					t.Errorf(utils.FAILURE_MESSAGE, tc.name, utils.VALUE, string(expectedFile), string(actualFile))
				}

				if err != nil {
					t.Errorf(utils.FAILURE_MESSAGE, tc.name, utils.ERROR, tc.expected_error, err)
				}
			}
		})
	}
}

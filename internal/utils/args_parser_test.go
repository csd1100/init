package utils_test

import (
	"errors"
	"flag"
	"os"
	"reflect"
	"testing"

	"github.com/csd1100/init/internal/cli"
	"github.com/csd1100/init/internal/helpers"
	"github.com/csd1100/init/internal/templates"
	"github.com/csd1100/init/internal/utils"
)

func TestParse(t *testing.T) {
	testTemplateFiles := []templates.TemplateFile{
		{
			Src: "./templates/main.go.tmpl",
			Dst: "./main.go",
		},
	}

	testTemplateData := make(map[string]string)
	testTemplateData[helpers.PROJECT_NAME] = "test"

	testTemplateDataWithOptions := make(map[string]string)
	testTemplateDataWithOptions[helpers.PROJECT_NAME] = "test"
	testTemplateDataWithOptions["key1"] = "val1"
	testTemplateDataWithOptions["key2"] = "val2"

	cases := []struct {
		name           string
		init           func()
		expected_value *utils.Options
		expected_error error
	}{
		{
			name:           "returns error if empty args",
			expected_value: nil,
			expected_error: utils.ErrArgNameRequired,
		},
		{
			name: "returns error if name not included",
			init: func() {
				utils.FSet.Set("t", "go")
			},
			expected_value: nil,
			expected_error: utils.ErrArgNameRequired,
		},
		{
			name: "returns error if template not included",
			init: func() {
				utils.FSet.Set("n", "test")
			},
			expected_value: nil,
			expected_error: utils.ErrArgTemplateRequired,
		},
		{
			name: "returns error if invalid template",
			init: func() {
				utils.FSet.Set("n", "test")
				utils.FSet.Set("t", "test")
			},
			expected_value: nil,
			expected_error: templates.ErrInvalidArgTemplate,
		},
		{
			name: "returns error if invalid path",
			init: func() {
				utils.FSet.Set("n", "test")
				utils.FSet.Set("t", "go")
				utils.FSet.Set("p", "invalid_directory_11111")
			},
			expected_value: nil,
			expected_error: utils.ErrInvalidArgPath,
		},
		{
			name: "returns options if only name and template",
			init: func() {
				utils.FSet.Set("n", "test")
				utils.FSet.Set("t", "go")
			},
			expected_value: &utils.Options{
				Name: "test",
				Template: templates.Template{
					Name:          "go",
					TemplateFiles: testTemplateFiles,
					TemplateData:  testTemplateData,
					BuildTool:     cli.Go,
				},
			},
			expected_error: nil,
		},
		{
			name: "returns options if valid arguments",
			init: func() {
				utils.FSet.Set("n", "test")
				utils.FSet.Set("t", "go")
				utils.FSet.Set("G", "true")
				utils.FSet.Set("S", "true")
				utils.FSet.Set("p", "/tmp/")
				utils.FSet.Set("o", "key1=val1,key2=val2")
			},
			expected_value: &utils.Options{
				Name: "test",
				Template: templates.Template{
					Name:          "go",
					TemplateFiles: testTemplateFiles,
					TemplateData:  testTemplateDataWithOptions,
					BuildTool:     cli.Go,
				},
				NoGit:  true,
				NoSync: true,
				Path:   "/tmp/",
			},
			expected_error: nil,
		},
		{
			name: "returns options if valid arguments with long version",
			init: func() {
				utils.FSet.Set("name", "test")
				utils.FSet.Set("template", "go")
				utils.FSet.Set("no-git", "true")
				utils.FSet.Set("no-sync", "true")
				utils.FSet.Set("path", "/tmp/")
				utils.FSet.Set("options", "key1=val1,key2=val2")
			},
			expected_value: &utils.Options{
				Name: "test",
				Template: templates.Template{
					Name:          "go",
					TemplateFiles: testTemplateFiles,
					TemplateData:  testTemplateDataWithOptions,
					BuildTool:     cli.Go,
				},
				NoGit:  true,
				NoSync: true,
				Path:   "/tmp/",
			},
			expected_error: nil,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			oldArgs := os.Args
			os.Args = oldArgs[:1]

			if tc.init != nil {
				tc.init()
			}

			t.Cleanup(func() {
				utils.FSet.Visit(func(f *flag.Flag) {
					f.Value.Set(f.DefValue)
				})
			})

			actual, err := utils.ParseArgs()

			if err != nil {
				if !errors.Is(err, tc.expected_error) {
					t.Errorf(utils.FAILURE_MESSAGE, tc.name, utils.ERROR, tc.expected_error, err)
				}
				if actual != nil {
					t.Errorf(utils.FAILURE_MESSAGE, tc.name, utils.VALUE, tc.expected_value, actual)
				}
			} else {
				if !reflect.DeepEqual(*actual, *tc.expected_value) {
					t.Errorf(utils.FAILURE_MESSAGE, tc.name, utils.VALUE, *tc.expected_value, *actual)
				}
				if err != nil {
					t.Errorf(utils.FAILURE_MESSAGE, tc.name, utils.ERROR, tc.expected_error, err)
				}
			}

			os.Args = oldArgs
		})
	}
}

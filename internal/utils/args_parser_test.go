package utils_test

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/csd1100/init/internal/cli"
	"github.com/csd1100/init/internal/helpers"
	"github.com/csd1100/init/internal/templates"
	"github.com/csd1100/init/internal/utils"
)

func TestParse(t *testing.T) {
	testTemplateData := make(map[string]string)
	testTemplateData[helpers.ProjectName] = "test"
	testTemplateData[helpers.GoPackageName] = "project"

	testTemplateDataWithOptions := make(map[string]string)
	testTemplateDataWithOptions[helpers.ProjectName] = "test"
	testTemplateDataWithOptions[helpers.GoPackageName] = "project"
	testTemplateDataWithOptions["key1"] = "val1"
	testTemplateDataWithOptions["key2"] = "val2"

	currentDir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	currentName := filepath.Base(currentDir)
	testTemplateDataForCurrent := make(map[string]string)
	testTemplateDataForCurrent[helpers.ProjectName] = currentName
	testTemplateDataForCurrent[helpers.GoPackageName] = "project"

	cases := []struct {
		name          string
		init          func()
		expectedValue *utils.Options
		expectedError error
	}{
		{
			name:          "returns error if empty args",
			expectedValue: nil,
			expectedError: helpers.ErrArgNameRequired,
		},
		{
			name: "returns error if name not included",
			init: func() {
				err := utils.FSet.Set("t", "go")
				if err != nil {
					panic("failed to set flag for test")
				}
			},
			expectedValue: nil,
			expectedError: helpers.ErrArgNameRequired,
		},
		{
			name: "returns error if template not included",
			init: func() {
				err := utils.FSet.Set("n", "test")
				if err != nil {
					panic("failed to set flag for test")
				}
			},
			expectedValue: nil,
			expectedError: helpers.ErrArgTemplateRequired,
		},
		{
			name: "returns error if invalid template",
			init: func() {
				err := utils.FSet.Set("n", "test")
				if err != nil {
					panic("failed to set flag for test")
				}
				err = utils.FSet.Set("t", "test")
				if err != nil {
					panic("failed to set flag for test")
				}
			},
			expectedValue: nil,
			expectedError: helpers.ErrInvalidArgTemplate,
		},
		{
			name: "returns error if invalid path",
			init: func() {
				err := utils.FSet.Set("n", "test")
				if err != nil {
					panic("failed to set flag for test")
				}
				err = utils.FSet.Set("t", "go")
				if err != nil {
					panic("failed to set flag for test")
				}
				err = utils.FSet.Set("p", "invalid_directory_11111")
				if err != nil {
					panic("failed to set flag for test")
				}
			},
			expectedValue: nil,
			expectedError: helpers.ErrInvalidArgPath,
		},
		{
			name: "returns options if only name and template",
			init: func() {
				_ = utils.FSet.Set("n", "test")
				_ = utils.FSet.Set("t", "go")
			},
			expectedValue: &utils.Options{
				Name: "test",
				Template: &templates.Template{
					Name:         "go",
					TemplateData: testTemplateData,
					BuildTool:    cli.Go,
				},
			},
			expectedError: nil,
		},
		{
			name: "returns options if valid arguments with path",
			init: func() {
				_ = utils.FSet.Set("n", "test")
				_ = utils.FSet.Set("t", "go")
				_ = utils.FSet.Set("G", "true")
				_ = utils.FSet.Set("S", "true")
				_ = utils.FSet.Set("p", "/tmp/")
				_ = utils.FSet.Set("o", "key1=val1,key2=val2")
			},
			expectedValue: &utils.Options{
				Name: "test",
				Template: &templates.Template{
					Name:         "go",
					TemplateData: testTemplateDataWithOptions,
					BuildTool:    cli.Go,
				},
				NoGit:  true,
				NoSync: true,
				Path:   "/tmp/",
			},
			expectedError: nil,
		},
		{
			name: "returns options if valid arguments with current",
			init: func() {
				_ = utils.FSet.Set("n", "test")
				_ = utils.FSet.Set("t", "go")
				_ = utils.FSet.Set("G", "true")
				_ = utils.FSet.Set("S", "true")
				_ = utils.FSet.Set("c", "true")
				_ = utils.FSet.Set("o", "key1=val1,key2=val2")
			},
			expectedValue: &utils.Options{
				Name: "test",
				Template: &templates.Template{
					Name:         "go",
					TemplateData: testTemplateDataWithOptions,
					BuildTool:    cli.Go,
				},
				NoGit:   true,
				NoSync:  true,
				Current: true,
			},
			expectedError: nil,
		},
		{
			name: "returns options if valid arguments with long version with path",
			init: func() {
				_ = utils.FSet.Set("name", "test")
				_ = utils.FSet.Set("template", "go")
				_ = utils.FSet.Set("no-git", "true")
				_ = utils.FSet.Set("no-sync", "true")
				_ = utils.FSet.Set("path", "/tmp/")
				_ = utils.FSet.Set("options", "key1=val1,key2=val2")
			},
			expectedValue: &utils.Options{
				Name: "test",
				Template: &templates.Template{
					Name:         "go",
					TemplateData: testTemplateDataWithOptions,
					BuildTool:    cli.Go,
				},
				NoGit:  true,
				NoSync: true,
				Path:   "/tmp/",
			},
			expectedError: nil,
		},
		{
			name: "returns options if valid arguments with long version with current",
			init: func() {
				_ = utils.FSet.Set("name", "test")
				_ = utils.FSet.Set("template", "go")
				_ = utils.FSet.Set("no-git", "true")
				_ = utils.FSet.Set("no-sync", "true")
				_ = utils.FSet.Set("current", "true")
				_ = utils.FSet.Set("options", "key1=val1,key2=val2")
			},
			expectedValue: &utils.Options{
				Name: "test",
				Template: &templates.Template{
					Name:         "go",
					TemplateData: testTemplateDataWithOptions,
					BuildTool:    cli.Go,
				},
				NoGit:   true,
				NoSync:  true,
				Current: true,
			},
			expectedError: nil,
		},
		{
			name: "returns error if current is used with path",
			init: func() {
				err := utils.FSet.Set("n", "test")
				if err != nil {
					panic("failed to set flag for test")
				}
				err = utils.FSet.Set("t", "js")
				if err != nil {
					panic("failed to set flag for test")
				}
				err = utils.FSet.Set("c", "true")
				if err != nil {
					panic("failed to set flag for test")
				}
				err = utils.FSet.Set("p", "/tmp")
				if err != nil {
					panic("failed to set flag for test")
				}
			},
			expectedValue: nil,
			expectedError: helpers.ErrArgCurrentNotWithPath,
		},
		{
			name: "returns options if only current and template",
			init: func() {
				_ = utils.FSet.Set("c", "true")
				_ = utils.FSet.Set("t", "go")
			},
			expectedValue: &utils.Options{
				Name: currentName,
				Template: &templates.Template{
					Name:         "go",
					TemplateData: testTemplateDataForCurrent,
					BuildTool:    cli.Go,
				},
				Current: true,
			},
			expectedError: nil,
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
					_ = f.Value.Set(f.DefValue)
				})
			})

			actual, err := utils.ParseArgs()

			helpers.ValidateExpectations(t, tc.name, actual, tc.expectedValue, err, tc.expectedError,
				func(actual any, expected any) error {
					if !reflect.DeepEqual(actual, expected) {
						return fmt.Errorf("expected %v, got %v", expected, actual)
					}
					return nil
				})

			os.Args = oldArgs
		})
	}
}

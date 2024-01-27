package cli

import (
	"errors"
	"reflect"
	"testing"

	"github.com/csd1100/init/internal/helpers"
)

type mockGo struct {
	actualArgs  []string
	actualError error
}

func (mg *mockGo) Exec(subcommand string, args []string) error {
	mg.actualArgs = append([]string{subcommand}, args...)
	return nil
}

func (mg *mockGo) GetCommand() string {
	return "mock"
}

func TestGoModInit(t *testing.T) {
	testcases := []struct {
		name           string
		projectName    string
		expectedArgs   []string
		exepectedError error
	}{
		{
			name:           "Go.ModInit passes correct args",
			projectName:    "go_test",
			expectedArgs:   []string{"mod", "init", "go_test"},
			exepectedError: nil,
		},
	}

	for _, tc := range testcases {

		t.Run(tc.name, func(t *testing.T) {

			mGo := goLang{exe: &mockGo{}}
			err := mGo.ModInit(tc.projectName)
			if err != nil {
				if !errors.Is(err, tc.exepectedError) {
					t.Errorf(helpers.FAILURE_MESSAGE,
						tc.name,
						helpers.ERROR,
						tc.exepectedError,
						err)
				}
			} else {
				if !reflect.DeepEqual(mGo.exe.(*mockGo).actualArgs, tc.expectedArgs) {
					t.Errorf(helpers.FAILURE_MESSAGE,
						tc.name,
						helpers.VALUE,
						tc.expectedArgs,
						mGo.exe.(*mockGo).actualArgs)
				}
			}

		})

	}

}

func TestGoModTidy(t *testing.T) {
	testcases := []struct {
		name           string
		projectName    string
		expectedArgs   []string
		exepectedError error
	}{
		{
			name:           "Go.ModTidy passes correct args",
			expectedArgs:   []string{"mod", "tidy"},
			exepectedError: nil,
		},
	}

	for _, tc := range testcases {

		t.Run(tc.name, func(t *testing.T) {

			mGo := goLang{exe: &mockGo{}}
			err := mGo.ModTidy()
			if err != nil {
				if !errors.Is(err, tc.exepectedError) {
					t.Errorf(helpers.FAILURE_MESSAGE,
						tc.name,
						helpers.ERROR,
						tc.exepectedError,
						err)
				}
			} else {
				if !reflect.DeepEqual(mGo.exe.(*mockGo).actualArgs, tc.expectedArgs) {
					t.Errorf(helpers.FAILURE_MESSAGE,
						tc.name,
						helpers.VALUE,
						tc.expectedArgs,
						mGo.exe.(*mockGo).actualArgs)
				}
			}

		})

	}

}

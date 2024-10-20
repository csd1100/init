package cli

import (
	"fmt"
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
		name          string
		projectName   string
		expectedArgs  []string
		expectedError error
	}{
		{
			name:          "Go.ModInit passes correct args",
			projectName:   "go_test",
			expectedArgs:  []string{"mod", "init", "go_test"},
			expectedError: nil,
		},
	}

	for _, tc := range testcases {

		t.Run(tc.name, func(t *testing.T) {
			mGo := goLang{exe: &mockGo{}}

			err := mGo.ModInit(tc.projectName)
			actualArgs := mGo.exe.(*mockGo).actualArgs

			helpers.ValidateExpectations(t, tc.name, actualArgs, tc.expectedArgs, err, tc.expectedError,
				func(actual any, expected any) error {
					if !reflect.DeepEqual(actual, expected) {
						return fmt.Errorf("expected %v, got %v", expected, actual)
					}
					return nil
				})
		})

	}

}

func TestGoModTidy(t *testing.T) {
	testcases := []struct {
		name          string
		projectName   string
		expectedArgs  []string
		expectedError error
	}{
		{
			name:          "Go.ModTidy passes correct args",
			expectedArgs:  []string{"mod", "tidy"},
			expectedError: nil,
		},
	}

	for _, tc := range testcases {

		t.Run(tc.name, func(t *testing.T) {
			mGo := goLang{exe: &mockGo{}}

			err := mGo.ModTidy()
			actualArgs := mGo.exe.(*mockGo).actualArgs

			helpers.ValidateExpectations(t, tc.name, actualArgs, tc.expectedArgs, err, tc.expectedError,
				func(actual any, expected any) error {
					if !reflect.DeepEqual(actual, expected) {
						return fmt.Errorf("expected %v, got %v", expected, actual)
					}
					return nil
				})
		})

	}

}

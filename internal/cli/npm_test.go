package cli

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/csd1100/init/internal/helpers"
)

type mockNpm struct {
	actualArgs  []string
	actualError error
}

func (mn *mockNpm) Exec(subcommand string, args []string) error {
	mn.actualArgs = append([]string{subcommand}, args...)
	return nil
}

func (mn *mockNpm) GetCommand() string {
	return "mock"
}

func TestNpmInstall(t *testing.T) {
	testcases := []struct {
		name          string
		expectedArgs  []string
		expectedError error
	}{
		{
			name:          "NpmInstall passes correct args",
			expectedArgs:  []string{"install"},
			expectedError: nil,
		},
	}

	for _, tc := range testcases {

		t.Run(tc.name, func(t *testing.T) {
			mNpm := npmCLI{exe: &mockNpm{}}

			err := mNpm.Install()
			actualArgs := mNpm.exe.(*mockNpm).actualArgs

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

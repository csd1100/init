package cli

import (
	"errors"
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
			if err != nil {
				if !errors.Is(err, tc.expectedError) {
					t.Errorf(helpers.FailureMessage,
						tc.name,
						helpers.ERROR,
						tc.expectedError,
						err)
				}
			} else {
				if !reflect.DeepEqual(mNpm.exe.(*mockNpm).actualArgs, tc.expectedArgs) {
					t.Errorf(helpers.FailureMessage,
						tc.name,
						helpers.VALUE,
						tc.expectedArgs,
						mNpm.exe.(*mockNpm).actualArgs)
				}
			}

		})

	}

}

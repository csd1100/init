package cli

import (
	"reflect"
	"testing"

	"github.com/csd1100/init/internal/utils"
)

func TestNpmInstall(t *testing.T) {
	testcases := []struct {
		name          string
		expectedArgs  []string
		expectedError error
	}{
		{
			name:          "Npm.Install passes proper args to Npm.Exec",
			expectedArgs:  []string{"install"},
			expectedError: nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			mockNpm := npm{mockCLI{}}
			err := mockNpm.Install()
			if err != nil {
				if tc.expectedError.Error() != mockNpm.cli.(mockCLI).actualError.Error() {
					t.Errorf(utils.FAILURE_MESSAGE, tc.name, utils.ERROR, tc.expectedError.Error(), err.Error())
				}
			} else {
				if reflect.DeepEqual(mockNpm.cli.(mockCLI).actualArgs, tc.expectedArgs) {
					t.Errorf(utils.FAILURE_MESSAGE, tc.name, utils.VALUE, tc.expectedArgs, mockNpm.cli.(mockCLI).actualArgs)
				}
			}
		})
	}
}

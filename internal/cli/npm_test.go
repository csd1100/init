package cli_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/csd1100/init/internal/cli"
	"github.com/csd1100/init/internal/utils"
)

type mockNpm struct {
	cli.NpmCLI
	actualArgs  []string
	actualError error
}

func (mn *mockNpm) Exec(subcommand string, args []string) ([]byte, error) {
	mn.actualArgs = args
	return nil, nil
}

func TestNpmInstall(t *testing.T) {
	testcases := []struct {
		name           string
		expectedArgs   []string
		exepectedError error
	}{
		{
			name:           "NpmInstall passes correct args",
			expectedArgs:   []string{"install"},
			exepectedError: nil,
		},
	}

	for _, tc := range testcases {

		t.Run(tc.name, func(t *testing.T) {

			mNpm := mockNpm{}
			err := mNpm.Install()
			if err != nil {
				if errors.Is(err, tc.exepectedError) {
					t.Errorf(utils.FAILURE_MESSAGE, tc.name, utils.ERROR, tc.exepectedError, err)
				}
			} else {
				if !reflect.DeepEqual(mNpm.actualArgs, tc.expectedArgs) {
					t.Errorf(utils.FAILURE_MESSAGE, tc.name, utils.VALUE, tc.expectedArgs, mNpm.actualArgs)
				}
			}

		})

	}

}

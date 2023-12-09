package cli_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/csd1100/init/internal/cli"
	"github.com/csd1100/init/internal/utils"
)

type mockCLI struct {
	cli.CLI
	actualArgs  []string
	actualError error
}

func (mc *mockCLI) Exec(subcommand string, args []string) ([]byte, error) {
	mc.actualArgs = args
	return nil, nil
}

func TestExec(t *testing.T) {
	testcases := []struct {
		name           string
		cli            cli.Executable
		expectedArgs   []string
		exepectedError error
	}{
		{
			name:           "cli.Exec receives correct args",
			cli:            nil,
			expectedArgs:   []string{"--test"},
			exepectedError: nil,
		},
		{
			name:           "cli.Exec returns error if not in path",
			cli:            cli.CLI{Command: "not_installed_executable"},
			expectedArgs:   []string{"--test"},
			exepectedError: fmt.Errorf("not_installed_executable is not installed"),
		},
	}

	for _, tc := range testcases {

		var testCLI cli.Executable
		if tc.cli == nil {
			testCLI = &mockCLI{}
		} else {
			testCLI = tc.cli
		}

		t.Run(tc.name, func(t *testing.T) {
			_, err := testCLI.Exec("", tc.expectedArgs)
			if err != nil {
				if err.Error() != tc.exepectedError.Error() {
					t.Errorf(utils.FAILURE_MESSAGE, tc.name, utils.ERROR, tc.exepectedError, err)
				}
			} else {
				if !reflect.DeepEqual(testCLI.(*mockCLI).actualArgs, tc.expectedArgs) {
					t.Errorf(utils.FAILURE_MESSAGE, tc.name, utils.VALUE, tc.expectedArgs, testCLI.(*mockCLI).actualArgs)
				}

			}
		})
	}

}

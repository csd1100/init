package cli_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/csd1100/init/internal/cli"
	"github.com/csd1100/init/internal/utils"
)

func TestNewCli(t *testing.T) {
	thisPath, _ := os.Executable()
	testcases := []struct {
		name          string
		command       string
		exepetedValue cli.CLI
	}{
		{
			name:    "NewCLI has IsInstalled true if program is in PATH",
			command: thisPath,
			exepetedValue: cli.CLI{
				Command:     thisPath,
				Path:        thisPath,
				IsInstalled: true,
			},
		},
		{
			name:    "NewCLI has IsInstalled false if program is NOT in PATH",
			command: "non_existent_executable",
			exepetedValue: cli.CLI{
				Command:     "non_existent_executable",
				Path:        "",
				IsInstalled: false,
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			actual := cli.NewCLI(tc.command)
			if actual != tc.exepetedValue {
				t.Errorf(utils.FAILURE_MESSAGE, tc.name, utils.VALUE, tc.exepetedValue, actual)
			}
		})
	}
}

func TestExec(t *testing.T) {
	thisPath, _ := os.Executable()
	testcases := []struct {
		name          string
		command       string
		exepetedError error
	}{
		{
			name:          "cli.Exec does not return error if executable is present in PATH",
			command:       thisPath,
			exepetedError: nil,
		},
		{
			name:          "cli.Exec does return error if executable is not present in PATH",
			command:       "non_existent_executable",
			exepetedError: fmt.Errorf("non_existent_executable is not installed"),
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			executable := cli.NewCLI(tc.command)
			actualError := executable.Exec([]string{"arg1"})
			if actualError != nil {
				if actualError.Error() != tc.exepetedError.Error() {
					t.Errorf(utils.FAILURE_MESSAGE, tc.name, utils.ERROR, tc.exepetedError, actualError)
				}
			}
		})
	}

}

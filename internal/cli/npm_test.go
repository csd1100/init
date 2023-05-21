package cli_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/csd1100/init/internal/cli"
	"github.com/csd1100/init/internal/utils"
)

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

			mcli := mockCLI{}
			err := cli.NpmInstall(&mcli)
			if err != nil {
				if errors.Is(err, tc.exepectedError) {
					t.Errorf(utils.FAILURE_MESSAGE, tc.name, utils.ERROR, tc.exepectedError, err)
				}
			} else {
				if !reflect.DeepEqual(mcli.actualArgs, tc.expectedArgs) {
					t.Errorf(utils.FAILURE_MESSAGE, tc.name, utils.VALUE, tc.expectedArgs, mcli.actualArgs)
				}
			}

		})

	}

}

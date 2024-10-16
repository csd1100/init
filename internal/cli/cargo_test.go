package cli

import (
	"errors"
	"reflect"
	"testing"

	"github.com/csd1100/init/internal/helpers"
)

type mockCargo struct {
	actualArgs  []string
	actualError error
}

func (mc *mockCargo) Exec(subcommand string, args []string) error {
	mc.actualArgs = append([]string{subcommand}, args...)
	return nil
}

func (mc *mockCargo) GetCommand() string {
	return "mock"
}

func TestCargo(t *testing.T) {
	testcases := []struct {
		name              string
		expectedArgs      []string
		expectedError     error
		functionUnderTest func(cargoCLI) error
	}{
		{
			name:          "CargoInit passes correct args",
			expectedArgs:  []string{"init"},
			expectedError: nil,
			functionUnderTest: func(cargo cargoCLI) error {
				return cargo.Init()
			},
		},
		{
			name:          "CargoCheck passes correct args",
			expectedArgs:  []string{"check"},
			expectedError: nil,
			functionUnderTest: func(cargo cargoCLI) error {
				return cargo.Check()
			},
		},
		{
			name:          "CargoClean passes correct args",
			expectedArgs:  []string{"clean"},
			expectedError: nil,
			functionUnderTest: func(cargo cargoCLI) error {
				return cargo.Clean()
			},
		},
		{
			name:          "CargoFetch passes correct args",
			expectedArgs:  []string{"fetch"},
			expectedError: nil,
			functionUnderTest: func(cargo cargoCLI) error {
				return cargo.Fetch()
			},
		},
	}

	for _, tc := range testcases {

		t.Run(tc.name, func(t *testing.T) {

			mCargo := cargoCLI{exe: &mockCargo{}}
			err := tc.functionUnderTest(mCargo)
			if err != nil {
				if !errors.Is(err, tc.expectedError) {
					t.Errorf(helpers.FailureMessage,
						tc.name,
						helpers.ERROR,
						tc.expectedError,
						err)
				}
			} else {
				if !reflect.DeepEqual(mCargo.exe.(*mockCargo).actualArgs, tc.expectedArgs) {
					t.Errorf(helpers.FailureMessage,
						tc.name,
						helpers.VALUE,
						tc.expectedArgs,
						mCargo.exe.(*mockCargo).actualArgs)
				}
			}

		})

	}

}

package cli

import (
	"fmt"
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
			actualArgs := mCargo.exe.(*mockCargo).actualArgs

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

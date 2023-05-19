package cli

import (
	"reflect"
	"testing"

	"github.com/csd1100/init/internal/utils"
)

type mockCLI struct {
	actualArgs  []string
	actualError error
}

func (mcli mockCLI) Exec(args []string) error {
	mcli.actualArgs = args
	mcli.actualError = nil
	return nil
}

func TestGitClone(t *testing.T) {
	testcases := []struct {
		name          string
		expectedArgs  []string
		expectedError error
	}{
		{
			name:          "Git.Clone passes proper args to Git.Exec",
			expectedArgs:  []string{"clone", "test"},
			expectedError: nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			mockGit := git{mockCLI{}}
			err := mockGit.Clone("test")
			if err != nil {
				if tc.expectedError != mockGit.cli.(mockCLI).actualError {
					t.Errorf(utils.FAILURE_MESSAGE, tc.name, utils.ERROR, tc.expectedError, mockGit.cli.(mockCLI).actualError)
				}
			} else {
				if reflect.DeepEqual(mockGit.cli.(mockCLI).actualArgs, tc.expectedArgs) {
					t.Errorf(utils.FAILURE_MESSAGE, tc.name, utils.VALUE, tc.expectedArgs, mockGit.cli.(mockCLI).actualArgs)
				}
			}
		})
	}
}

func TestGitInit(t *testing.T) {
	testcases := []struct {
		name          string
		expectedArgs  []string
		expectedError error
	}{
		{
			name:          "Git.Init passes proper args to Git.Exec",
			expectedArgs:  []string{"init"},
			expectedError: nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			mockGit := git{mockCLI{}}
			err := mockGit.Init()
			if err != nil {
				if tc.expectedError.Error() != mockGit.cli.(mockCLI).actualError.Error() {
					t.Errorf(utils.FAILURE_MESSAGE, tc.name, utils.ERROR, tc.expectedError.Error(), err.Error())
				}
			} else {
				if reflect.DeepEqual(mockGit.cli.(mockCLI).actualArgs, tc.expectedArgs) {
					t.Errorf(utils.FAILURE_MESSAGE, tc.name, utils.VALUE, tc.expectedArgs, mockGit.cli.(mockCLI).actualArgs)
				}
			}
		})
	}
}

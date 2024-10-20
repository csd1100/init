package cli

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/csd1100/init/internal/helpers"
)

type mockGit struct {
	actualArgs  []string
	actualError error
}

func (mg *mockGit) Exec(subcommand string, args []string) error {
	mg.actualArgs = append([]string{subcommand}, args...)
	return nil
}

func (mg *mockGit) GetCommand() string {
	return "mock"
}

func TestGitClone(t *testing.T) {
	testcases := []struct {
		name          string
		repo          string
		args          []string
		expectedArgs  []string
		expectedError error
	}{
		{
			name:          "GitClone without args",
			repo:          "test",
			args:          []string{},
			expectedArgs:  []string{"clone", "test"},
			expectedError: nil,
		},
		{
			name:          "GitClone with args",
			repo:          "test",
			args:          []string{"xxxx"},
			expectedArgs:  []string{"clone", "xxxx", "test"},
			expectedError: nil,
		},
	}

	for _, tc := range testcases {

		t.Run(tc.name, func(t *testing.T) {
			mGit := gitCLI{exe: &mockGit{}}
			err := mGit.Clone(tc.repo, tc.args)
			actualArgs := mGit.exe.(*mockGit).actualArgs

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

func TestGitCloneSingleBranch(t *testing.T) {
	testcases := []struct {
		name          string
		repo          string
		branch        string
		expectedArgs  []string
		expectedError error
	}{
		{
			name:   "GitCloneSingleBranch",
			repo:   "https://github.com/test/test",
			branch: "test-1",
			expectedArgs: []string{
				"clone", "--single-branch", "-b", "test-1",
				"--depth", "1", "https://github.com/test/test",
			},
			expectedError: nil,
		},
	}

	for _, tc := range testcases {

		t.Run(tc.name, func(t *testing.T) {

			mGit := gitCLI{exe: &mockGit{}}
			err := mGit.CloneSingleBranch(tc.repo, tc.branch)
			actualArgs := mGit.exe.(*mockGit).actualArgs

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
func TestGitInit(t *testing.T) {
	testcases := []struct {
		name          string
		expectedArgs  []string
		expectedError error
	}{
		{
			name:          "GitInit passes correct args",
			expectedArgs:  []string{"init"},
			expectedError: nil,
		},
	}

	for _, tc := range testcases {

		t.Run(tc.name, func(t *testing.T) {

			mGit := gitCLI{exe: &mockGit{}}
			err := mGit.Init()
			actualArgs := mGit.exe.(*mockGit).actualArgs

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

package cli_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/csd1100/init/internal/cli"
	"github.com/csd1100/init/internal/utils"
)

type mockGit struct {
	cli.GitCLI
	actualArgs  []string
	actualError error
}

func (mg *mockGit) Exec(subcommand string, args []string) ([]byte, error) {
	mg.actualArgs = args
	return nil, nil
}

func TestGitClone(t *testing.T) {
	testcases := []struct {
		name           string
		repo           string
		actualArgs     []string
		expectedArgs   []string
		exepectedError error
	}{
		{
			name:           "GitClone without args",
			repo:           "test",
			actualArgs:     []string{},
			expectedArgs:   []string{"clone", "test"},
			exepectedError: nil,
		},
		{
			name:           "GitClone with args",
			repo:           "test",
			actualArgs:     []string{"xxxx"},
			expectedArgs:   []string{"clone", "xxxx", "test"},
			exepectedError: nil,
		},
	}

	for _, tc := range testcases {

		t.Run(tc.name, func(t *testing.T) {

			mGit := mockGit{}
			err := mGit.Clone(tc.repo, tc.actualArgs)
			if err != nil {
				if errors.Is(err, tc.exepectedError) {
					t.Errorf(utils.FAILURE_MESSAGE, tc.name, utils.ERROR, tc.exepectedError, err)
				}
			} else {
				if !reflect.DeepEqual(mGit.actualArgs, tc.expectedArgs) {
					t.Errorf(utils.FAILURE_MESSAGE, tc.name, utils.VALUE, tc.expectedArgs, mGit.actualArgs)
				}
			}

		})

	}

}

func TestGitCloneSingleBranch(t *testing.T) {
	testcases := []struct {
		name           string
		repo           string
		branch         string
		expectedArgs   []string
		exepectedError error
	}{
		{
			name:   "GitCloneSingleBranch",
			repo:   "https://github.com/test/test",
			branch: "test-1",
			expectedArgs: []string{
				"clone", "--single-branch", "-b", "test-1",
				"--depth", "1", "https://github.com/test/test",
			},
			exepectedError: nil,
		},
	}

	for _, tc := range testcases {

		t.Run(tc.name, func(t *testing.T) {

			mGit := mockGit{}
			err := mGit.CloneSingleBranch(tc.repo, tc.branch)
			if err != nil {
				if errors.Is(err, tc.exepectedError) {
					t.Errorf(utils.FAILURE_MESSAGE, tc.name, utils.ERROR, tc.exepectedError, err)
				}
			} else {
				if !reflect.DeepEqual(mGit.actualArgs, tc.expectedArgs) {
					t.Errorf(utils.FAILURE_MESSAGE, tc.name, utils.VALUE, tc.expectedArgs, mGit.actualArgs)
				}
			}

		})

	}

}
func TestGitInit(t *testing.T) {
	testcases := []struct {
		name           string
		expectedArgs   []string
		exepectedError error
	}{
		{
			name:           "GitInit passes correct args",
			expectedArgs:   []string{"init"},
			exepectedError: nil,
		},
	}

	for _, tc := range testcases {

		t.Run(tc.name, func(t *testing.T) {

			mGit := mockGit{}
			err := mGit.Init()
			if err != nil {
				if errors.Is(err, tc.exepectedError) {
					t.Errorf(utils.FAILURE_MESSAGE, tc.name, utils.ERROR, tc.exepectedError, err)
				}
			} else {
				if !reflect.DeepEqual(mGit.actualArgs, tc.expectedArgs) {
					t.Errorf(utils.FAILURE_MESSAGE, tc.name, utils.VALUE, tc.expectedArgs, mGit.actualArgs)
				}
			}

		})

	}

}

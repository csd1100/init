package cli_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/csd1100/init/internal/cli"
	"github.com/csd1100/init/internal/utils"
)

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

			mcli := mockCLI{}
			err := cli.GitClone(&mcli, tc.repo, tc.actualArgs)
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

			mcli := mockCLI{}
			err := cli.GitCloneSingleBranch(&mcli, tc.repo, tc.branch)
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

			mcli := mockCLI{}
			err := cli.GitInit(&mcli)
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

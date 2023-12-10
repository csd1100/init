package cli

import (
	"github.com/csd1100/init/internal/helpers"
)

type GitCLI struct {
	CLI
}

func (Git GitCLI) Clone(repo string, args []string) error {
	if args == nil {
		args = []string{}
	}
	cloneArgs := append(args, repo)
	op, err := Git.Exec("clone", cloneArgs)
	helpers.AppLogger.Debug("Output of git clone:\n %s", op)
	if err != nil {
		return err
	}
	return nil
}

func (Git GitCLI) CloneSingleBranch(repo string, branch string) error {
	return Git.Clone(repo, []string{"--single-branch", "-b", branch, "--depth", "1"})
}

func (Git GitCLI) Init() error {
	op, err := Git.Exec("init", []string{})
	helpers.AppLogger.Debug("Output of git init:\n %s", op)
	if err != nil {
		return err
	}
	return nil
}

var Git = GitCLI{CLI{Command: "git"}}

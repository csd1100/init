package cli

import (
	"fmt"
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
	fmt.Printf("%s\n", op)
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
	fmt.Printf("%s\n", op)
	if err != nil {
		return err
	}
	return nil
}

var Git = GitCLI{CLI{Command: "git"}}

package cli

type gitCLI struct {
	Command string
}

var Git = gitCLI{
	Command: "git",
}

func (git gitCLI) GetCommand() string {
	return git.Command
}

func (git gitCLI) Exec(subcommand string, args []string) error {
	return execute(git, subcommand, args)
}

func (git gitCLI) Clone(repo string, args []string) error {
	if args == nil {
		args = []string{}
	}

	cloneArgs := append(args, repo)
	return git.Exec("clone", cloneArgs)
}

func (git gitCLI) CloneSingleBranch(repo string, branch string) error {
	return git.Clone(repo, []string{"--single-branch", "-b", branch, "--depth", "1"})
}

func (git gitCLI) Init() error {
	return git.Exec("init", []string{})
}

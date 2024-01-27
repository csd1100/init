package cli

type gitExe struct {
	Command string
}

func (git gitExe) GetCommand() string {
	return git.Command
}

func (git gitExe) Exec(subcommand string, args []string) error {
	return execute(git, subcommand, args)
}

func getGitExecutable() gitExe {
	return gitExe{Command: "git"}
}

type gitCLI struct {
	exe Executable
}

var Git = gitCLI{
	exe: getGitExecutable(),
}

func (git gitCLI) Clone(repo string, args []string) error {
	if args == nil {
		args = []string{}
	}

	cloneArgs := append(args, repo)
	return git.exe.Exec("clone", cloneArgs)
}

func (git gitCLI) CloneSingleBranch(repo string, branch string) error {
	return git.Clone(repo, []string{"--single-branch", "-b", branch, "--depth", "1"})
}

func (git gitCLI) Init() error {
	return git.exe.Exec("init", []string{})
}

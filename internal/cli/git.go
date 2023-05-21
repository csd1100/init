package cli

func GitClone(exe Executable, repo string, args []string) error {
	if args == nil {
		args = []string{}
	}
	cloneArgs := append([]string{"clone"}, args...)
	cloneArgs = append(cloneArgs, repo)
	return exe.Exec(cloneArgs)
}

func GitCloneSingleBranch(exe Executable, repo string, branch string) error {
	return GitClone(exe, repo, []string{"--single-branch", "-b", branch, "--depth", "1"})
}

func GitInit(exe Executable) error {
	return exe.Exec([]string{"init"})
}

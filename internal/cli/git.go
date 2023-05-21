package cli

func GitClone(exe Executable, repo string, args []string) error {
	if args == nil {
		args = []string{}
	}
	cloneArgs := append([]string{"clone"}, args...)
	cloneArgs = append(cloneArgs, repo)
	return exe.Exec(cloneArgs)
}

func GitInit(exe Executable) error {
	return exe.Exec([]string{"init"})
}

package cli

func NpmInstall(exe Executable) error {
	return exe.Exec([]string{"install"})
}

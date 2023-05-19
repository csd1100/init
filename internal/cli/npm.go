package cli

type npmInterface interface {
	Install() error
}

type npm struct {
	cli Executable
}

func (npmCLI npm) Install() error {
	return npmCLI.cli.Exec([]string{"install"})
}

var Npm = npm{NewCLI("npm")}

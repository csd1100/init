package cli

type gitInterface interface {
	Clone(repo string) error
	Init() error
}

type git struct {
	cli Executable
}

func (gitCLI git) Clone(repo string) error {
	return gitCLI.cli.Exec([]string{"clone", repo})
}

func (gitCLI git) Init() error {
	return gitCLI.cli.Exec([]string{"init"})
}

var Git = git{NewCLI("git")}

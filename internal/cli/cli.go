package cli

import (
	"fmt"
	"os/exec"
)

type Executable interface {
	Exec(args []string) error
}

type CLI struct {
	Command     string
	Path        string
	IsInstalled bool
}

func (cli CLI) Exec(args []string) error {
	if cli.IsInstalled == false {
		return fmt.Errorf("%s is not installed", cli.Command)
	}

	// cmd := exec.Command(cli.Path, args...)
	// fmt.Println(cmd)
	// err := cmd.Run()
	// if err != nil {
	// 	return err
	// }

	return nil
}

func NewCLI(command string) CLI {
	cli := CLI{}
	cli.Command = command
	path, err := exec.LookPath(cli.Command)
	if err != nil {
		cli.IsInstalled = false
	} else {
		cli.Path = path
		cli.IsInstalled = true
	}
	return cli
}

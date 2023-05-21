package cli

import (
	"fmt"
	"os/exec"
)

type Executable interface {
	Exec([]string) error
}

type CLI struct {
	Command string
}

func (cli CLI) Exec(args []string) error {
	path, err := exec.LookPath(cli.Command)
	if err != nil {
		return fmt.Errorf("%s is not installed", cli.Command)
	}
	fmt.Println(path)
	return nil
}

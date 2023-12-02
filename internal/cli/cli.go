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
	cmd := exec.Command(path, args...)
	fmt.Println(cmd)

	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Printf("%s\n", stdoutStderr)

	return nil
}

package cli

import (
	"fmt"
	"os/exec"

	"github.com/csd1100/init/internal/helpers"
)

type Executable interface {
	Exec(string, []string) ([]byte, error)
}

type BuildTool interface {
	Executable
	Sync(map[string]string) error
}

type CLI struct {
	Command string
}

func (cli CLI) Exec(subcommand string, args []string) ([]byte, error) {
	path, err := exec.LookPath(cli.Command)
	if err != nil {
		return nil, fmt.Errorf("%s is not installed", cli.Command)
	}

	arguments := append([]string{subcommand}, args...)
	cmd := exec.Command(path, arguments...)
	helpers.AppLogger.Debug("Executing Command: %v", cmd)

	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		return nil, err
	}

	return stdoutStderr, nil
}

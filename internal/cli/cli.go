package cli

import (
	"os/exec"

	"github.com/csd1100/init/internal/helpers"
)

type Executable interface {
	GetCommand() string
	Exec(string, []string) error
}

type BuildTool interface {
	Sync(map[string]string) error
}

func execute(exe Executable, subcommand string, args []string) error {
	path, err := exec.LookPath(exe.GetCommand())
	if err != nil {
		return helpers.ErrExecNotFound
	}

	arguments := append([]string{subcommand}, args...)
	cmd := exec.Command(path, arguments...)
	helpers.AppLogger.Debug("Executing Command: %v", cmd)

	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		helpers.AppLogger.Error("Error for command %s:\n%v", cmd, string(stdoutStderr))
		return err
	}

	helpers.AppLogger.Debug("Output of command %s:\n%v", cmd, string(stdoutStderr))

	return nil
}

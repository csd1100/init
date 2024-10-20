package cli

import (
	"github.com/csd1100/init/internal/helpers"
)

type npmExe struct {
	Command string
}

func (npm npmExe) Exec(subcommand string, args []string) error {
	return execute(npm, subcommand, args)
}

func (npm npmExe) GetCommand() string {
	return npm.Command
}

type npmCLI struct {
	exe Executable
}

var Npm = npmCLI{exe: npmExe{Command: "pnpm"}}

func (npm npmCLI) Install() error {
	return npm.exe.Exec("install", []string{})
}

func (npm npmCLI) Sync(data map[string]string) error {
	helpers.AppLogger.Trace("Running pnpm Sync method")
	helpers.AppLogger.Debug("Using options %v for Sync", data)
	return Npm.Install()
}

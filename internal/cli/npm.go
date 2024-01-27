package cli

import (
	"github.com/csd1100/init/internal/helpers"
)

type NpmCLI struct {
	Command string
}

var Npm = NpmCLI{Command: "npm"}

func (npm NpmCLI) GetCommand() string {
	return npm.Command
}

func (npm NpmCLI) Exec(subcommand string, args []string) error {
	return execute(npm, subcommand, args)
}

func (npm NpmCLI) Install() error {
	return npm.Exec("install", []string{})
}

func (npm NpmCLI) Sync(data map[string]string) error {
	helpers.AppLogger.Trace("Running npm Sync method")
	helpers.AppLogger.Debug("Using options %v for Sync", data)
	return Npm.Install()
}

package cli

import "github.com/csd1100/init/internal/helpers"

type NpmCLI struct {
	CLI
}

func (Npm NpmCLI) Install() error {
	op, err := Npm.Exec("install", []string{})
	helpers.AppLogger.Debug("Output of npm install:\n %s", op)
	if err != nil {
		return err
	}
	return nil
}

func (Npm NpmCLI) Sync(data map[string]string) error {
	helpers.AppLogger.Trace("Running npm Sync method")
	helpers.AppLogger.Debug("Using options %v for Sync", data)
	return Npm.Install()
}

var Npm = NpmCLI{CLI{Command: "npm"}}

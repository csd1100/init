package cli

import (
	"github.com/csd1100/init/internal/helpers"
)

type goLang struct {
	Command string
}

var Go = goLang{Command: "go"}

func (goCLI goLang) GetCommand() string {
	return Go.Command
}

func (goCLI goLang) Exec(subcommand string, args []string) error {
	return execute(goCLI, subcommand, args)
}

func (goCLI goLang) ModInit(projectName string) error {
	return Go.Exec("mod", []string{"init", projectName})
}

func (goCLI goLang) ModTidy() error {
	err := Go.Exec("mod", []string{"tidy"})
	if err != nil {
		return err
	}
	return nil
}

func (goCLI goLang) Sync(data map[string]string) error {
	helpers.AppLogger.Trace("Running go Sync method")
	helpers.AppLogger.Debug("Using options %v for Sync", data)

	err := Go.ModInit(data[helpers.PROJECT_NAME])
	if err != nil {
		return err
	}

	err = Go.ModTidy()
	if err != nil {
		return err
	}

	return nil
}

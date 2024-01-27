package cli

import (
	"github.com/csd1100/init/internal/helpers"
)

type goExe struct {
	Command string
}

func (goExe goExe) Exec(subcommand string, args []string) error {
	return execute(goExe, subcommand, args)
}

func (goExe goExe) GetCommand() string {
	return goExe.Command
}

type goLang struct {
	exe Executable
}

var Go = goLang{exe: goExe{Command: "go"}}

func (goCLI goLang) ModInit(projectName string) error {
	return goCLI.exe.Exec("mod", []string{"init", projectName})
}

func (goCLI goLang) ModTidy() error {
	return goCLI.exe.Exec("mod", []string{"tidy"})
}

func (goCLI goLang) Sync(data map[string]string) error {
	helpers.AppLogger.Trace("Running go Sync method")
	helpers.AppLogger.Debug("Using options %v for Sync", data)

	err := goCLI.ModInit(data[helpers.PROJECT_NAME])
	if err != nil {
		return err
	}

	err = goCLI.ModTidy()
	if err != nil {
		return err
	}

	return nil
}

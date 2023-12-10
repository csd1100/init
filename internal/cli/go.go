package cli

import (
	"github.com/csd1100/init/internal/helpers"
)

type goLang struct {
	CLI
}

func (Go goLang) ModInit(projectName string) error {
	op, err := Go.Exec("mod", []string{"init", projectName})
	helpers.AppLogger.Debug("Output of go mod init:\n %s", op)
	if err != nil {
		return err
	}
	return nil
}
func (Go goLang) ModTidy() error {
	op, err := Go.Exec("mod", []string{"tidy"})
	helpers.AppLogger.Debug("Output of go mod tidy:\n %s", op)
	if err != nil {
		return err
	}
	return nil
}

func (Go goLang) Sync(data map[string]string) error {
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

var Go = goLang{CLI{Command: "go"}}

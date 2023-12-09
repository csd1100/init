package cli

import (
	"fmt"
)

type NpmCLI struct {
	CLI
}

func (Npm NpmCLI) Install() error {
	op, err := Npm.Exec("install", []string{})
	fmt.Printf("%s\n", op)
	if err != nil {
		return err
	}
	return nil
}

func (Npm NpmCLI) Sync(data map[string]string) error {
	return Npm.Install()
}

var Npm = NpmCLI{CLI{Command: "npm"}}

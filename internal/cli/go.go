package cli

import (
	"fmt"
)

type goLang struct {
	CLI
}

func (Go goLang) ModInit(projectName string) error {
	op, err := Go.Exec("mod", []string{"init", projectName})
	fmt.Printf("%s", op)
	fmt.Printf("%s", err.Error())
	if err != nil {
		return err
	}
	return nil
}
func (Go goLang) ModTidy() error {
	op, err := Go.Exec("mod", []string{"tidy"})
	fmt.Printf("%s\n", op)
	if err != nil {
		return err
	}
	return nil
}

func (Go goLang) Sync(data map[string]string) error {
	err := Go.ModTidy()
	if err != nil {
		return err
	}
	return nil
}

var Go = goLang{CLI{Command: "go"}}

package core

import (
	"fmt"
	"os"
	"path"

	"github.com/csd1100/init/internal/cli"
	"github.com/csd1100/init/internal/templates"
	"github.com/csd1100/init/internal/utils"
)

var git cli.CLI

func init() {
	git = cli.CLI{Command: "git"}
}

func Init(options utils.Options) error {
	tmpDir, err := createTempDirAndChangeCWD()
	fmt.Println(*tmpDir)
	if err != nil {
		fmt.Println(err.Error())
	}

	err = cloneTemplateRepoAndChangeCWD(options)
	if err != nil {
		fmt.Println("err", err.Error())
	}

	err = options.Template.ParseTemplates()
	if err != nil {
		fmt.Println("err", err.Error())
	}

	// 3. git init
	// 4. run Init on template
	return nil
}

func createTempDirAndChangeCWD() (*string, error) {
	tmpDir, err := os.MkdirTemp(os.TempDir(), "init-*")
	if err != nil {
		fmt.Println("er", err.Error())
		return nil, err
	}
	fmt.Println(tmpDir)

	err = os.Chdir(tmpDir)
	if err != nil {
		return nil, err
	}
	return &tmpDir, nil
}

func cloneTemplateRepoAndChangeCWD(options utils.Options) error {
	err := cli.GitCloneSingleBranch(git, "https://github.com/csd1100/templates/", options.Template.(templates.Template).Name)

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}
	err = os.Chdir(path.Join(cwd, "templates"))
	if err != nil {
		return err
	}
	return err
}

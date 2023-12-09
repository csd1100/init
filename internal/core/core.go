package core

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/csd1100/init/internal/cli"
	"github.com/csd1100/init/internal/templates"
	"github.com/csd1100/init/internal/utils"
)

var git cli.CLI

func init() {
	git = cli.CLI{Command: "git"}
}

func Init(options utils.Options) error {
	projectPath, err := os.Getwd()
	if err != nil {
		fmt.Println(err.Error())
	}
	if options.Path != "" {
		projectPath = options.Path
	}
	projectAbsPath, err := filepath.Abs(projectPath)
	if err != nil {
		fmt.Println("err", err.Error())
	}

	projectAbsPath = path.Join(projectAbsPath, options.Name)

	fmt.Println(projectPath)
	fmt.Println(projectAbsPath)

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

	os.RemoveAll(".git")
	os.RemoveAll("templates")

	// 3. git init
	if !options.NoGit {
		err = cli.GitInit(git)
		if err != nil {
			fmt.Println("err", err.Error())
		}
	}

	// 4. run Init on template
	if !options.NoSync {
	}

	os.Rename(path.Join(*tmpDir, "templates"), projectAbsPath)
	if err != nil {
		fmt.Println("err", err.Error())
	}

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

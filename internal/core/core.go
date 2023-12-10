package core

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/csd1100/init/internal/cli"
	"github.com/csd1100/init/internal/templates"
	"github.com/csd1100/init/internal/utils"
)

func Init(options utils.Options) error {
	projectAbsPath, err := getProjectAbsolutePath(options)
	if err != nil {
		return errors.New(fmt.Sprintf("Error Generating Project Path, Err:%v", err.Error()))
	}

	tmpDir, err := setupRepo(options)
	if err != nil {
		return err
	}

	// 2. Parse templates and cleanup directory
	err = options.Template.ParseTemplates()
	if err != nil {
		return errors.New(fmt.Sprintf("Error while Parsing a Template, Err:%v", err.Error()))
	}

	err = cleanupProjectDir()
	if err != nil {
		return err
	}

	// 3. git init
	if !options.NoGit {
		err = cli.Git.Init()
		if err != nil {
			return errors.New(fmt.Sprintf("Error while running git init, Err:%v", err.Error()))
		}
	}

	// 4. run Init on template
	if !options.NoSync {
		err = options.Template.Sync(options.Template.(templates.Template).TemplateData)
		if err != nil {
			return errors.New(fmt.Sprintf("Error while running sync, Err:%v", err.Error()))
		}
	}

	os.Rename(path.Join(tmpDir, "templates"), projectAbsPath)
	if err != nil {
		return errors.New(fmt.Sprintf("Error while moving the project, Err:%v", err.Error()))
	}

	return nil
}

func getProjectAbsolutePath(options utils.Options) (string, error) {
	projectPath, err := os.Getwd()
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error Getting Current Directory, Err:%v", err.Error()))
	}
	if options.Path != "" {
		projectPath = options.Path
	}
	projectAbsPath, err := filepath.Abs(projectPath)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error Converting Path to Absolute Path, Err:%v", err.Error()))
	}

	return path.Join(projectAbsPath, options.Name), nil
}

func createTempDirAndChangeCWD() (*string, error) {
	tmpDir, err := os.MkdirTemp(os.TempDir(), "init-*")
	if err != nil {
		return nil, err
	}

	err = os.Chdir(tmpDir)
	if err != nil {
		return nil, err
	}
	return &tmpDir, nil
}

func cloneTemplateRepoAndChangeCWD(options utils.Options) error {
	err := cli.Git.CloneSingleBranch("https://github.com/csd1100/templates/", options.Template.(templates.Template).Name)

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

func cleanupProjectDir() error {
	err := os.RemoveAll(".git")
	if err != nil {
		return errors.New("Error removing .git directory")
	}
	err = os.RemoveAll("templates")
	if err != nil {
		return errors.New("Error removing templates directory")
	}
	return nil
}

func setupRepo(options utils.Options) (string, error) {

	// 1. create temo Directory and clone repo in it
	tmpDir, err := createTempDirAndChangeCWD()
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error Creating a Temporary Directory, Err:%v", err.Error()))
	}

	err = cloneTemplateRepoAndChangeCWD(options)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error while Cloning a Template, Err:%v", err.Error()))
	}

	return *tmpDir, nil
}

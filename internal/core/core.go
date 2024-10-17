package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/csd1100/init/internal/cli"
	"github.com/csd1100/init/internal/helpers"
	"github.com/csd1100/init/internal/templates"
	"github.com/csd1100/init/internal/utils"
)

func shouldGitInit(options utils.Options, projectPath string) bool {
	_, err := os.ReadDir(filepath.Join(projectPath, ".git"))
	return !options.NoGit && err != nil
}

func cleanupProjectDir(current bool) error {
	if !current {
		err := os.RemoveAll(".git")
		if err != nil {
			return errors.New("got an error while removing .git directory")
		}
	}

	err := os.RemoveAll("templates")
	if err != nil {
		return errors.New("got an error while removing templates directory")
	}

	err = os.Remove(helpers.TemplatesFilesConfig)
	if err != nil {
		return errors.New("error removing " + helpers.TemplatesFilesConfig)
	}
	return nil
}

func cloneTemplateRepoAndChangeCWD(options utils.Options) error {
	repo := "https://github.com/csd1100/templates/"
	branch := options.Template.(*templates.Template).Name

	if dev, isPresent := os.LookupEnv(helpers.DEV); isPresent && dev == "true" {
		if repoPath, isPresent := os.LookupEnv(helpers.InitDevRepoPath); isPresent {
			repoAbsPath, err := filepath.Abs(repoPath)
			if err != nil {
				helpers.AppLogger.Error("Invalid Git Repository Path: %s set for %s",
					repoPath, helpers.InitDevRepoPath)
			} else {
				helpers.AppLogger.Info("Using Local Git Repository: %s", repoAbsPath)
				if strings.HasSuffix(repoAbsPath, ".git") {
					repo = fmt.Sprintf("file://%s", repoAbsPath)
				} else {
					repo = fmt.Sprintf("file://%s.git", repoAbsPath)
				}
			}
		}
		if branchName, isPresent := os.LookupEnv(helpers.InitDevBranchName); isPresent {
			helpers.AppLogger.Info("Using Git Branch: %s from %s Repository", branchName, repo)
			branch = branchName
		}
	}

	err := cli.Git.CloneSingleBranch(repo, branch)

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

func setupRepo(options utils.Options) (string, error) {

	// 1. Create temp Directory and
	tmpDir, err := createTempDirAndChangeCWD()
	if err != nil {
		return "", errors.New(fmt.Sprintf("error Creating a Temporary Directory, Err:%v", err.Error()))
	}

	// 2. Clone repo in it and change into that directory
	err = cloneTemplateRepoAndChangeCWD(options)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error while Cloning a Template, Err:%v", err.Error()))
	}

	// 3. Parse template-files.json
	var templateFiles templates.TemplateFiles
	contents, err := os.ReadFile(helpers.TemplatesFilesConfig)
	if err != nil {
		return "", fmt.Errorf("unable to read config file '%v'", helpers.TemplatesFilesConfig)
	}

	err = json.Unmarshal(contents, &templateFiles)
	if err != nil {
		return "", fmt.Errorf("unable to parse config '%v', due to error: '%w'", helpers.TemplatesFilesConfig, err)
	}

	options.Template.(*templates.Template).TemplateFiles = templateFiles

	return *tmpDir, nil
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

	if options.Current {
		return projectAbsPath, nil
	}

	return path.Join(projectAbsPath, options.Name), nil
}

func Init(options utils.Options) error {
	projectAbsPath, err := getProjectAbsolutePath(options)
	if err != nil {
		return errors.New(fmt.Sprintf("Error Generating Project Path, Err:%v", err.Error()))
	}

	if _, err := os.Stat(projectAbsPath); err == nil {
		return errors.New(fmt.Sprintf("'%s' is already present.", projectAbsPath))
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

	err = cleanupProjectDir(options.Current)
	if err != nil {
		return err
	}

	// 3. git init
	if shouldGitInit(options, projectAbsPath) {
		err = cli.Git.Init()
		if err != nil {
			return errors.New(fmt.Sprintf("Error while running git init, Err:%v", err.Error()))
		}
	}

	// 4. run Init on template
	if !options.NoSync {
		err = options.Template.Sync(options.Template.(*templates.Template).TemplateData)
		if err != nil {
			return errors.New(fmt.Sprintf("Error while running sync, Err:%v", err.Error()))
		}
	}

	err = helpers.MoveDir(path.Join(tmpDir, "templates"), projectAbsPath)
	if err != nil {
		return errors.New(fmt.Sprintf("Error while moving the project, Err:%v", err.Error()))
	}

	return os.RemoveAll(tmpDir)
}

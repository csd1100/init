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
	repo := "https://github.com/csd1100/templates/"
	branch := options.Template.(*templates.Template).Name

	if dev, isPresent := os.LookupEnv(helpers.DEV); isPresent && dev == "true" {
		if repo_path, isPresent := os.LookupEnv(helpers.INIT_DEV_REPO_PATH); isPresent {
			repo_abs_path, err := filepath.Abs(repo_path)
			if err != nil {
				helpers.AppLogger.Error("Invalid Git Repository Path: %s set for %s",
					repo_path, helpers.INIT_DEV_REPO_PATH)
			} else {
				helpers.AppLogger.Info("Using Local Git Repository: %s", repo_abs_path)
				if strings.HasSuffix(repo_abs_path, ".git") {
					repo = fmt.Sprintf("file://%s", repo_abs_path)
				} else {
					repo = fmt.Sprintf("file://%s.git", repo_abs_path)
				}
			}
		}
		if branch_name, isPresent := os.LookupEnv(helpers.INIT_DEV_BRANCH_NAME); isPresent {
			helpers.AppLogger.Info("Using Git Branch: %s from %s Repository", branch_name, repo)
			branch = branch_name
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

func cleanupProjectDir() error {
	err := os.RemoveAll(".git")
	if err != nil {
		return errors.New("Error removing .git directory")
	}
	err = os.RemoveAll("templates")
	if err != nil {
		return errors.New("Error removing templates directory")
	}

	err = os.Remove(helpers.TEMPLATES_FILES_CONFIG)
	if err != nil {
		return errors.New("Error removing " + helpers.TEMPLATES_FILES_CONFIG)
	}
	return nil
}

func setupRepo(options utils.Options) (string, error) {

	// 1. Create temp Directory and
	tmpDir, err := createTempDirAndChangeCWD()
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error Creating a Temporary Directory, Err:%v", err.Error()))
	}

	// 2. Clone repo in it and change into that direcotry
	err = cloneTemplateRepoAndChangeCWD(options)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error while Cloning a Template, Err:%v", err.Error()))
	}

	// 3. Parse template-files.json
	var templateFiles templates.TemplateFiles
	contents, err := os.ReadFile(helpers.TEMPLATES_FILES_CONFIG)
	if err != nil {
		return "", fmt.Errorf("Unable to read config file '%v'", helpers.TEMPLATES_FILES_CONFIG)
	}

	err = json.Unmarshal(contents, &templateFiles)
	if err != nil {
		return "", fmt.Errorf("Unable to parse config '%v', due to error: '%w'", helpers.TEMPLATES_FILES_CONFIG, err)
	}

	options.Template.(*templates.Template).TemplateFiles = templateFiles

	return *tmpDir, nil
}

package utils

import (
	"errors"
	"flag"
	"os"
	"path/filepath"

	"github.com/csd1100/init/internal/helpers"
	"github.com/csd1100/init/internal/templates"
)

type Options struct {
	Name     string
	Template templates.Project
	NoGit    bool
	NoSync   bool
	Current  bool
	Path     string
	Help     bool
}

var name string
var templateName string
var noGit bool
var noSync bool
var path string
var current bool
var templateOptions string
var verbosity = -99
var FSet = flag.FlagSet{}

func init() {
	FSet.StringVar(&name, "n", "", "name of the project")
	FSet.StringVar(&name, "name", "", "name of the project")
	FSet.StringVar(&templateName, "t", "", "template for the project")
	FSet.StringVar(&templateName, "template", "", "name for the project")
	FSet.StringVar(&path, "p", "", "path for the project")
	FSet.StringVar(&path, "path", "", "path for the project")
	FSet.StringVar(&templateOptions, "o", "", "options for the template")
	FSet.StringVar(&templateOptions, "options", "", "options for the template")
	FSet.BoolVar(&current, "c", false, "do not create separate directory for project")
	FSet.BoolVar(&current, "current", false, "do not create separate directory for project")
	FSet.BoolVar(&noGit, "G", false, "do not initialize git repository")
	FSet.BoolVar(&noGit, "no-git", false, "do not initialize git repository")
	FSet.BoolVar(&noSync, "S", false, "do not sync project (e.g. npm install, go mod tidy)")
	FSet.BoolVar(&noSync, "no-sync", false, "do not sync project (e.g. npm install, go mod tidy)")
	FSet.IntVar(&verbosity, "v", 2, "verbosity of output from 0 - 5 where 0 is less verbose")
	FSet.IntVar(&verbosity, "verbose", 2, "verbosity of output from 0 - 5 where 0 is less verbose")
}

func validateArgs() error {

	if !current && name == "" {
		return helpers.ErrArgNameRequired
	}

	if name != "" && string(name[0]) == "-" {
		return helpers.ErrInvalidArgName
	}

	if templateName == "" {
		return helpers.ErrArgTemplateRequired
	}

	if current && path != "" {
		return helpers.ErrArgCurrentNotWithPath
	}

	if path != "" {
		_, err := os.ReadDir(path)
		if err != nil {
			return helpers.ErrInvalidArgPath
		}
	}

	return nil
}

func ParseArgs() (*Options, error) {
	helpers.AppLogger.Trace("Starting Parsing Arguments")

	err := FSet.Parse(os.Args[1:])
	if err != nil {
		if errors.Is(err, flag.ErrHelp) {
			return &Options{Help: true}, nil
		} else {
			return nil, err
		}
	}

	err = validateArgs()
	if err != nil {
		return nil, err
	}

	if verbosity != -99 {
		helpers.AppLogger.CurrentLevel = helpers.GetLevel(verbosity)
	}

	if current && name == "" {
		cwd, err := os.Getwd()
		if err != nil {
			helpers.AppLogger.Error("failed to get current directory")
			return nil, err
		}
		name = filepath.Base(cwd)
	}

	template, err := templates.GetTemplate(templateName, name, templateOptions)
	if err != nil {
		return nil, err
	}

	options := Options{
		Name:     name,
		Template: template,
		NoGit:    noGit,
		NoSync:   noSync,
		Path:     path,
		Current:  current,
	}

	return &options, nil
}

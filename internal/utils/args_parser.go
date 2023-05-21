package utils

import (
	"errors"
	"flag"
	"os"

	"github.com/csd1100/init/internal/templates"
)

type Options struct {
	Name     string
	Template templates.Project
	NoGit    bool
	NoSync   bool
	Path     string
	Help     bool
}

var name string
var templateName string
var noGit bool
var noSync bool
var path string
var FSet = flag.FlagSet{}

func init() {
	FSet.StringVar(&name, "n", "", "name of the project")
	FSet.StringVar(&name, "name", "", "name of the project")
	FSet.StringVar(&templateName, "t", "", "template for the project")
	FSet.StringVar(&templateName, "template", "", "name for the project")
	FSet.StringVar(&path, "p", "", "path for the project")
	FSet.StringVar(&path, "path", "", "path for the project")
	FSet.BoolVar(&noGit, "G", false, "do not initialize git repository")
	FSet.BoolVar(&noGit, "no-git", false, "do not initialize git repository")
	FSet.BoolVar(&noSync, "S", false, "do not sync project (e.g. npm install, go mod tidy)")
	FSet.BoolVar(&noSync, "no-sync", false, "do not sync project (e.g. npm install, go mod tidy)")
}

func validateArgs() error {

	if name == "" {
		return ErrArgNameRequired
	}

	if string(name[0]) == "-" {
		return ErrInvalidArgName
	}

	if templateName == "" {
		return ErrArgTemplateRequired
	}

	if path != "" {
		_, err := os.ReadDir(path)
		if err != nil {
			return ErrInvalidArgPath
		}
	}

	return nil
}

func ParseArgs() (*Options, error) {
	err := FSet.Parse(os.Args[1:])
	if err != nil {
		// TODO: throw error if invalid flag and refactor tests accordingly
		if errors.Is(err, flag.ErrHelp) {
			return &Options{Help: true}, nil
		}
	}

	err = validateArgs()
	if err != nil {
		return nil, err
	}

	template, err := templates.GetTemplate(templateName)
	if err != nil {
		return nil, err
	}

	options := Options{
		Name:     name,
		Template: template,
		NoGit:    noGit,
		NoSync:   noSync,
		Path:     path,
	}

	return &options, nil
}

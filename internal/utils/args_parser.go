package utils

import (
	"fmt"

	"github.com/csd1100/init/internal/templates"
)

type Options struct {
	Name     string
	Template templates.Template
	NoGit    bool
	NoSync   bool
	Path     string
	Help     bool
}

func validateArgs(options Options) error {
	if options.Name == "" {
		return fmt.Errorf("expected argument: name")
	}
	if options.Template.Name == "" {
		return fmt.Errorf("expected argument: template")
	}
	return nil
}

func ParseArgs(args []string) (*Options, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("expected at least 2 arguments: name and template")
	}
	options := Options{}
	for i := 0; i < len(args); i++ {
		switch {
		case args[i] == "-n", args[i] == "--name":
			options.Name = args[i+1]
			i++
		case args[i] == "-t", args[i] == "--template":
			template, err := templates.NewTemplate(args[i+1])
			if err != nil {
				return nil, err
			}
			options.Template = *template
			i++
		case args[i] == "-p", args[i] == "--path":
			options.Path = args[i+1]
			i++
		case args[i] == "-h", args[i] == "--help":
			options.Help = true
		case args[i] == "-G", args[i] == "--no-git":
			options.NoGit = true
		case args[i] == "-S", args[i] == "--no-sync":
			options.NoSync = true
		default:
			return nil, fmt.Errorf("invalid argument: %s", args[i])
		}
	}
	if err := validateArgs(options); err != nil {
		return nil, err
	}
	return &options, nil
}

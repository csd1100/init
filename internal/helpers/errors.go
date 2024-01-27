package helpers

import "errors"

var ErrExecNotFound = errors.New("Executable not found in path")
var ErrInvalidArgTemplate = errors.New("invalid template")
var ErrArgNameRequired = errors.New("The name for the project is not passed")
var ErrInvalidArgName = errors.New("Invalid name provided for project")
var ErrInvalidArgPath = errors.New("Invalid path provided for project")
var ErrArgTemplateRequired = errors.New("Template for the project is required")
var ErrNotYetImplmented = errors.New("Not Yet Implemented")

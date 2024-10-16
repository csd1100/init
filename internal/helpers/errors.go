package helpers

import "errors"

var ErrExecNotFound = errors.New("executable not found in path")
var ErrInvalidArgTemplate = errors.New("invalid template")
var ErrArgNameRequired = errors.New("the name for the project is not passed")
var ErrInvalidArgName = errors.New("invalid name provided for project")
var ErrInvalidArgPath = errors.New("invalid path provided for project")
var ErrArgTemplateRequired = errors.New("template for the project is required")

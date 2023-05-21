package utils

import "errors"

var ErrArgNameRequired = errors.New("name for project required")
var ErrInvalidArgName = errors.New("invalid name for project")
var ErrInvalidArgPath = errors.New("invalid path for project")
var ErrArgTemplateRequired = errors.New("template for project required")
var ErrNotYetImplmented = errors.New("not yet implemented")

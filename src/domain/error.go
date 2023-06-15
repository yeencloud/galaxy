package domain

import "errors"

var ErrInvalidMethodName = errors.New("Invalid method name")
var ErrNoComponentsToRegister = errors.New("No components to register")
var ErrMultipleModuleExports = errors.New("Multiple module exports are not supported")
var ErrApiVersionMismatch = errors.New("API version mismatch")
var ErrHealthFunctionNotExported = errors.New("Health function not exported")
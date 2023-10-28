package commands

import (
	"errors"
)

var (
	NoDataErr         = errors.New("user has no data ")
	UserNotAuthorized = errors.New("user not authorized")
	UnknownDataType   = errors.New("unknown data type")
)

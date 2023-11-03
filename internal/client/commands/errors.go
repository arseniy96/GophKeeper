package commands

import (
	"errors"
)

var (
	ErrNoData            = errors.New("user has no data")
	ErrUserNotAuthorized = errors.New("user not authorized")
	ErrUnknownDataType   = errors.New("unknown data type")
)

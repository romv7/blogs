package errors

import "errors"

var (
	ErrInsufficientArguments = errors.New("insufficient number of arguments provided.")
)

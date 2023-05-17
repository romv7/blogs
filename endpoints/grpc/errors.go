package grpc

import "errors"

var (
	ErrNotEnoughArgument = errors.New("Provided an insufficient number of required arguments.")
)

package store

import (
	"errors"
)

// This type corresponds to a set of enum values that points to store type.
type StoreType uint

const (
	// SqlStore is an enum value that indicates to use an SQL-based database
	// as the main store of the defined instance of Store type.
	SqlStore StoreType = iota
)

var (

	// When an invoker tries to instantiate a Store type using an invalid store
	// this error is returned.
	ErrInvalidStore             = errors.New("invalid store selected")
	ErrUnauthorizedToCreatePost = errors.New("user unauthorized to create post")
)

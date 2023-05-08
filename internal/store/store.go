package store

import (
	"errors"
)

type StoreType uint

const (
	SqlStore StoreType = iota
)

var (
	ErrInvalidStore = errors.New("invalid store selected")
)

package storage

import (
	"errors"
	"log"

	"github.com/romv7/blogs/internal/storage/driver"
	"github.com/romv7/blogs/internal/storage/driver/plain"
)

type StorageDriverType uint

const (
	Plain StorageDriverType = iota
)

type StorageDriver interface {
	//
	Get(key string) (p []byte, err error)

	//
	Put(key string, b []byte) (err error)

	//
	Remove(key string) (err error)

	//
	Describe(key string) (res []*driver.PathInfo, err error)

	//
	Contains(key string) bool
}

var (
	ErrorInvalidStorageDriver = errors.New("invalid storage driver selected")
)

type Storage struct {
	t   StorageDriverType
	drv StorageDriver
}

func NewStorage(t StorageDriverType) (out *Storage) {
	out = &Storage{t, nil}

	switch t {
	case Plain:
		out.drv = plain.Default
	default:
		log.Panic(ErrorInvalidStorageDriver)
	}

	return
}

func NewPlainStorage() StorageDriver {
	return NewStorage(Plain).drv
}

func (s *Storage) Driver() StorageDriver {
	return s.drv
}

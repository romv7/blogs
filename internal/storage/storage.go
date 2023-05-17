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

// This interface is the contract that contains a set of methods that must be satified
// by any implemented storage driver. This will serve as our gateway to any storage
// driver that we had implemented (so far it's only the Plain driver that is available).
type StorageDriver interface {
	// Gets the associated bytes of data with the key argument. If there is no
	// data it should return an error instead.
	Get(key string) (p []byte, err error)

	// Tries to put the data argument to the storage driver using the key arguemnt.
	// When an internal error occured or any other exception occured, an error must
	// be returned
	Put(key string, b []byte) (err error)

	// Removes the key and the data associated to it from the storage driver.
	Remove(key string) (err error)

	// Describes returns a set of PathInfo. Returns an error when something
	// went wrong.
	Describe(key string) (res []*driver.PathInfo, err error)

	// Contains checks whether the key argument exists in the storage driver.
	Contains(key string) bool
}

var (
	// If the end user of the Storage type accidently selected an invalid or unimplemented
	// storage driver, this error is returned.
	ErrorInvalidStorageDriver = errors.New("invalid storage driver selected")
)

// Storage type can have a variety of shape depending on what type of
// storage driver the end user created. To check what are the possible
// drivers that this type can currently be shaped into you can lookup
// the package internal/storage/driver/...
//
// t specifies what is the type of the storage driver that is
// going to be used.
//
// drv contains the instance of the main storage driver that the Storage
// will use.
type Storage struct {
	t   StorageDriverType
	drv StorageDriver
}

// Creates a new Storage type using the Default storage driver defined in the storage driver
// package. Panics when the argument is an invalid storage driver (ErrorInvalidStorageDriver).
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

// Conveniently creates a Storage type that uses the Plain driver.
func NewPlainStorage(rootPath string) StorageDriver {
	return plain.NewPlain(rootPath)
}

// Returns the driver of the Storage instance. Useful for debugging the storage driver
// code.
func (s *Storage) Driver() StorageDriver {
	return s.drv
}

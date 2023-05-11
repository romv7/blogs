package storage

import (
	"errors"
	"log"
	"os"
)

var (
	ErrStorageDirNotSet = errors.New("STORAGE_DIR was not set")

	storageDir = os.Getenv("STORAGE_DIR")
)

func init() {
	if storageDir == "" {
		log.Panic(ErrStorageDirNotSet)
	}
}

type Storage struct{}

func (s *Storage) WriteTo() {

}

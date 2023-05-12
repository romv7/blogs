package storageTest

import (
	"testing"

	"github.com/romv7/blogs/internal/storage"
)

func TestNewStorage(t *testing.T) {
	if S := storage.NewStorage(storage.Plain); S == nil {
		t.Errorf("NewStorage should instantiate a new storage driver")
	}
}

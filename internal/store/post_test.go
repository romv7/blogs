package store_test

import (
	"testing"

	"github.com/rommms07/blogs/internal/store"
)

var (
	postStore *store.PostStore
)

func init() {
	postStore = store.NewPostStore(globalFakeUser)
}

func Test_theTestPostStoreMustNotBeDefined(t *testing.T) {
	if postStore == nil {
		t.Errorf("[fail] store.NewPostStore failed to defined the test store...")
	}
}

func Test_mustInstantiateANewPost(t *testing.T) {
	post := postStore.New()

	if post.User != globalFakeUser {
		t.Errorf("[fail] postStore.New did not create the expected post.")
	}

}

package store

import (
	"github.com/rommms07/blogs/internal/entities"
)

type PostStore struct {
	user *entities.User
	UnimplementedStore
}

func NewPostStore(user *entities.User) *PostStore {
	return &PostStore{user: user}
}

func (s *PostStore) New() *entities.Post {
	return entities.NewPost(s.user)
}

func (s *PostStore) Read(query string) {

}

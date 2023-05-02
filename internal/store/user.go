package store

import (
	"github.com/romv7/blogs/internal/entities"
)

type UserStore struct {
	UnimplementedStore
}

func NewUserStore() *UserStore {
	return &UserStore{}
}

func (s *UserStore) New(name, fullName, email string) *entities.User {
	return entities.NewUser(name, fullName, email)
}

package store

import (
	"github.com/romv7/blogs/internal/entities"
)

type CommentStore struct {
	user *entities.User
	UnimplementedStore
}

func NewCommentStore(user *entities.User) (*CommentStore, error) {
	return &CommentStore{user: user}, nil
}

func (s *CommentStore) New(commentText string) *entities.Comment {
	return nil
}

func (s *CommentStore) Save(p *entities.Comment) error {
	return nil
}

func (s *CommentStore) Delete(p *entities.Comment) error {
	return nil
}

func (s *CommentStore) Read() error {

	return nil
}

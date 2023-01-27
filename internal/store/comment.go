package store

import (
	"github.com/rommms07/blogs/internal/entities"
)

type CommentStore struct {
	user *entities.User
	UnimplementedStore
}

func NewCommentStore(user *entities.User) (*CommentStore, error) {
	return &CommentStore{user: user}, nil
}

func (s *CommentStore) New(commentText, targetUuid string) *entities.Comment {
	return entities.NewComment(s.user, commentText, targetUuid)
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

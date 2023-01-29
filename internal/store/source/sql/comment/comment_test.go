package comment_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/rommms07/blogs/internal/entities"
	"github.com/rommms07/blogs/internal/store/source/sql/comment"
)

var (
	commentsDb    *comment.CommentStoreSql
	mockUser       = entities.NewUser("rommms07", "Rom Vales Villanueva", "romdevmod@gmail.com")
)

func Test_shouldCreateAMockCommentOnBehalfOfAUser(t *testing.T) {
	err := commentsDb.Save(entities.NewComment(mockUser, "The quick brown fox jumps over the lazy dog.", uuid.New().String()))
	if err != nil {
		t.Errorf(err.Error())
	}
}

func Test_ableToEditAMockCommentByWhoCreatedIt(t *testing.T) {
}

func Test_shouldBeAbleToDeleteAMockComment(t *testing.T) {

}

func Test_shouldReturnASetOfCommentsViaReadQuery(t *testing.T) {

}


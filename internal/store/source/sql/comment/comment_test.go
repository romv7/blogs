package comment_test

import (
	"testing"

	"github.com/romv7/blogs/internal/entities"
	"github.com/romv7/blogs/internal/store/source/sql/comment"
)

var (
	commentsDb *comment.CommentStoreSql
	mockUser   = entities.NewUser("romv7", "Rom Vales Villanueva", "romdevmod@gmail.com")
	mockPost   = entities.NewPost(mockUser, "What makes societal collapse deadly?", "Blah blah blah", "test", "test1", "test2")
)

type TestCase struct {
	User        *entities.User
	commentText string
}

type TestCases []*TestCase

var tcases = TestCases{
	{
		User:        entities.NewUser("jodiee", "Jodiee Maria Santos", "jodiee9@gmail.com"),
		commentText: "Your profile is full of personality and love! I like it.",
	},
}

func Test_shouldCreateAMockCommentOnBehalfOfAMockUser(t *testing.T) {
	commentStore := comment.CommentStoreSql{}

	for _, tcase := range tcases {
		comment := entities.NewComment(tcase.User, tcase.commentText, mockPost)
		commentStore.T = comment.State.CreatedAt.AsTime().Unix()
		if err := commentStore.Save(comment); err != nil {
			t.Fatalf(err.Error())
		}

	}

}

func Test_ableToEditAMockCommentByWhoCreatedIt(t *testing.T) {
}

func Test_shouldBeAbleToDeleteAMockComment(t *testing.T) {

}

func Test_shouldReturnASetOfCommentsViaReadQuery(t *testing.T) {

}

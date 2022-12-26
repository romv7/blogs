package comment_test

import (
	"testing"

	"github.com/rommms07/blogs/internal"
	"github.com/rommms07/blogs/internal/entities"
	"github.com/rommms07/blogs/internal/store/source/sql"
	"github.com/rommms07/blogs/internal/store/source/sql/comment"
	"github.com/rommms07/blogs/pb"
)

var (
	testCommentsDb *sql.SQLDataSource
	comment_ins    *comment.CommentStoreSql
	mockUser       = entities.NewUser("rommms07", "Rom Vales Villanueva", "romdevmod@gmail.com")

	config *internal.ConfigSchema
)

func init() {
	config, _ = internal.LoadConfig()
	config.Setenv("test")

	testCommentsDb = sql.NewSQLDataSource(config.Database.Drv_name, "")
	testCommentsDb.InitWithMockDb("commentsDb")

	comment_ins = comment.NewSQLCommentStore(testCommentsDb)
}

func Test_shouldCreateAMockCommentOnBehalfOfAUser(t *testing.T) {
	// expectsTbl :=

	err := comment_ins.Save(entities.NewComment(mockUser, "The quick brown fox jumps over the lazy dog.", 1, pb.Comment_T_POST))
	if err != nil {
		t.Errorf(err.Error())
	}
}

func Test_ableToEditAMockCommentByWhoCreatedIt(t *testing.T) {
	// expectsTbl :=
	// mockComment := &entities.Comment{}

}

func Test_shouldBeAbleToDeleteAMockComment(t *testing.T) {

}

func Test_shouldReturnASetOfCommentsViaReadQuery(t *testing.T) {

}

func Test_cleanUp(t *testing.T) {
	testCommentsDb.DetachIfMock()
}

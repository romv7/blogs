package post_test

import (
	"testing"

	"github.com/romv7/blogs/internal/entities"
	"github.com/romv7/blogs/internal/store/source/sql/post"
	"github.com/romv7/blogs/internal/store/source/sql/uuidindexes"
)

type TestCase struct {
	User *entities.User
	Post *entities.Post
}

type TestCases = []*TestCase

func Test_shouldCreateNewMockPostOnBehalfOfAFakeUser(t *testing.T) {
	post.InitSql()

	me := entities.NewUser("romv7", "Rom Vales Villanueva", "romdevmod@gmail.com")

	tcases := TestCases{
		{
			User: me,
			Post: entities.NewPost(me, "How to Become Obese In Just One Month!", ""),
		},
	}

	postStore := post.PostStoreSql{}

	for _, tcase := range tcases {
		postStore.T = tcase.Post.State.CreatedAt.AsTime().Unix()
		if err := postStore.Save(tcase.Post); err != nil {
			t.Fatalf(err.Error())
		}

		ui, err := uuidindexes.GetUuidIndex(post.DbName(), tcase.Post.UniqueKey())
		if err != nil {
			t.Fatalf(err.Error())
		}

		if !ui.Exists() {
			t.Fatalf("unable to save post \"%s\"", tcase.Post.HeadlineText)
		}

		if err := postStore.Delete(tcase.Post); err != nil {
			t.Fatalf(err.Error())
		}
	}
}

func Test_ableToEditAMockPostByUserABCWhoCreatedIt(t *testing.T) {
	post.InitSql()
}

func Test_shouldBeAbleToDeleteAMockPostCreatedByAUser(t *testing.T) {
	post.InitSql()
}

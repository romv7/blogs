package storeTest

import (
	"testing"

	"github.com/romv7/blogs/internal/store"
)

func TestNewPostStore(t *testing.T) {
	if store.NewPostStore(store.SqlStore) == nil {
		t.Fatalf("NewPostStore returned an invalid store")
	}
}

func TestPostStore_SqlStore(t *testing.T) {

	pstore := store.NewPostStore(store.SqlStore)
	ustore := store.NewUserStore(store.SqlStore)

	for _, tcase := range globalPostTestCases {
		u := ustore.NewUser(tcase.u)
		p := pstore.NewPost(tcase.u, tcase.p)

		// Save the user from the test case.
		if err := ustore.Save(u); err != nil {
			t.Fatal(err)
		}

		// Save the blog post to the database.
		if err := pstore.Save(p); err != nil {
			t.Fatal(err)
		}

		defer ustore.Delete(u)
		defer pstore.Delete(p)

		// Expect that the post `p` should now be in the db.
		if ex, err := pstore.GetByUuid(p.Proto().Uuid); err != nil {
			t.Fatal(err)
		} else {

			if ex.Proto().Id != p.Proto().Id {
				t.Error(ErrPropNotMatched)
			}

			if ex.Proto().HeadlineText != p.Proto().HeadlineText {
				t.Error(ErrPropNotMatched)
			}

			if ex.Proto().User.Id != u.Proto().Id {
				t.Error(ErrPropNotMatched)
			}
		}
	}

}

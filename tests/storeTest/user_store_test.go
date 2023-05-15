package storeTest

import (
	"testing"

	"github.com/romv7/blogs/internal/store"
)

func TestNewUserStore(t *testing.T) {
	if store.NewUserStore(store.SqlStore) == nil {
		t.Fatalf("NewUserStore returned an invalid store.")
	}
}

func TestUserStore_SqlStore(t *testing.T) {

	ustore := store.NewUserStore(store.SqlStore)

	for _, tcase := range globalUserTestCases {
		u := ustore.NewUser(tcase.u)
		if err := u.Save(); err != nil {
			t.Fatal(err)
		}

		defer u.Delete()

		var ex *store.User

		if expected, err := ustore.GetByUuid(u.Proto().Uuid); err != nil {
			t.Error(err)
		} else {
			ex = expected
		}

		if u.Proto().Id != ex.Proto().Id {
			t.Error(ErrPropNotMatched)
			t.Error("expected model Id did not match with the test case model Id")
			t.Error(ex.Proto().Id, " != ", u.Proto().Id)
		}

		if ex.Proto().Uuid != u.Proto().Uuid {
			t.Error(ErrPropNotMatched)
			t.Error("expected model Uuid did not match with the test case model Uuid")
			t.Error(ex.Proto().Uuid, " != ", u.Proto().Uuid)
		}

		if ex.Proto().FullName != u.Proto().FullName {
			t.Error(ErrPropNotMatched)
			t.Error("expected model FullName did not match with the test case model FullName")
		}
	}
}

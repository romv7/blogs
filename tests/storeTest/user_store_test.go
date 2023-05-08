package storeTest

import (
	"testing"

	"github.com/romv7/blogs/internal/store"
)

func TestNewUserStore(t *testing.T) {
	if store.NewUserStore(store.SqlStore) == nil {
		t.Fatalf("NewUserStore returned an invalid store. =What the fuck?")
	}
}

func TestUserStore_SqlStore(t *testing.T) {
	ustore := store.NewUserStore(store.SqlStore)

	for _, tcase := range globalUserTestCases {
		u := ustore.NewUser(tcase.u)
		if err := ustore.Save(u); err != nil {
			t.Fatal(err)
		}

		defer ustore.Delete(u)

		var ex *store.User

		if expected, err := ustore.GetByUuid(u.Proto().Uuid); err != nil {
			t.Error(err)
		} else {
			ex = expected
		}

		if u.Proto().Id != ex.Proto().Id {
			t.Error(ErrPropNotMatched)
			t.Error("expected model Id did not match with the test case model Id")
		}

		if ex.Proto().FullName != u.Proto().FullName {
			t.Error(ErrPropNotMatched)
			t.Error("expected model FullName did not match with the test case model FullName")
		}
	}
}

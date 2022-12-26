package store_test

import (
	"testing"

	"github.com/rommms07/blogs/internal/entities"
	"github.com/rommms07/blogs/internal/store"
	"github.com/rommms07/blogs/pb"
)

var (
	userStore *store.UserStore
)

func init() {
	userStore = store.NewUserStore()
}

func Test_theTestUserStoreMustNowBeDefined(t *testing.T) {
	if userStore == nil {
		t.Errorf("[fail] store.NewUserStore did not properly instantiate the test store...")
	}
}

func Test_mustInstantiateANewUserEntity(t *testing.T) {
	me := userStore.New("rommms07", "Rom Vales Villanueva", "romdevmod@gmail.com")
	if me == nil {
		t.Errorf("[fail] userStore.New was not able to create a new user entity.")
	}

	expect := &entities.User{
		User: &pb.User{
			Name:     "rommms07",
			FullName: "Rom Vales Villanueva",
			Email:    "romdevmod@gmail.com",
		},
	}

	if me.Name != expect.Name {
		t.Errorf("[fail] me.Name does not contain the expected value `%s`", expect.Name)
	}

	if me.Email != expect.Email {
		t.Errorf("[fail] me.Email does not contain the expected email `%s`", expect.Email)
	}

	if me.FullName != expect.FullName {
		t.Errorf("[fail] me.FullName does not contain the expected full name `%s`", expect.FullName)
	}
}

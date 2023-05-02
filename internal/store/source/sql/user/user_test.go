package user_test

import (
	"encoding/base64"
	"fmt"
	"testing"

	"github.com/romv7/blogs/internal/entities"
	"github.com/romv7/blogs/internal/store/source/sql/user"
	"github.com/romv7/blogs/internal/store/source/sql/uuidindexes"
)

type TestCase struct {
	User *entities.User
}

type TestCases = []*TestCase

func Test_shouldBeAbleToCreateANewMockUser(t *testing.T) {
	user.InitSql()

	tcases := TestCases{
		{
			User: entities.NewUser("romv7", "Rom Vales Villanueva", "romdevmod@gmail.com"),
		},
		{
			User: entities.NewUser("jackf", "Jack Frost", "jackfrostt@yahoo.com.ph"),
		},
		{
			User: entities.NewUser("cidhighwind", "Cid Highwind", "cid@ff.com"),
		},
	}

	userStore := &user.UserStoreSql{}

	for _, tcase := range tcases {
		userStore.T = tcase.User.State.CreatedAt.AsTime().Unix()
		userStore.Save(tcase.User)

		ui, err := uuidindexes.GetUuidIndex("users", tcase.User.UniqueKey())
		if err != nil {
			t.Errorf(err.Error())
			t.Fatalf("user was not possible pushed to the db")
		}

		if ui.Resource_key != fmt.Sprintf("users:%s", base64.StdEncoding.EncodeToString([]byte(tcase.User.UniqueKey()))) {
			t.Fatalf("did not matched the uuid of the user that was pushed to the db")
		}

		userStore.Delete(tcase.User)
	}

}

func Test_mustBeAbleToEditAnExistingMockUser(t *testing.T) {
	user.InitSql()

}

func Test_ableToRemoveOrDeleteAMockUser(t *testing.T) {
	user.InitSql()

	tcases := TestCases{
		{
			User: entities.NewUser("romv7", "Rom Vales Villanueva", "romdevmod@gmail.com"),
		},
	}

	userStore := &user.UserStoreSql{}

	for _, tcase := range tcases {
		userStore.T = tcase.User.State.CreatedAt.AsTime().Unix()
		userStore.Save(tcase.User)

		ui, err := uuidindexes.GetUuidIndex("users", tcase.User.UniqueKey())
		if err != nil {
			t.Errorf(err.Error())
			t.Fatalf("user was not possible pushed to the db")
		}

		if !ui.Exists() {
			t.Fatalf("user %s does not exist in the db", tcase.User.UniqueKey())
		}

		userStore.Delete(tcase.User)

		if ui.Exists() {
			t.Fatalf("unable to delete %s from the db", tcase.User.UniqueKey())
		}
	}
}

func Test_ableToRemoveOrDeleteAMockUserById(t *testing.T) {
	t.Errorf("DeleteById: not implemented")
}

func Test_ableToRemoveOrDeleteAMockUserByUuid(t *testing.T) {
	t.Errorf("DeleteByUuid: not implemented")
}

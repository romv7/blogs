package mutationsTest

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/romv7/blogs/endpoints/gql"
	"github.com/romv7/blogs/internal/pb"
	"github.com/romv7/blogs/internal/store"
	"github.com/romv7/blogs/internal/utils"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

func TestSuperOpsAddUser(t *testing.T) {
	S := gql.DefaultSchema()
	query := `
	mutation SuperOpsForAddUser($new: UserSuperMutationsParameter!) {
		superOps {
			addUser(input: $new) {
				op
				message
				code
				uuid
			}
		}
	}
	`

	var (
		expectedName     = "romdevmod"
		expectedFullName = "Rom Vales Villanueva"
		expectedEmail    = "rom@gmail.com"
	)

	result := S.Exec(context.TODO(), query, "SuperOpsForAddUser", map[string]any{
		"new": map[string]any{
			"name":     expectedName,
			"fullName": expectedFullName,
			"email":    expectedEmail,
		},
	})

	if len(result.Errors) > 0 {
		t.Errorf("%+v", result.Errors)
	}

	opResult := map[string]any{}
	if err := json.Unmarshal(result.Data, &opResult); err != nil {
		t.Fatal(err)
	}

	opResult = opResult["superOps"].(map[string]any)["addUser"].(map[string]any)

	ustore := store.NewUserStore(store.SqlStore)
	uuid := opResult["uuid"].(string)

	if opResult["code"].(float64) != float64(http.StatusCreated) {
		t.Error("failed to create user, returned an unexpected code.")
	}

	var storedUser *store.User

	if u, err := ustore.GetByUuid(uuid); err != nil {
		t.Error(err)
	} else {
		storedUser = u
	}

	defer storedUser.Delete()

	if storedUser.Proto().FullName != expectedFullName {
		t.Error("did not match the expected fullName by the created user.")
	}

	if storedUser.Proto().Name != expectedName {
		t.Error("did not match the expected fullName by the created user.")
	}

	if storedUser.Proto().Email != expectedEmail {
		t.Error("did not match the expected email by the created user.")
	}
}

func TestSuperOpsUpdateUser(t *testing.T) {
	ustore := store.NewUserStore(store.SqlStore)
	S := gql.DefaultSchema()
	query := `
	mutation SuperOpsForUpdateUser($updated: UserSuperMutationsParameter!) {
		superOps {
			updateUser(input: $updated) {
				op
				code
				uuid
			}
		}
	}
	`

	u := ustore.NewUser(&pb.User{
		Id:       utils.RandomUniqueId(),
		Uuid:     uuid.NewString(),
		Name:     "rommms",
		FullName: "Rom Vales Villanueva",
		Email:    "romdevmod@gmail.com",
		Type:     pb.User_T_AUTHOR,
		State: &pb.UserState{
			CreatedAt: timestamppb.Now(),
			UpdatedAt: timestamppb.Now(),
			Disabled:  false,
			UVerified: false,
		},
	})

	if err := u.Save(); err != nil {
		t.Error(err)
	}

	defer u.Delete()

	var (
		expectedName = "romdevmod"
	)

	result := S.Exec(context.TODO(), query, "SuperOpsForUpdateUser", map[string]any{
		"updated": map[string]any{
			"id":         float64(u.Proto().Id),
			"name":       expectedName,
			"email":      "idream.rommms@gmail.com",
			"isDisabled": false,
			"isVerified": true,
		},
	})

	if len(result.Errors) > 0 {
		t.Errorf("%+v", result.Errors)
	}

	var updatedUser *store.User

	if up, err := ustore.GetById(u.Proto().Id); err != nil {
		t.Error(err)
	} else {
		updatedUser = up
	}

	updatedpb := updatedUser.Proto()
	upb := u.Proto()

	if updatedpb.Id != upb.Id {
		t.Error("did not match the expected id by the updated user.")
	}

	if updatedpb.Name == upb.Name {
		t.Errorf("the query should have updated the user name with the new one. \"%s\"", expectedName)
	}

}

func TestSuperOpsDeleteUser(t *testing.T) {
	ustore := store.NewUserStore(store.SqlStore)
	S := gql.DefaultSchema()
	query := `
	mutation SuperOpsForDeleteUser($userUuid: String!) {
		superOps {
			deleteUser(uuid: $userUuid) {
				op
				code
				message
				uuid
			}
		}
	}
	`

	u := ustore.NewUser(&pb.User{
		Id:       utils.RandomUniqueId(),
		Uuid:     uuid.NewString(),
		Name:     "rommms",
		FullName: "Rom Vales Villanueva",
		Email:    "romdevmod.delete@gmail.com",
		Type:     pb.User_T_AUTHOR,
		State: &pb.UserState{
			CreatedAt: timestamppb.Now(),
			UpdatedAt: timestamppb.Now(),
			Disabled:  false,
			UVerified: false,
		},
	})

	if err := u.Save(); err != nil {
		t.Error(err)
	}

	result := S.Exec(context.TODO(), query, "SuperOpsForDeleteUser", map[string]any{
		"userUuid": u.Proto().Uuid,
	})

	if len(result.Errors) > 0 {
		t.Errorf("%+v", result.Errors)
	}

	if u, err := ustore.GetByUuid(u.Proto().Uuid); u != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error("user was not deleted from the database..")
	}
}

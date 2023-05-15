package user_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/romv7/blogs/internal/pb"
	"github.com/romv7/blogs/internal/store"
	"github.com/romv7/blogs/internal/utils"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	ustore   = store.NewUserStore(store.SqlStore)
	mockUser = ustore.NewUser(&pb.User{
		Id:       utils.RandomUniqueId() + uint64(time.Now().Unix()),
		Uuid:     uuid.NewString(),
		Name:     "quietfox",
		FullName: "Rom Vales Villanueva",
		Email:    "quitefox@gmail.com",
		Type:     pb.User_T_AUTHOR,
		State: &pb.UserState{
			CreatedAt: timestamppb.Now(),
			UpdatedAt: nil,
			Disabled:  false,
			UVerified: false,
		},
	})
)

func TestShouldUpdateUserType(t *testing.T) {
	mockUser.Save()
	defer mockUser.Delete()

	if err := ustore.Save(mockUser.ToNormal()); err != nil {
		t.Error(err)
	}

	if u, err := ustore.GetById(mockUser.Proto().Id); err != nil {
		t.Error(err)
	} else {

		if u.Proto().Type != pb.User_T_NORMAL {
			t.Errorf("u.Proto().Type did not match the expected enum pb.USER_T_NORMAL")
		}
	}
}

func TestExpectToVerifyAnUnverifiedUser(t *testing.T) {
	ustore.Save(mockUser)
	defer ustore.Delete(mockUser)

	if err := mockUser.Verify(); err != nil {
		t.Error("failed to verify user: ")
		t.Error(err)
	}
}

package gqlmodelsTest

import (
	"testing"

	"github.com/google/uuid"
	"github.com/romv7/blogs/endpoints/gql/models"
	"github.com/romv7/blogs/internal/pb"
	"github.com/romv7/blogs/internal/utils"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestUserProtoFunc(t *testing.T) {
	upb := &pb.User{
		Id:       utils.RandomUniqueId(),
		Uuid:     uuid.NewString(),
		FullName: "Rom Vales Villanueva",
		Name:     "romdevmod",
		Email:    "romdevmod@gmail.com",
		Type:     pb.User_T_AUTHOR,
		State: &pb.UserState{
			CreatedAt: timestamppb.Now(),
			UpdatedAt: timestamppb.Now(),
			Disabled:  false,
			UVerified: true,
		},
	}

	u := models.Proto_GQLModelUser(upb)

	if uint64(u.Id) != upb.Id {
		t.Error("Something is definitely wrong with floats.")
	}
}

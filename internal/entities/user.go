package entities

import (
	"github.com/google/uuid"
	"github.com/romv7/blogs/internal/utils/rand"
	"github.com/romv7/blogs/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type User struct {
	*pb.User
}

func NewUser(name, fullName, email string) (user *User) {
	now := timestamppb.Now()
	N, _ := rand.Rand()

	user = &User{
		User: &pb.User{
			Id:       uint32(now.AsTime().Unix()) + uint32(N),
			Name:     name,
			FullName: fullName,
			Email:    email,
			Type:     pb.User_T_NORMAL,
			Uuid:     uuid.New().String(),
			State: &pb.UserState{
				CreatedAt: now,
				UpdatedAt: now,
				Disabled:  false,
			},
		},
	}

	return
}

func (u *User) UniqueKey() string {
	return u.Email
}

func (u *User) ChangeType(typ pb.User_Type) *User {
	u.Type = typ
	return u
}

func (u *User) ToNormal() *User {
	u.ChangeType(pb.User_T_NORMAL)
	return u
}

func (u *User) ToAuthor() *User {
	u.ChangeType(pb.User_T_AUTHOR)
	return u
}

package entities

import (
	"github.com/rommms07/blogs/pb"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type User struct {
	*pb.User
}

func NewUser(name, fullName, email string) (user *User) {
	user = &User{
		User: &pb.User{
			Name:      name,
			FullName:  fullName,
			Email:     email,
			Type:      pb.User_T_NORMAL,
			Uuid:      uuid.New().String(),
			State:     &pb.UserState {
				CreatedAt: timestamppb.Now(),
				Disabled:  false,
			},
		},
	}

	return
}

func (u *User) Save() error {
	return nil		
}

func (u *User) Delete() error {
	return nil
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

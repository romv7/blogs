package entities

import (
	"time"

	"github.com/rommms07/blogs/pb"
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
			CreatedAt: uint64(time.Now().Unix()),
		},
	}

	return
}

func (u *User) ChangeType(typ pb.User_Type) *User {

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

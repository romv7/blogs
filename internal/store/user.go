package store

import (
	"fmt"
	"log"

	"github.com/romv7/blogs/internal/pb"
)

func (u *User) ToAuthor() *User {
	if u.Proto().Type == pb.User_T_AUTHOR {
		return u
	}

	switch u.t {
	case SqlStore:
		u.sqlModel.Type = pb.User_T_AUTHOR
	default:
		log.Panic(ErrInvalidStore)
	}

	return u
}

func (u *User) ToNormal() *User {

	switch u.t {
	case SqlStore:
		u.sqlModel.Type = pb.User_T_NORMAL
	default:
		log.Panic(ErrInvalidStore)
	}

	return u
}

func (u *User) Verify() (err error) {

	switch u.t {
	case SqlStore:
		ustore := NewUserStore(SqlStore)
		u.sqlModel.Verified = true
		err = ustore.Save(u)
	default:
		log.Panic(ErrInvalidStore)
	}

	return
}

func (u *User) Disable() (err error) {
	if u.Proto().State.Disabled {
		err = fmt.Errorf("%s: already disabled", u.Proto().Name)
	}

	if err == nil {
		err = toggleDisabledProperty(u, true)
	}

	return
}

func (u *User) Enable() (err error) {
	if !u.Proto().State.Disabled {
		err = fmt.Errorf("%s: already enabled", u.Proto().Name)
	}

	if err == nil {
		err = toggleDisabledProperty(u, false)
	}

	return
}

func (u *User) AuthorRootResourceId() string {
	auth := u.Proto()

	if auth.Type != pb.User_T_AUTHOR {
		log.Panic("normal user cannot have an author resource")
	}

	return auth.Uuid
}

func toggleDisabledProperty(u *User, val bool) (err error) {
	switch u.t {
	case SqlStore:
		ustore := NewUserStore(SqlStore)
		u.sqlModel.Disabled = val
		err = ustore.Save(u)
	default:
		log.Panic(ErrInvalidStore)
	}

	return
}

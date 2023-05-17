package store

import (
	"fmt"
	"log"

	"github.com/romv7/blogs/internal/pb"
	"github.com/romv7/blogs/internal/storage"
	"github.com/romv7/blogs/internal/utils/author"
)

// Only applicable for users that are recognized as pb.USER_T_AUTHOR,
// returns the metadata of the author stored in a storage. Take note,
// this method will panic when the User field (u.s) selects an
// invalid storage driver.
func (u *User) Metadata() *author.AuthorInfo {
	if u.Proto().Type != pb.User_T_AUTHOR {
		return nil
	}

	switch u.s {
	case storage.Plain:
		ah := author.NewAuthorHelper(u.Proto(), u.s)
		return ah.GetAuthorMetadata()
	default:
		log.Panic(storage.ErrorInvalidStorageDriver)
	}

	return nil
}

func (u *User) Save() (err error) {
	ustore := NewUserStore(u.t)
	err = ustore.Save(u)

	return
}

func (u *User) Delete() (err error) {
	ustore := NewUserStore(u.t)

	if err = ustore.Delete(u); err != nil {
		return
	}

	if u.Proto().Type == pb.User_T_AUTHOR {
		ah := author.NewAuthorHelper(u.Proto(), storage.Plain)
		err = ah.DeleteAuthorMetadata()
	}

	return
}

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
	if u.Proto().Type == pb.User_T_NORMAL {
		return u
	}

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

func (u *User) IsAuthor() bool {
	if u.Proto().Type == pb.User_T_AUTHOR {
		return true
	}

	return false
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

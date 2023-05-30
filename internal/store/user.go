package store

import (
	"fmt"
	"log"

	"github.com/romv7/blogs/internal/pb"
	"github.com/romv7/blogs/internal/storage"
	"github.com/romv7/blogs/internal/utils/author"
)

// TODO: Has side effects,
func (u *User) Proto() (out *pb.User) {

	// This method will not rely on the u.IsAuthor method to determine whether the user is an author or not
	// because it causes an infinite recursion when we rely on the u.IsAuthor method.
	var isAuthor bool

	switch u.t {
	case SqlStore:
		out = u.sqlModel.Proto()
		isAuthor = out.Type == pb.User_T_AUTHOR
	default:
		log.Panic(ErrInvalidStore)
	}

	if isAuthor {
		out.StoragePath = author.AuthorRootResourceId(out)
		ah := author.NewAuthorHelper(out, storage.Plain)
		metadata := ah.GetAuthorMetadata()

		out.Bio = metadata.Bio
		out.AltName = metadata.AltName
		for plat, links := range metadata.SocialLinks {
			out.SocialLinks[string(plat)] = &pb.SocialLinks{Data: links}
		}
	}

	return
}

// Only applicable for users that are recognized as pb.USER_T_AUTHOR,
// returns the metadata of the author stored in a storage. Take note,
// this method will panic when the User field (u.s) selects an
// invalid storage driver.
func (u *User) Metadata() *author.AuthorInfo {
	if !u.IsAuthor() {
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
	err = ustore.Delete(u)

	return
}

// Convert a user into a just an author, when the user is already a pb.User_T_AUTHOR do nothing and return the user.
func (u *User) ToAuthor() *User {
	if u.IsAuthor() {
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

// Convert a user into a just a normal one, when the user is already a pb.User_T_NORMAL do nothing and return the user.
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

// TODO: Make a periodically executing code that purges all of the unverified and disabled users from the database.

// Toggles the u.UVerified property of a user. Doing so will verify that the user is a legitimate
// resource consumer.
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

// A simple helper method used for checking whether a user is an author or just a normal user.
func (u *User) IsAuthor() bool {
	return u.Proto().Type == pb.User_T_AUTHOR
}

// A helper function used for toggling the u.Disabled property of a user.
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

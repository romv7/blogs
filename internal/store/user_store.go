package store

import (
	"errors"
	"log"

	"github.com/romv7/blogs/internal/pb"
	"github.com/romv7/blogs/internal/storage"
	sqlStore "github.com/romv7/blogs/internal/store/sql"
	sqlModels "github.com/romv7/blogs/internal/store/sql/models"

	"gorm.io/gorm"
)

type User struct {
	t        StoreType
	s        storage.StorageDriverType
	sqlModel *sqlModels.User
}

func (u *User) Proto() *pb.User {

	switch u.t {
	case SqlStore:
		return u.sqlModel.Proto()
	default:
		log.Panic(ErrInvalidStore)
	}

	return nil
}

type UserStore struct {
	t StoreType
	s storage.StorageDriverType
}

func NewUserStore(t StoreType) *UserStore {
	return &UserStore{t, storage.Plain}
}

func (s *UserStore) GetMainStore() (S any) {
	switch s.t {
	case SqlStore:
		S = sqlStore.Store()
	default:
		log.Panic(ErrInvalidStore)
	}

	return
}

func (s *UserStore) NewUser(u *pb.User) (out *User) {
	out = &User{}

	switch s.t {
	case SqlStore:
		out.t = SqlStore
		out.sqlModel = sqlModels.NewUser(u)
	default:
		log.Panic(ErrInvalidStore)
	}

	// Set the storage type
	out.s = s.s

	return
}

func (s *UserStore) Save(u *User) (err error) {

	switch s.t {
	case SqlStore:
		db := sqlStore.Store()
		res := db.Save(u.sqlModel)
		err = res.Error
	default:
		log.Panic(ErrInvalidStore)
	}

	return
}

func (s *UserStore) Delete(u *User) (err error) {

	switch s.t {
	case SqlStore:
		db := sqlStore.Store()

		if res := db.Where("uuid = ?", u.sqlModel.Uuid).Delete(u.sqlModel); res != nil {
			return res.Error
		}

		u = nil
	default:
		log.Panic(ErrInvalidStore)
	}

	return
}

// PROBLEM: Sometimes GetById is returning record not found even though the record exist.

// RESOLVED:
//
//	Turns out that this function is not the problem, the reason why this method
//	is returning a record not found is that the unit test were tightly interconnected
//	as a result, when one test deleted a user and the other tries to GetById the record it
//	is apparent that it should return a not found error because the other test already deleted
//	the record we are trying to lookup with GetById.

func (s *UserStore) GetById(id uint64) (out *User, err error) {
	out = &User{}

	switch s.t {
	case SqlStore:
		db := sqlStore.Store()
		out.t = s.t
		out.sqlModel = &sqlModels.User{ID: id}

		if res := db.First(out.sqlModel); errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, res.Error
		}

	default:
		log.Panic(ErrInvalidStore)
	}

	return
}

func (s *UserStore) GetByUuid(uuid string) (out *User, err error) {
	out = &User{}

	switch s.t {
	case SqlStore:
		db := sqlStore.Store()
		out.t = s.t
		out.sqlModel = &sqlModels.User{}

		if res := db.Where("uuid = ?", uuid).Limit(1).Find(out.sqlModel); errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, res.Error
		}

	default:
		log.Panic(ErrInvalidStore)
	}

	return
}

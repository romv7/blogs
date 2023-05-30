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

type Post struct {
	t         StoreType
	s         storage.StorageDriverType
	isUpdated bool
	sqlModel  *sqlModels.Post
}

type PostStore struct {
	t StoreType
	s storage.StorageDriverType
}

func NewPostStore(t StoreType) *PostStore {
	return &PostStore{t, storage.Plain}
}

func (s *PostStore) GetMainStore() (S any) {
	switch s.t {
	case SqlStore:
		S = sqlStore.Store()
	default:
		log.Panic(ErrInvalidStore)
	}

	return
}

func (s *PostStore) NewPost(u *pb.User, p *pb.Post) (out *Post) {
	out = &Post{}

	if p.User == nil {
		p.User = u
	}

	switch s.t {
	case SqlStore:
		out.t = s.t
		out.sqlModel = sqlModels.NewPost(p)
	default:
		log.Panic(ErrInvalidStore)
	}

	// Set the storage type.
	out.s = s.s

	return
}

func (s *PostStore) Save(p *Post) (err error) {
	if p.Proto().User.Type == pb.User_T_NORMAL {
		return ErrUnauthorizedToCreatePost
	}

	switch s.t {
	case SqlStore:
		db := sqlStore.Store()
		res := db.Save(p.sqlModel)
		err = res.Error
	default:
		log.Panic(ErrInvalidStore)
	}

	return
}

func (s *PostStore) Delete(p *Post) (err error) {

	switch s.t {
	case SqlStore:
		db := sqlStore.Store()

		if res := db.Where("uuid = ?", p.sqlModel.Uuid).Delete(p.sqlModel); res.Error != nil {
			return res.Error
		}

		p = nil
	default:
		log.Panic(ErrInvalidStore)
	}

	return
}

// TODO: Add a documentation to this method.
func (s *PostStore) GetById(id uint64) (out *Post, err error) {
	out = &Post{}

	switch s.t {
	case SqlStore:
		db := sqlStore.Store()
		out.t = s.t
		out.sqlModel = &sqlModels.Post{ID: id}
		if res := db.First(out.sqlModel); errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, res.Error
		}

		if out.sqlModel.ID != id {
			return nil, gorm.ErrRecordNotFound
		}
	default:
		log.Panic(ErrInvalidStore)
	}

	return
}

// TODO: Add a documentation to this method.
func (s *PostStore) GetByUuid(uuid string) (out *Post, err error) {
	out = &Post{}

	switch s.t {
	case SqlStore:
		db := sqlStore.Store()
		out.t = s.t
		out.sqlModel = &sqlModels.Post{}
		if res := db.Where("uuid = ?", uuid).Limit(1).Find(out.sqlModel); errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, res.Error
		}

		if out.sqlModel.Uuid != uuid {
			return nil, gorm.ErrRecordNotFound
		}
	default:
		log.Panic(ErrInvalidStore)
	}

	return
}

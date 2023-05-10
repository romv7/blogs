package store

import (
	"errors"

	"github.com/romv7/blogs/internal/pb"
	sqlStore "github.com/romv7/blogs/internal/store/sql"
	sqlModels "github.com/romv7/blogs/internal/store/sql/models"

	"gorm.io/gorm"
)

type Post struct {
	t        StoreType
	sqlModel *sqlModels.Post
}

func (p *Post) Proto() *pb.Post {
	ustore := NewUserStore(SqlStore)
	cstore := NewCommentStore(SqlStore)

	switch p.t {
	case SqlStore:
		pout := p.sqlModel.Proto()

		if u, err := ustore.GetById(p.sqlModel.UserId); errors.Is(err, gorm.ErrRecordNotFound) {
			// @TODO (Handle post that has no owner)
		} else {
			pout.User = u.Proto()
		}

		pout.Comments = cstore.TargetCommentProtoTree(pout.Uuid)

		// @TODO (Add an edit history)

		return pout
	default:
		panic(ErrInvalidStore)
	}

}

type PostStore struct {
	t StoreType
}

func NewPostStore(t StoreType) *PostStore {
	return &PostStore{t}
}

func (s *PostStore) NewPost(u *pb.User, p *pb.Post) (out *Post) {
	out = &Post{}

	p.User = u

	switch s.t {
	case SqlStore:
		out.t = s.t
		out.sqlModel = sqlModels.NewPost(p)
	default:
		panic(ErrInvalidStore)
	}

	return
}

func (s *PostStore) Save(p *Post) (err error) {
	switch s.t {
	case SqlStore:
		db := sqlStore.Store()
		res := db.Save(p.sqlModel)
		err = res.Error
	default:
		panic(ErrInvalidStore)
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
		panic(ErrInvalidStore)
	}

	return
}

func (s *PostStore) GetById(id uint32) (out *Post, err error) {
	out = &Post{}

	switch s.t {
	case SqlStore:
		db := sqlStore.Store()
		out.t = s.t
		out.sqlModel = &sqlModels.Post{}
		if res := db.Where("id = ?", id).First(out.sqlModel); errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, res.Error
		}
	default:
		panic(ErrInvalidStore)
	}

	return
}

func (s *PostStore) GetByUuid(uuid string) (out *Post, err error) {
	out = &Post{}

	switch s.t {
	case SqlStore:
		db := sqlStore.Store()
		out.t = s.t
		out.sqlModel = &sqlModels.Post{}
		if res := db.Where("uuid = ?", uuid).First(out.sqlModel); errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, res.Error
		}
	default:
		panic(ErrInvalidStore)
	}

	return
}

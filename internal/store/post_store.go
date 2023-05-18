package store

import (
	"errors"
	"log"
	"os"

	"github.com/romv7/blogs/internal/pb"
	"github.com/romv7/blogs/internal/storage"
	sqlStore "github.com/romv7/blogs/internal/store/sql"
	sqlModels "github.com/romv7/blogs/internal/store/sql/models"
	"github.com/romv7/blogs/internal/utils/author"

	"gorm.io/gorm"
)

type Post struct {
	t        StoreType
	s        storage.StorageDriverType
	sqlModel *sqlModels.Post
}

func (p *Post) Proto() *pb.Post {
	ustore := NewUserStore(SqlStore)
	cstore := NewCommentStore(SqlStore)

	var pout *pb.Post

	switch p.t {
	case SqlStore:
		pout = p.sqlModel.Proto()

		if u, err := ustore.GetById(p.sqlModel.UserId); errors.Is(err, gorm.ErrRecordNotFound) {
			// TODO: Handle post that has no owner.
		} else {
			pout.User = u.Proto()

			if pout.User.Type == pb.User_T_AUTHOR {
				ah := author.NewAuthorHelper(pout.User, p.s)

				// TODO: Get post metadata from the p.s storage.
				switch p.s {
				case storage.Plain:
					m, content, err := ah.GetAuthorPostMetadata(pout)

					var Err error

					// The post still isn't save in the storage.
					if err == nil {
						pout.HeadlineText = m.HeadlineText
						pout.SummaryText = m.SummaryText
						pout.Refs = m.References
						pout.Content = content
					} else if os.IsNotExist(err) {
						Err = ah.SaveAuthorPost(pout)
					} else {
						Err = err
					}

					if Err != nil {
						log.Panic(Err)
					}
				default:
					log.Panic(storage.ErrorInvalidStorageDriver)
				}
			}
		}

		pout.Comments = cstore.TargetCommentProtoTree(pout.Uuid)

		return pout
	default:
		log.Panic(ErrInvalidStore)
	}

	return nil
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
	if u.Type == pb.User_T_NORMAL {
		log.Panic(author.ErrInvalidArgument)
	}

	out = &Post{}
	p.User = u

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
	default:
		log.Panic(ErrInvalidStore)
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
		if res := db.Where("uuid = ?", uuid).Limit(1).Find(out.sqlModel); errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, res.Error
		}
	default:
		log.Panic(ErrInvalidStore)
	}

	return
}

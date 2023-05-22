package store

import (
	"errors"
	"log"

	"github.com/romv7/blogs/internal/pb"
	"github.com/romv7/blogs/internal/storage"
	sqlStore "github.com/romv7/blogs/internal/store/sql"
	sqlModels "github.com/romv7/blogs/internal/store/sql/models"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func setTheTargetFor(c *pb.Comment, t_uuid string, T pb.Comment_TargetType) (err error) {
	switch T {
	case pb.Comment_TT_COMMENT:
		commentStore := NewCommentStore(SqlStore)
		t, err := commentStore.GetByUuid(t_uuid)
		if err != nil {
			return err
		}

		c.Target = &pb.Comment_TComment{TComment: t.Proto()}
	case pb.Comment_TT_USER:
		userStore := NewUserStore(SqlStore)

		t, err := userStore.GetByUuid(t_uuid)
		if err != nil {
			return err
		}

		c.Target = &pb.Comment_TUser{TUser: t.Proto()}
	case pb.Comment_TT_POST:
		postStore := NewPostStore(SqlStore)

		t, err := postStore.GetByUuid(t_uuid)
		if err != nil {
			return err
		}

		c.Target = &pb.Comment_TPost{TPost: t.Proto()}
	default:
		return status.Errorf(codes.InvalidArgument, "provided an invalid target type (%d)", T)
	}

	return
}

type Comment struct {
	t        StoreType
	s        storage.StorageDriverType
	sqlModel *sqlModels.Comment
}

type CommentStore struct {
	t StoreType
	s storage.StorageDriverType
}

func NewCommentStore(t StoreType) *CommentStore {
	return &CommentStore{t, storage.Plain}
}

func (s *CommentStore) GetMainStore() (S any) {
	switch s.t {
	case SqlStore:
		S = sqlStore.Store()
	default:
		log.Panic(ErrInvalidStore)
	}

	return
}

func (s *CommentStore) NewComment(c *pb.Comment, t_uuid string, T pb.Comment_TargetType) (out *Comment) {
	out = &Comment{}

	if s.t == SqlStore {
		if err := setTheTargetFor(c, t_uuid, T); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			log.Panic(err)
		}

		cout, err := sqlModels.NewComment(c, t_uuid)
		if err != nil {
			log.Panic(err)
		}

		out.sqlModel = cout
	} else {
		log.Panic(ErrInvalidStore)
	}

	return
}

func (s *CommentStore) Save(c *Comment) (err error) {
	switch s.t {
	case SqlStore:
		db := sqlStore.Store()
		res := db.Save(c.sqlModel)
		err = res.Error
	default:
		log.Panic(ErrInvalidStore)
	}

	return
}

func (s *CommentStore) Delete(c *Comment) (err error) {

	switch s.t {
	case SqlStore:
		db := sqlStore.Store()

		if res := db.Where("uuid = ?", c.sqlModel.Uuid).Delete(c.sqlModel); res.Error != nil {
			return res.Error
		}

		c = nil
	default:
		log.Panic(ErrInvalidStore)
	}

	return
}

func (s *CommentStore) GetById(id uint64) (out *Comment, err error) {
	out = &Comment{}

	switch s.t {
	case SqlStore:
		db := sqlStore.Store()
		out.t = s.t
		out.sqlModel = &sqlModels.Comment{ID: id}
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

func (s *CommentStore) GetByUuid(uuid string) (out *Comment, err error) {
	out = &Comment{}

	switch s.t {
	case SqlStore:
		db := sqlStore.Store()
		out.t = s.t
		out.sqlModel = &sqlModels.Comment{}
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

func (s *CommentStore) TargetCommentProtoTree(t_uuid string) (out []*pb.Comment) {
	out = make([]*pb.Comment, 0)
	ustore := NewUserStore(SqlStore)

	switch s.t {
	case SqlStore:
		db := sqlStore.Store()
		C := make([]*sqlModels.Comment, 0)

		if res := db.Where("target_uuid = ?", t_uuid).Find(&C); errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil
		}

		for _, res := range C {
			c := res.Proto()

			if u, err := ustore.GetById(res.UserId); err != nil {
				// TODO: Handle comment that has no owner
			} else {
				c.User = u.Proto()
			}

			// Get all replies targeting this comment.
			c.Replies = s.TargetCommentProtoTree(c.Uuid)

			out = append(out, c)
		}

	default:
		log.Panic(ErrInvalidStore)
	}

	return
}

package store

import (
	"log"

	"github.com/romv7/blogs/internal/pb"
)

func (c *Comment) Proto() (out *pb.Comment) {
	cstore := NewCommentStore(SqlStore)
	ustore := NewUserStore(SqlStore)

	switch c.t {
	case SqlStore:
		out = c.sqlModel.Proto()
		out.Replies = cstore.TargetCommentProtoTree(out.Uuid)

		if u, err := ustore.GetById(c.sqlModel.UserId); err != nil {
			// TODO: Handle comment that has no owner
		} else {
			out.User = u.Proto()
		}

	default:
		log.Panic(ErrInvalidStore)
	}

	return
}

func (c *Comment) Save() (err error) {
	cstore := NewCommentStore(c.t)
	err = cstore.Save(c)

	return
}

func (c *Comment) Delete() (err error) {
	cstore := NewCommentStore(c.t)
	err = cstore.Delete(c)

	return
}

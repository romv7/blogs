package store

import (
	"log"

	"github.com/romv7/blogs/internal/pb"
)

// TODO: Add a documentation to this method.
// TODO: Should set the c.Target field of the comment using the retrieved target_uuid from the database.
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

// Just another helper method for the CommentStore.Save() method. Used this as much as possible instead of
// directly making a new instance of the CommentStore each time you want to save a comment to the data source.
func (c *Comment) Save() (err error) {
	cstore := NewCommentStore(c.t)
	err = cstore.Save(c)

	return
}

// Just another helper method for the CommentStore.Delete() method. Used this as much as possible instead of
// directly making a new instance of the CommentStore each time you want to delete a comment to the data source.
func (c *Comment) Delete() (err error) {
	cstore := NewCommentStore(c.t)
	err = cstore.Delete(c)

	return
}

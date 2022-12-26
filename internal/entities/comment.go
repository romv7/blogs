package entities

import (
	"time"

	"github.com/rommms07/blogs/pb"
)

type Comment struct {
	*User
	*pb.Comment
}

func NewComment(user *User, commentText string, targetId uint64, targetType pb.Comment_TargetType) (comment *Comment) {
	comment = &Comment{
		User: user,
		Comment: &pb.Comment{
			User:        user.User,
			CommentText: commentText,
			TargetType:  targetType,
			TargetId:    targetId,
			CreatedAt:   uint64(time.Now().Unix()),
		},
	}

	return
}

func (c *Comment) React(typ pb.React_Type) *Comment {
	return c
}

func (c *Comment) Reply(reply *Comment) (*Comment, error) {
	return c, nil
}

func (c *Comment) Remove() (*Comment, error) {
	return c, nil
}

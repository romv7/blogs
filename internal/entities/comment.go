package entities

import (
	"github.com/rommms07/blogs/pb"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Comment struct {
	*User
	*pb.Comment
}

func NewComment(user *User, commentText string, targetUuid string) (comment *Comment) {
	comment = &Comment{
		User: user,
		Comment: &pb.Comment{
			Uuid:		uuid.New().String(),
			UserUuid:	user.Uuid,
			CommentText:	commentText,
			TargetUuid:	targetUuid,
			State:		&pb.CommentState {
				Edited:    false,
				CreatedAt: timestamppb.Now(),
				EditedAt:  timestamppb.Now(),
			},
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

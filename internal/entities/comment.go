package entities

import (
	"github.com/google/uuid"
	"github.com/romv7/blogs/internal/utils/rand"
	"github.com/romv7/blogs/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Comment struct {
	*User
	*pb.Comment
}

func NewComment(user *User, commentText string, target any) (comment *Comment) {
	now := timestamppb.Now()
	N, _ := rand.Rand()

	comment = &Comment{
		User: user,
		Comment: &pb.Comment{
			Id:          uint32(now.AsTime().Unix()) + uint32(N),
			Uuid:        uuid.New().String(),
			User:        user.User,
			CommentText: commentText,
			State: &pb.CommentState{
				Reacts:    &pb.Reacts{},
				Edited:    false,
				CreatedAt: now,
				EditedAt:  now,
			},
		},
	}

	comment.SetTarget(target)

	return
}

func (c *Comment) UniqueKey() string {
	return c.CommentText + "#" + c.User.UniqueKey()
}

func (c *Comment) React() *Comment {
	return c
}

func (c *Comment) Reply(reply *Comment) (*Comment, error) {
	return c, nil
}

func (c *Comment) Remove() (*Comment, error) {
	return c, nil
}

func (c *Comment) GetTargetUuid() (res string) {
	switch v := c.Target.(type) {
	case *pb.Comment_TComment:
		res = v.TComment.Uuid
	case *pb.Comment_TPost:
		res = v.TPost.Uuid
	case *pb.Comment_TUser:
		res = v.TUser.Uuid
	}

	return
}

func (c *Comment) SetTarget(target any) {
	switch v := target.(type) {
	case *Comment:
		c.Target = &pb.Comment_TComment{TComment: v.Comment}
		c.TargetType = pb.Comment_TT_COMMENT
	case *Post:
		c.Target = &pb.Comment_TPost{TPost: v.Post}
		c.TargetType = pb.Comment_TT_POST
	case *User:
		c.Target = &pb.Comment_TUser{TUser: v.User}
		c.TargetType = pb.Comment_TT_USER
	}
}

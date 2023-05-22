package models

import (
	"time"

	"github.com/romv7/blogs/internal/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Comment struct {
	Id          float64               `json:"id"`
	Uuid        string                `json:"uuid"`
	User        *User                 `json:"user"`
	CommentText string                `json:"commentText"`
	Replies     []*Comment            `json:"replies"`
	TargetType  pb.Comment_TargetType `json:"targetType"`
	// Target      any                   `json:"target"`
	Reacts    *Reacts   `json:"reacts"`
	CreatedAt time.Time `json:"createdAt"`
	EditedAt  time.Time `json:"editedAt"`
}

func Proto_GQLModelComment(c *pb.Comment) (out *Comment) {
	out = &Comment{
		Id:          float64(c.Id),
		Uuid:        c.Uuid,
		User:        Proto_GQLModelUser(c.User),
		CommentText: c.CommentText.Data,
		TargetType:  c.TargetType,
		// Target:      c.Target,
		Reacts:    Proto_GQLModelReacts(c.State.Reacts),
		CreatedAt: c.State.CreatedAt.AsTime(),
		EditedAt:  c.State.EditedAt.AsTime(),
	}

	for _, reply := range c.Replies {
		out.Replies = append(out.Replies, Proto_GQLModelComment(reply))
	}

	return
}

func (c *Comment) Proto() (out *pb.Comment) {
	out = &pb.Comment{
		Id:          uint64(c.Id),
		Uuid:        c.Uuid,
		User:        c.User.Proto(),
		CommentText: &pb.CommentText{Data: c.CommentText},
		TargetType:  c.TargetType,
		State: &pb.CommentState{
			Reacts:    c.Reacts.Proto(),
			CreatedAt: timestamppb.New(c.CreatedAt),
			EditedAt:  timestamppb.New(c.EditedAt),
		},
	}

	for _, reply := range c.Replies {
		out.Replies = append(out.Replies, reply.Proto())
	}

	return
}

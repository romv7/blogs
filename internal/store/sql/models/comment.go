package models

import (
	"time"

	"github.com/romv7/blogs/internal/pb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Comment struct {
	ID          uint32                `gorm:"column:id;autoIncrement:false"`
	Uuid        string                `gorm:"column:uuid;unique"`
	UserId      uint32                `gorm:"column:user_id"`
	CommentText []byte                `gorm:"column:comment_text;type:blob"`
	Reacts      []byte                `gorm:"column:reacts;type:blob"`
	CreatedAt   time.Time             `gorm:"column:created_at"`
	EditedAt    time.Time             `gorm:"column:edited_at"`
	TargetType  pb.Comment_TargetType `gorm:"column:target_type"`
	TargetUuid  string                `gorm:"column:target_uuid"`
}

func NewComment(c *pb.Comment, t_uuid string) (*Comment, error) {
	out := &Comment{
		ID:         c.Id,
		Uuid:       c.Uuid,
		UserId:     c.User.Id,
		CreatedAt:  c.State.CreatedAt.AsTime(),
		EditedAt:   c.State.EditedAt.AsTime(),
		TargetUuid: t_uuid,
		TargetType: c.TargetType,
	}

	commentText_b, err := proto.Marshal(c.CommentText)
	if err != nil {
		return nil, err
	}

	reacts_b, err := proto.Marshal(c.State.Reacts)
	if err != nil {
		return nil, err
	}

	out.CommentText = commentText_b
	out.Reacts = reacts_b

	return out, nil
}

func (c *Comment) Proto() (out *pb.Comment) {
	out = &pb.Comment{
		Id:         c.ID,
		Uuid:       c.Uuid,
		TargetType: c.TargetType,
		Target:     nil,
		State: &pb.CommentState{
			Reacts:    nil,
			CreatedAt: timestamppb.New(c.CreatedAt),
			EditedAt:  timestamppb.New(c.EditedAt),
		},
	}

	reacts := new(pb.Reacts)
	commentText := new(pb.CommentText)

	if err := proto.Unmarshal(c.Reacts, reacts); err != nil {
		panic(err)
	}

	if err := proto.Unmarshal(c.CommentText, commentText); err != nil {
		panic(err)
	}

	out.State.Reacts = reacts
	out.CommentText = commentText

	return
}

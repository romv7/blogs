package models

import (
	"time"

	"github.com/romv7/blogs/internal/pb"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Post struct {
	ID          uint32              `gorm:"column:id;autoIncrement:false"`
	UserId      uint32              `gorm:"column:user_id"`
	Uuid        string              `gorm:"column:uuid;unique"`
	Tags        []byte              `gorm:"column:tags;type:blob"`
	Reacts      []byte              `gorm:"column:reacts;type:blob"`
	Stage       pb.PostState_Stage  `gorm:"column:stage"`
	Status      pb.PostState_Status `gorm:"column:status"`
	RevisedAt   time.Time           `gorm:"column:revised_at"`
	ArchivedAt  time.Time           `gorm:"column:archived_at"`
	PublishedAt time.Time           `gorm:"column:published_at"`
	CreatedAt   time.Time           `gorm:"column:created_at"`
	PrevId      uint32              `gorm:"column:prev_id"`
}

func NewPost(p *pb.Post) (out *Post) {
	tags_b, err := proto.Marshal(p.Tags)
	if err != nil {
		panic(err)
	}

	reacts_b, err := proto.Marshal(p.State.Reacts)
	if err != nil {
		panic(err)
	}

	out = &Post{
		ID:          p.Id,
		UserId:      p.User.Id,
		Uuid:        p.Uuid,
		Tags:        tags_b,
		Reacts:      reacts_b,
		Stage:       p.State.Stage,
		Status:      p.State.Status,
		RevisedAt:   p.State.RevisedAt.AsTime(),
		ArchivedAt:  p.State.ArchivedAt.AsTime(),
		PublishedAt: p.State.PublishedAt.AsTime(),
		CreatedAt:   p.State.CreatedAt.AsTime(),
	}

	if p.Prev != nil {
		out.PrevId = p.Prev.Id
	}

	return
}

func (p *Post) Proto() (out *pb.Post) {
	out = &pb.Post{
		Id:   p.ID,
		Uuid: p.Uuid,
		State: &pb.PostState{
			Stage:  p.Stage,
			Status: p.Status,

			CreatedAt:   timestamppb.New(p.CreatedAt),
			PublishedAt: timestamppb.New(p.PublishedAt),
			RevisedAt:   timestamppb.New(p.RevisedAt),
			ArchivedAt:  timestamppb.New(p.ArchivedAt),
		},
	}

	tags := new(pb.Tags)
	reacts := new(pb.Reacts)

	if err := proto.Unmarshal(p.Tags, tags); err != nil {
		panic(err)
	}

	if err := proto.Unmarshal(p.Reacts, reacts); err != nil {
		panic(err)
	}

	out.Tags = tags
	out.State.Reacts = reacts

	return
}

package models

import (
	"time"

	"github.com/romv7/blogs/internal/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Post struct {
	Id           float64    `json:"id"`
	Uuid         string     `json:"uuid"`
	HeadlineText string     `json:"headlineText"`
	SummaryText  string     `json:"summaryText"`
	User         *User      `json:"user"`
	Comments     []*Comment `json:"comments"`
	Tags         []string   `json:"tags"`
	Attachments  []string   `json:"attachments"`
	Refs         []string   `json:"refs"`

	// TODO: Instead of including the entire content to the
	//       response of the query, use a temporary generated
	//       URL pointing the content itself.
	Content     string              `json:"content"`
	Reacts      *Reacts             `json:"reacts"`
	Stage       pb.PostState_Stage  `json:"stage"`
	Status      pb.PostState_Status `json:"status"`
	CreatedAt   time.Time           `json:"createdAt"`
	RevisedAt   time.Time           `json:"revisedAt"`
	ArchivedAt  time.Time           `json:"archivedAt"`
	PublishedAt time.Time           `json:"publishedAt"`

	// TODO: Add a way to historically navigate to old post versions. (???)
	Prev *Post `json:"-"`
}

func Proto_GQLModelPost(p *pb.Post) (out *Post) {
	out = &Post{
		Id:           float64(p.Id),
		Uuid:         p.Uuid,
		HeadlineText: p.HeadlineText,
		SummaryText:  p.SummaryText,
		User:         Proto_GQLModelUser(p.User),
		Tags:         p.Tags.Data,
		Attachments:  p.Attachments,
		Refs:         p.Refs,
		Content:      p.Content,
		Stage:        p.State.Stage,
		Status:       p.State.Status,
		CreatedAt:    p.State.CreatedAt.AsTime(),
		RevisedAt:    p.State.RevisedAt.AsTime(),
		ArchivedAt:   p.State.ArchivedAt.AsTime(),
		PublishedAt:  p.State.PublishedAt.AsTime(),
	}

	for _, comment := range p.Comments {
		out.Comments = append(out.Comments, Proto_GQLModelComment(comment))
	}

	return
}

func (p *Post) Proto() (out *pb.Post) {
	out = &pb.Post{
		Id:           uint64(p.Id),
		Uuid:         p.Uuid,
		HeadlineText: p.HeadlineText,
		SummaryText:  p.SummaryText,
		User:         p.User.Proto(),
		Tags:         &pb.Tags{Data: p.Tags},
		Attachments:  p.Attachments,
		Refs:         p.Refs,
		Content:      p.Content,
		State: &pb.PostState{
			Reacts:      p.Reacts.Proto(),
			Stage:       p.Stage,
			Status:      p.Status,
			RevisedAt:   timestamppb.New(p.RevisedAt),
			ArchivedAt:  timestamppb.New(p.ArchivedAt),
			PublishedAt: timestamppb.New(p.PublishedAt),
			CreatedAt:   timestamppb.New(p.CreatedAt),
		},
	}

	for _, comment := range p.Comments {
		out.Comments = append(out.Comments, comment.Proto())
	}

	return
}

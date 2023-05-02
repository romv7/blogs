package entities

import (
	"github.com/google/uuid"
	"github.com/romv7/blogs/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Post struct {
	User *User
	*pb.Post
}

func NewPost(user *User, headlineText, summaryText string, tags ...string) (post *Post) {
	createdAt := timestamppb.Now()

	post = &Post{
		User: user,
		Post: &pb.Post{
			Id:           uint32(createdAt.AsTime().Unix()),
			Uuid:         uuid.New().String(),
			User:         user.User,
			HeadlineText: headlineText,
			SummaryText:  summaryText,
			Tags:         tags,
			Comments:     make([]*pb.Comment, 0),
			Original:     nil,
			State: &pb.PostState{
				Stage:       pb.PostState_S_WIP,
				Status:      pb.PostState_S_DRAFT,
				RevisedAt:   createdAt,
				ArchivedAt:  nil,
				PublishedAt: nil,
				CreatedAt:   createdAt,
				Reacts:      &pb.Reacts{},
			},
		},
	}

	return
}

func (p *Post) UniqueKey() string {
	return p.HeadlineText + "#" + p.User.UniqueKey()
}

func (p *Post) SetHeadlineText(htxt string) *Post {
	p.HeadlineText = htxt
	return p
}

func (p *Post) SetSummaryText(summary string) *Post {
	p.SummaryText = summary
	return p
}

func (p *Post) ToArchive() *Post {
	return p
}

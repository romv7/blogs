package models

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/romv7/blogs/internal/pb"
)

type GQLModel_PostResolver struct {
	*Post
}

func (p *GQLModel_PostResolver) Id() float64 {
	return p.Post.Id
}

func (p *GQLModel_PostResolver) Uuid() string {
	return p.Post.Uuid
}

func (p *GQLModel_PostResolver) HeadlineText() string {
	return p.Post.HeadlineText
}

func (p *GQLModel_PostResolver) SummaryText() string {
	return p.Post.SummaryText
}

func (p *GQLModel_PostResolver) User() *GQLModel_UserResolver {
	return NewGQLModel_UserResolver(p.Post.User)
}

func (p *GQLModel_PostResolver) Comments() []*GQLModel_CommentResolver {
	out := make([]*GQLModel_CommentResolver, 0)

	for _, comment := range p.Post.Comments {
		out = append(out, NewGQLModel_CommentResolver(comment))
	}

	return out
}

func (p *GQLModel_PostResolver) Tags() []string {
	return p.Post.Tags
}

func (p *GQLModel_PostResolver) Attachments() []string {
	return p.Post.Attachments
}

func (p *GQLModel_PostResolver) Refs() []string {
	return p.Post.Refs
}

func (p *GQLModel_PostResolver) Content() string {
	return p.Post.Content
}

func (p *GQLModel_PostResolver) Reacts() *GQLModel_ReactsResolver {
	return NewGQLModel_ReactsResolver(p.Post.Reacts)
}

func (p *GQLModel_PostResolver) Stage() pb.PostState_Stage {
	return p.Post.Stage
}

func (p *GQLModel_PostResolver) Status() pb.PostState_Status {
	return p.Post.Status
}

func (p *GQLModel_PostResolver) CreatedAt() graphql.Time {
	return graphql.Time{Time: p.Post.CreatedAt}
}

func (p *GQLModel_PostResolver) RevisedAt() graphql.Time {
	return graphql.Time{Time: p.Post.RevisedAt}
}

func (p *GQLModel_PostResolver) ArchivedAt() (out *graphql.Time) {
	out = &graphql.Time{Time: p.Post.ArchivedAt}

	return
}

func (p *GQLModel_PostResolver) PublishedAt() (out *graphql.Time) {
	out = &graphql.Time{Time: p.Post.PublishedAt}

	return
}

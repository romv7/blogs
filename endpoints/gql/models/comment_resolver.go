package models

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/romv7/blogs/internal/pb"
)

type GQLModel_CommentResolver struct {
	*Comment
}

func NewGQLModel_CommentResolver(c *Comment) *GQLModel_CommentResolver {
	return &GQLModel_CommentResolver{c}
}

func (c *GQLModel_CommentResolver) Id() float64 {
	return c.Comment.Id
}

func (c *GQLModel_CommentResolver) Uuid() string {
	return c.Comment.Uuid
}

func (c *GQLModel_CommentResolver) User() *GQLModel_UserResolver {
	return NewGQLModel_UserResolver(c.Comment.User)
}

func (c *GQLModel_CommentResolver) CommentText() string {
	return c.Comment.CommentText
}

func (c *GQLModel_CommentResolver) Replies() []*GQLModel_CommentResolver {
	out := make([]*GQLModel_CommentResolver, 0)

	for _, reply := range c.Comment.Replies {
		out = append(out, NewGQLModel_CommentResolver(reply))
	}

	return out
}

// PROBLEM: Comment resolver cannot resolve a target!

func (c *GQLModel_CommentResolver) TargetType() pb.Comment_TargetType {
	return c.Comment.TargetType
}

func (c *GQLModel_CommentResolver) Reacts() *GQLModel_ReactsResolver {
	return NewGQLModel_ReactsResolver(c.Comment.Reacts)
}

func (c *GQLModel_CommentResolver) CreatedAt() graphql.Time {
	return graphql.Time{Time: c.Comment.CreatedAt}
}

func (c *GQLModel_CommentResolver) EditedAt() graphql.Time {
	return graphql.Time{Time: c.Comment.EditedAt}
}

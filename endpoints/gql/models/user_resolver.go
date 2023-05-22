package models

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/romv7/blogs/internal/pb"
)

type GQLModel_UserResolver struct {
	*User
}

func NewGQLModel_UserResolver(u *User) *GQLModel_UserResolver {
	return &GQLModel_UserResolver{u}
}

func (r *GQLModel_UserResolver) Id() float64 {
	return r.User.Id
}

func (r *GQLModel_UserResolver) Uuid() string {
	return r.User.Uuid
}

func (r *GQLModel_UserResolver) Name() string {
	return r.User.Name
}

func (r *GQLModel_UserResolver) Email() string {
	return r.User.Email
}

func (r *GQLModel_UserResolver) FullName() string {
	return r.User.FullName
}

func (r *GQLModel_UserResolver) Type() pb.User_Type {
	return r.User.Type
}

func (r *GQLModel_UserResolver) CreatedAt() graphql.Time {
	return graphql.Time{Time: r.User.CreatedAt}
}

func (r *GQLModel_UserResolver) UpdatedAt() graphql.Time {
	return graphql.Time{Time: r.User.UpdatedAt}
}

func (r *GQLModel_UserResolver) IsDisabled() bool {
	return r.User.IsDisabled
}

func (r *GQLModel_UserResolver) IsVerified() bool {
	return r.User.IsVerified
}

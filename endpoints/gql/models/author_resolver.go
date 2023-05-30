package models

import "context"

type GQLModel_AuthorResolver struct {
	authorInfo *GQLModel_UserResolver
}

func NewGQLModel_AuthorResolver(ctx context.Context, args *ArgsAuthor) *GQLModel_AuthorResolver {
	return &GQLModel_AuthorResolver{}
}

func (r *GQLModel_AuthorResolver) AuthorInfo() *GQLModel_UserResolver {
	return r.authorInfo
}

func (r *GQLModel_AuthorResolver) LatestPosts(ctx context.Context, args *ArgsAuthorLatestPosts) *GQLModel_PostConnectionResolver {
	return NewGQLModel_PostConnectionResolver(&PostConnection{})
}

func (r *GQLModel_AuthorResolver) Posts(ctx context.Context, args *ArgsAuthorPosts) *GQLModel_PostConnectionResolver {
	return NewGQLModel_PostConnectionResolver(&PostConnection{})
}

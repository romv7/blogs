package models

import "context"

type GQLModel_GlobalLatestPosts_Resolver struct {
	postConn *GQLModel_PostConnectionResolver
}

func NewGQLModel_GlobalLatestPosts_Resolver(ctx context.Context, args *ArgsGlobalLatestPosts) *GQLModel_GlobalLatestPosts_Resolver {
	return &GQLModel_GlobalLatestPosts_Resolver{
		postConn: NewGQLModel_PostConnectionResolver(&PostConnection{}),
	}
}

func (r *GQLModel_GlobalLatestPosts_Resolver) Edges() *[]*GQLModel_PostEdgeResolver {
	return r.postConn.Edges()
}

func (r *GQLModel_GlobalLatestPosts_Resolver) PageInfo() *GQLModel_PageInfoResolver {
	return r.PageInfo()
}

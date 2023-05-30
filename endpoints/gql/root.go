package gql

import (
	"context"

	"github.com/romv7/blogs/endpoints/gql/models"
)

type RootQuery struct{}

func NewRootQuery() *RootQuery {
	return &RootQuery{}
}

func (RootQuery) GlobalSearch(ctx context.Context, args *models.ArgsGlobalSearch) (*models.GQLModel_GlobalSearchResultsResolver, error) {
	if args.Keyword == nil {
		return &models.GQLModel_GlobalSearchResultsResolver{}, nil
	}

	return models.NewGQLModel_GlobalSearchResultsResolver(ctx, *args.Keyword), nil
}

func (RootQuery) GlobalLatestPosts(ctx context.Context, args *models.ArgsGlobalLatestPosts) (*models.GQLModel_GlobalLatestPosts_Resolver, error) {
	return models.NewGQLModel_GlobalLatestPosts_Resolver(ctx, args), nil
}

func (RootQuery) Author(ctx context.Context, args *models.ArgsAuthor) *models.GQLModel_AuthorResolver {
	return models.NewGQLModel_AuthorResolver(ctx, args)
}

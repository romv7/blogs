package gql

import (
	"context"

	"github.com/romv7/blogs/endpoints/gql/models"
)

type RootQuery struct{}

func NewRootQuery() *RootQuery {
	return &RootQuery{}
}

func (RootQuery) GlobalSearch(ctx context.Context, args models.ArgsGlobalSearch) (*models.GQLModel_GlobalSearchResultsResolver, error) {
	if args.Keyword == nil {
		return &models.GQLModel_GlobalSearchResultsResolver{}, nil
	}

	return models.NewGlobalSearch(ctx, *args.Keyword), nil
}

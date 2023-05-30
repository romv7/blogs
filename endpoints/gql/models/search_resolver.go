package models

import (
	"context"
	"time"

	"github.com/graph-gophers/graphql-go"
)

type GQLModel_GlobalSearchResultsResolver struct {
	res *GlobalSearchResults
}

func NewGQLModel_GlobalSearchResultsResolver(ctx context.Context, keyword string) *GQLModel_GlobalSearchResultsResolver {

	return &GQLModel_GlobalSearchResultsResolver{
		res: &GlobalSearchResults{
			Keyword:  keyword,
			People:   &GQLModel_UserConnectionResolver{},
			Comments: &GQLModel_CommentConnectionResolver{},
			Posts:    &GQLModel_PostConnectionResolver{},
			Stats: &GQLModel_SearchResultsStatsResolver{
				res: &GlobalSearchResultsStats{},
			},
			StartTime: time.Now(),
			EndTime:   time.Now(),
		},
	}
}

// Type aliases for conviniently accessing some structs.
type ctxt = context.Context
type userFilter = *ArgsSearchResultUserFilter
type postFilter = *ArgsSearchResultPostFilter
type commentFilter = *ArgsSearchResultCommentFilter

func (s *GQLModel_GlobalSearchResultsResolver) Keyword() string {
	return s.res.Keyword
}

func (s *GQLModel_GlobalSearchResultsResolver) People(ctx ctxt, args userFilter) *GQLModel_UserConnectionResolver {
	return s.res.People
}

func (s *GQLModel_GlobalSearchResultsResolver) Posts(ctx ctxt, args postFilter) *GQLModel_PostConnectionResolver {
	return s.res.Posts
}

func (s *GQLModel_GlobalSearchResultsResolver) Comments(ctx ctxt, args commentFilter) *GQLModel_CommentConnectionResolver {
	return s.res.Comments
}

func (s *GQLModel_GlobalSearchResultsResolver) Stats() *GQLModel_SearchResultsStatsResolver {
	return s.res.Stats
}

func (s *GQLModel_GlobalSearchResultsResolver) StartTime() graphql.Time {
	return graphql.Time{Time: s.res.StartTime}
}

func (s *GQLModel_GlobalSearchResultsResolver) EndTime() graphql.Time {
	return graphql.Time{Time: s.res.EndTime}
}

type GQLModel_SearchResultsStatsResolver struct {
	res *GlobalSearchResultsStats
}

func NewGQLModel_SearchResultsStatsResolver(res *GlobalSearchResultsStats) *GQLModel_SearchResultsStatsResolver {
	return &GQLModel_SearchResultsStatsResolver{res}
}

func (s *GQLModel_SearchResultsStatsResolver) PeopleCount() float64 {
	return float64(s.res.PeopleCount)
}

func (s *GQLModel_SearchResultsStatsResolver) PostsCount() float64 {
	return float64(s.res.PostsCount)
}

func (s *GQLModel_SearchResultsStatsResolver) CommentsCount() float64 {
	return float64(s.res.CommentsCount)
}

package models

import (
	"context"
	"time"

	"github.com/graph-gophers/graphql-go"
)

type GQLModel_GlobalSearchResultsResolver struct {
	res *GlobalSearchResults
}

func NewGlobalSearch(ctx context.Context, keyword string) *GQLModel_GlobalSearchResultsResolver {

	return &GQLModel_GlobalSearchResultsResolver{
		res: &GlobalSearchResults{
			Keyword:  keyword,
			People:   &[]*GQLModel_UserResolver{},
			Comments: &[]*GQLModel_CommentResolver{},
			Posts:    &[]*GQLModel_PostResolver{},
			Stats: &GQLModel_SearchResultsStatsResolver{
				res: &GlobalSearchResultsStats{},
			},
			StartTime: time.Now(),
			EndTime:   time.Now(),
		},
	}
}

func (s *GQLModel_GlobalSearchResultsResolver) Keyword() string {
	return s.res.Keyword
}

func (s *GQLModel_GlobalSearchResultsResolver) People() []*GQLModel_UserResolver {
	return *s.res.People
}

func (s *GQLModel_GlobalSearchResultsResolver) Posts() []*GQLModel_PostResolver {
	return *s.res.Posts
}

func (s *GQLModel_GlobalSearchResultsResolver) Comments() []*GQLModel_CommentResolver {
	return *s.res.Comments
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

package models

import "time"

type GlobalSearchResults struct {
	Keyword   string                               `json:"keyword"`
	People    *[]*GQLModel_UserResolver            `json:"people"`
	Posts     *[]*GQLModel_PostResolver            `json:"posts"`
	Comments  *[]*GQLModel_CommentResolver         `json:"comments"`
	Stats     *GQLModel_SearchResultsStatsResolver `json:"stats"`
	StartTime time.Time                            `json:"startTime"`
	EndTime   time.Time                            `json:"endTime"`
}

type GlobalSearchResultsStats struct {
	PeopleCount   uint64 `json:"peopleCount"`
	PostsCount    uint64 `json:"postsCount"`
	CommentsCount uint64 `json:"commentsCount"`
}

type ArgsGlobalSearch struct {
	Keyword *string
}

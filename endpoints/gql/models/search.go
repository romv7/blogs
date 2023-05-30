package models

import "time"

type GlobalSearchResults struct {
	Keyword   string                               `json:"keyword"`
	People    *GQLModel_UserConnectionResolver     `json:"people"`
	Posts     *GQLModel_PostConnectionResolver     `json:"posts"`
	Comments  *GQLModel_CommentConnectionResolver  `json:"comments"`
	Stats     *GQLModel_SearchResultsStatsResolver `json:"stats"`
	StartTime time.Time                            `json:"startTime"`
	EndTime   time.Time                            `json:"endTime"`
}

type GlobalSearchResultsStats struct {
	PeopleCount   uint64 `json:"peopleCount"`
	PostsCount    uint64 `json:"postsCount"`
	CommentsCount uint64 `json:"commentsCount"`
}

// SearchResultFilter
type SearchResultFilter struct {
}

type UserFilter struct {
	SearchResultFilter
}

type PostFilter struct {
	SearchResultFilter
}

type CommentFilter struct {
	SearchResultFilter
}

type ArgsGlobalSearch struct {
	Keyword  *string
	Category *[]string
}

type ArgsSearchResultUserFilter struct {
	Input *UserFilter
}

type ArgsSearchResultPostFilter struct {
	Input *PostFilter
}

type ArgsSearchResultCommentFilter struct {
	Input *CommentFilter
}

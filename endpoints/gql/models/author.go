package models

type ArgsAuthor struct {
	Id   *float64
	Uuid *string
}

type ArgsAuthorLatestPosts struct {
	Filter *PostConnectionFilter
}

type ArgsAuthorPosts struct {
	First, Last *int64
	After       *PostForwardCursor
	Before      *PostBackwardCursor
	Filter      *PostConnectionFilter
}

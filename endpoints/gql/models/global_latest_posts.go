package models

type ArgsGlobalLatestPosts struct {
	First, Last *int32
	After       *PostForwardCursor
	Before      *PostBackwardCursor
}

package models

type PageInfo struct {
	HasNextPage     bool
	HasPreviousPage bool
	StartCursor     *string
	EndCursor       *string
}

type PostEdge struct {
	Cursor string
	Node   *GQLModel_PostResolver
}

type PostConnection struct {
	Edges    *[]*GQLModel_PostEdgeResolver
	PageInfo *GQLModel_PageInfoResolver
}

type UserEdge struct {
	Cursor string
	Node   *GQLModel_UserResolver
}

type UserConnection struct {
	Edges    *[]*GQLModel_UserEdgeResolver
	PageInfo *GQLModel_PageInfoResolver
}

type CommentEdge struct {
	Cursor string
	Node   *GQLModel_CommentResolver
}

type CommentConnection struct {
	Edges    *[]*GQLModel_CommentEdgeResolver
	PageInfo *GQLModel_PageInfoResolver
}

type PostForwardCursor struct{}

type PostBackwardCursor struct{}

type UserForwardCursor struct{}

type UserBackwardCursor struct{}

type CommentForwardCursor struct{}

type CommentBackwardCursor struct{}

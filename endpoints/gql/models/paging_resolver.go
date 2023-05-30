package models

type GQLModel_PageInfoResolver struct {
	pageInfo *PageInfo
}

func NewGQLModel_PageInfoResolver(pinfo *PageInfo) *GQLModel_PageInfoResolver {
	return &GQLModel_PageInfoResolver{pinfo}
}

func (r *GQLModel_PageInfoResolver) HasNextPage() bool {
	return r.pageInfo.HasNextPage
}

func (r *GQLModel_PageInfoResolver) HasPreviousPage() bool {
	return r.pageInfo.HasPreviousPage
}

func (r *GQLModel_PageInfoResolver) StartCursor() *string {
	return r.pageInfo.StartCursor
}

func (r *GQLModel_PageInfoResolver) EndCursor() *string {
	return r.pageInfo.EndCursor
}

type GQLModel_PostEdgeResolver struct {
	postEdge *PostEdge
}

func NewGQLModel_PostEdgeResolver(edge *PostEdge) *GQLModel_PostEdgeResolver {
	return &GQLModel_PostEdgeResolver{edge}
}

func (r *GQLModel_PostEdgeResolver) Cursor() string {
	return r.postEdge.Cursor
}

func (r *GQLModel_PostEdgeResolver) Node() *GQLModel_PostResolver {
	return r.postEdge.Node
}

type GQLModel_PostConnectionResolver struct {
	postConn *PostConnection
}

func NewGQLModel_PostConnectionResolver(conn *PostConnection) *GQLModel_PostConnectionResolver {
	return &GQLModel_PostConnectionResolver{conn}
}

func (r *GQLModel_PostConnectionResolver) Edges() *[]*GQLModel_PostEdgeResolver {
	return r.postConn.Edges
}

func (r *GQLModel_PostConnectionResolver) PageInfo() *GQLModel_PageInfoResolver {
	return r.postConn.PageInfo
}

type GQLModel_UserEdgeResolver struct {
	userEdge *UserEdge
}

func (r *GQLModel_UserEdgeResolver) Cursor() string {
	return r.userEdge.Cursor
}

func (r *GQLModel_UserEdgeResolver) Node() *GQLModel_UserResolver {
	return r.userEdge.Node
}

type GQLModel_UserConnectionResolver struct {
	userConn *UserConnection
}

func (r *GQLModel_UserConnectionResolver) Edges() *[]*GQLModel_UserEdgeResolver {
	return r.userConn.Edges
}

func (r *GQLModel_UserConnectionResolver) PageInfo() *GQLModel_PageInfoResolver {
	return r.userConn.PageInfo
}

type GQLModel_CommentEdgeResolver struct {
	commentEdge *CommentEdge
}

func (r *GQLModel_CommentEdgeResolver) Cursor() string {
	return r.commentEdge.Cursor
}

func (r *GQLModel_CommentEdgeResolver) Node() *GQLModel_CommentResolver {
	return r.commentEdge.Node
}

type GQLModel_CommentConnectionResolver struct {
	commentConn *CommentConnection
}

func (r *GQLModel_CommentConnectionResolver) Edges() *[]*GQLModel_CommentEdgeResolver {
	return r.commentConn.Edges
}

func (r *GQLModel_CommentConnectionResolver) PageInfo() *GQLModel_PageInfoResolver {
	return r.commentConn.PageInfo
}

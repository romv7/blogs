package grpc

import (
	"context"

	"github.com/google/uuid"
	"github.com/romv7/blogs/internal/pb"
	"github.com/romv7/blogs/internal/store"
	"github.com/romv7/blogs/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type BlogService struct {
	pb.UnimplementedBlogServiceServer
}

// TODO: NewBlogPost
func (svc *BlogService) NewBlogPost(ctx context.Context, params *pb.BlogService_NewBlogPost_Params) (*pb.Post, error) {
	if params.User == nil {
		return nil, status.Errorf(codes.InvalidArgument, ErrNotEnoughArgument.Error())
	}

	out := &pb.Post{Id: utils.RandomUniqueId(), Uuid: uuid.NewString()}
	out.User = params.User
	out.HeadlineText = params.HeadlineText
	out.SummaryText = params.SummaryText
	out.Tags = &pb.Tags{Data: params.Tags}
	out.Attachments = params.Attachments
	out.Refs = params.Refs
	out.Content = params.Content
	out.Prev = nil
	out.State = &pb.PostState{
		Stage:       pb.PostState_S_WIP,
		Status:      pb.PostState_S_DRAFT,
		RevisedAt:   timestamppb.Now(),
		CreatedAt:   timestamppb.Now(),
		ArchivedAt:  nil,
		PublishedAt: nil,
		Reacts:      &pb.Reacts{},
	}

	return out, nil
}

// TODO: SaveBlogPost
func (svc *BlogService) SaveBlogPost(ctx context.Context, params *pb.BlogService_SaveBlogPost_Params) (*pb.BlogService_SaveBlogPost_Response, error) {
	if params.Post == nil || params.User == nil {
		return &pb.BlogService_SaveBlogPost_Response{}, status.Errorf(codes.InvalidArgument, ErrNotEnoughArgument.Error())
	}

	pstore := store.NewPostStore(store.SqlStore)

	if err := pstore.NewPost(params.User, params.Post).Save(); err != nil {
		return &pb.BlogService_SaveBlogPost_Response{}, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.BlogService_SaveBlogPost_Response{}, nil
}

// TODO: DeleteBlogPost
func (svc *BlogService) DeleteBlogPost(ctx context.Context, params *pb.BlogService_DeleteBlogPost_Params) (*pb.BlogService_DeleteBlogPost_Response, error) {
	if params.Post == nil {
		return &pb.BlogService_DeleteBlogPost_Response{}, status.Errorf(codes.InvalidArgument, ErrNotEnoughArgument.Error())
	}

	pstore := store.NewPostStore(store.SqlStore)

	if err := pstore.NewPost(params.Post.User, params.Post).Delete(); err != nil {
		return &pb.BlogService_DeleteBlogPost_Response{}, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.BlogService_DeleteBlogPost_Response{}, nil
}

// TODO: @NewComment Recognize a comment text as spam or automated.

func (svc *BlogService) NewComment(ctx context.Context, params *pb.BlogService_NewComment_Params) (*pb.Comment, error) {
	if params.User == nil || len(params.CommentText) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, ErrNotEnoughArgument.Error())
	} else {

		// Checks the UUID passed as an argument.
		if _, err := uuid.Parse(params.TargetUuid); err != nil {
			return nil, status.Errorf(codes.Internal, err.Error())
		}
	}

	cstore := store.NewCommentStore(store.SqlStore)

	out := cstore.NewComment(
		&pb.Comment{
			Id:          utils.RandomUniqueId(),
			Uuid:        uuid.NewString(),
			User:        params.User,
			CommentText: &pb.CommentText{Data: params.CommentText},
			State: &pb.CommentState{
				CreatedAt: timestamppb.Now(),
				EditedAt:  timestamppb.Now(),
				Reacts:    &pb.Reacts{},
			},
		},
		params.TargetUuid,
		params.TargetType,
	).Proto()

	return out, nil
}

// TODO: SaveComment
func (svc *BlogService) SaveComment(ctx context.Context, params *pb.BlogService_SaveComment_Params) (*pb.BlogService_SaveComment_Response, error) {
	if params.Comment == nil {
		return &pb.BlogService_SaveComment_Response{}, status.Errorf(codes.InvalidArgument, ErrNotEnoughArgument.Error())
	}

	cstore := store.NewCommentStore(store.SqlStore)

	if err := cstore.NewComment(params.Comment, params.Comment.Uuid, params.Comment.TargetType).Save(); err != nil {
		return &pb.BlogService_SaveComment_Response{}, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.BlogService_SaveComment_Response{}, nil
}

// TODO: DeleteComment
func (svc *BlogService) DeleteComment(ctx context.Context, params *pb.BlogService_DeleteComment_Params) (*pb.BlogService_DeleteComment_Response, error) {
	if params.Comment == nil {
		return &pb.BlogService_DeleteComment_Response{}, status.Errorf(codes.InvalidArgument, ErrNotEnoughArgument.Error())
	}

	cstore := store.NewCommentStore(store.SqlStore)

	if err := cstore.NewComment(params.Comment, params.Comment.Uuid, params.Comment.TargetType).Delete(); err != nil {
		return &pb.BlogService_DeleteComment_Response{}, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.BlogService_DeleteComment_Response{}, nil
}

func (svc *BlogService) GlobalSearch(ctx context.Context, params *pb.BlogService_GlobalSearch_Params) (*pb.SearchResults, error) {

	return nil, nil
}

func (svc *BlogService) GlobalLatestBlogPosts(ctx context.Context, params *pb.BlogService_GlobalLatestBlogPosts_Params) (*pb.BlogService_GlobalLatestBlogPosts_Response, error) {

	return nil, nil
}

func (svc *BlogService) Author_GetAuthorInfo(ctx context.Context, params *pb.BlogService_AuthorGetInfo_Params) (*pb.User, error) {

	return nil, nil
}

func (svc *BlogService) Author_LatestBlogPosts(ctx context.Context, params *pb.BlogService_AuthorLatestBlogPosts_Params) (*pb.BlogService_AuthorLatestBlogPosts_Response, error) {

	return nil, nil
}

func (svc *BlogService) Author_GetBlogPosts(ctx context.Context, params *pb.BlogService_AuthorGetBlogPosts_Params) (*pb.BlogService_AuthorGetBlogPosts_Response, error) {

	return nil, nil
}

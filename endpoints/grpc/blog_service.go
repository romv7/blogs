package grpc

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/romv7/blogs/internal/pb"
	"github.com/romv7/blogs/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type B struct {
	pb.UnimplementedBlogServiceServer
}

var (
	ErrNotEnoughArgument = errors.New("Provided an insufficient number of required arguments.")
)

// TODO: NewBlogPost
func (svc *B) NewBlogPost(ctx context.Context, params *pb.BlogService_NewBlogPost_Params) (*pb.Post, error) {
	out := &pb.Post{Id: utils.RandomUniqueId(), Uuid: uuid.NewString()}

	if len(params.HeadlineText) == 0 || params.User == nil {
		return nil, status.Errorf(codes.InvalidArgument, ErrNotEnoughArgument.Error())
	}

	out.User = params.User
	out.HeadlineText = params.HeadlineText
	out.SummaryText = params.SummaryText
	out.Tags = &pb.Tags{Data: params.Tags}
	out.Images = params.Images
	out.Attachments = params.Attachments
	out.Refs = params.Refs
	out.Content = params.Content
	out.Prev = nil
	out.State = &pb.PostState{
		Stage:       pb.PostState_S_WIP,
		Status:      pb.PostState_S_DRAFT,
		RevisedAt:   nil,
		CreatedAt:   timestamppb.Now(),
		ArchivedAt:  nil,
		PublishedAt: nil,
		Reacts:      &pb.Reacts{},
	}

	return out, nil
}

// TODO: SaveBlogPost
func (svc *B) SaveBlogPost(ctx context.Context, params *pb.BlogService_SaveBlogPost_Params) (*pb.BlogService_SaveBlogPost_Response, error) {

	return &pb.BlogService_SaveBlogPost_Response{}, nil
}

// TODO: DeleteBlogPost
func (svc *B) DeleteBlogPost(ctx context.Context, params *pb.BlogService_DeleteBlogPost_Params) (*pb.BlogService_DeleteBlogPost_Response, error) {

	return &pb.BlogService_DeleteBlogPost_Response{}, nil
}

// TODO: NewComment
func (svc *B) NewComment(ctx context.Context, params *pb.BlogService_NewComment_Params) (*pb.Comment, error) {
	out := &pb.Comment{Id: utils.RandomUniqueId(), Uuid: uuid.NewString()}

	return out, nil
}

// TODO: SaveComment
func (svc *B) SaveComment(ctx context.Context, params *pb.BlogService_SaveComment_Params) (*pb.BlogService_SaveComment_Response, error) {

	return &pb.BlogService_SaveComment_Response{}, nil
}

// TODO: DeleteComment
func (svc *B) DeleteComment(ctx context.Context, params *pb.BlogService_DeleteComment_Params) (*pb.BlogService_DeleteComment_Response, error) {

	return &pb.BlogService_DeleteComment_Response{}, nil
}

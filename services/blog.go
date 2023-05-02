package services

import (
	"context"

	"github.com/romv7/blogs/pb"
)

type blogSvcServer struct {
	pb.UnimplementedBlogServiceServer
}

func (s *blogSvcServer) NewPost(ctx context.Context, req *pb.PostRequest) (res *pb.PostResponse, err error) {
	return
}

func (s *blogSvcServer) NewComment(ctx context.Context, req *pb.CommentRequest) (res *pb.CommentResponse, err error) {

	return
}

func (s *blogSvcServer) EditComment(ctx context.Context, req *pb.CommentRequest) (res *pb.CommentResponse, err error) {

	return
}

func (s *blogSvcServer) DeleteCommment(ctx context.Context, req *pb.CommentRequest) (res *pb.CommentResponse, err error) {

	return
}

func (s *blogSvcServer) Reply(ctx context.Context, req *pb.CommentRequest) (res *pb.CommentResponse, err error) {

	return
}

func (s *blogSvcServer) React(ctx context.Context, req *pb.ReactRequest) (res *pb.ReactResponse, err error) {

	return
}

func NewBlogService() *blogSvcServer {
	return nil
}

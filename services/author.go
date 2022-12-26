package services

import (
	"context"

	"github.com/rommms07/blogs/pb"
)

type authorSvcServer struct {
	pb.UnimplementedAuthorServiceServer
}

func (s *authorSvcServer) NewPost(ctx context.Context, req *pb.AuthorRequest) (res *pb.AuthorResponse, err error) {

	return
}

func (s *authorSvcServer) EditPost(ctx context.Context, req *pb.AuthorRequest) (res *pb.AuthorResponse, err error) {

	return
}

func (s *authorSvcServer) DeletePost(ctx context.Context, req *pb.AuthorRequest) (res *pb.AuthorResponse, err error) {

	return
}

func NewAuthorService() *authorSvcServer {
	return nil
}

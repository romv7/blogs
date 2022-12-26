package services

import (
	"context"

	"github.com/rommms07/blogs/pb"
)

type userSvcServer struct {
	pb.UnimplementedUserServiceServer
}

func (s *userSvcServer) NewUser(ctx context.Context, req *pb.AuthorRequest) (res *pb.AuthorResponse, err error) {

	return
}

func (s *userSvcServer) EditUser(ctx context.Context, req *pb.AuthorRequest) (res *pb.AuthorResponse, err error) {

	return
}

func (s *userSvcServer) DeleteUser(ctx context.Context, req *pb.AuthorRequest) (res *pb.AuthorResponse, err error) {

	return
}

func NewUserService() *userSvcServer {
	return nil
}

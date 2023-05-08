package grpcTest

import (
	"github.com/romv7/blogs/internal/pb"
	"google.golang.org/grpc"
)

type grpcTestCase struct {
	For        int
	testValues map[string]any
}

type grpcTestCases []*grpcTestCase

const (
	BlogServiceTest = iota
	UserServiceTest
)

func dialGrpcClientConn() *grpc.ClientConn {
	conn, err := grpc.Dial(ListenerAddr, grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	return conn
}

func createTestBlogClient() (conn *grpc.ClientConn, result pb.BlogServiceClient) {
	conn = dialGrpcClientConn()
	result = pb.NewBlogServiceClient(conn)
	return
}

func createTestUserClient() (conn *grpc.ClientConn, result pb.UserServiceClient) {
	conn = dialGrpcClientConn()
	result = pb.NewUserServiceClient(conn)
	return
}

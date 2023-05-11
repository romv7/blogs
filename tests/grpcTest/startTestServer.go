package grpcTest

import (
	"log"
	"net"

	"google.golang.org/grpc"

	endpoint "github.com/romv7/blogs/endpoints/grpc"
	"github.com/romv7/blogs/internal/pb"
)

var (
	lis          net.Listener
	ListenerAddr = ":5000"
)

func startTestServer(service interface{}) {
	var err error

	lis, err = net.Listen("tcp", ListenerAddr)
	if err != nil {
		log.Panic(err)
	}

	grpcServer := grpc.NewServer()

	switch S := service.(type) {
	case *endpoint.B:
		pb.RegisterBlogServiceServer(grpcServer, S)
	case *endpoint.U:
		pb.RegisterUserServiceServer(grpcServer, S)
	}

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			grpcServer.Stop()
		}
	}()

}

func closeTestServer() {
	lis.Close()
}

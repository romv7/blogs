package grpcTest

import (
	"testing"

	"github.com/romv7/blogs/endpoints/grpc"
)

func TestNewBlogPost(t *testing.T) {
	b := &grpc.B{}

	startTestServer(b)
	defer closeTestServer()

	conn, _ := createTestBlogClient()
	defer conn.Close()

}

func TestSaveBlogPost(t *testing.T) {

}

func TestDeleteBlogPost(t *testing.T) {

}

func TestNewComment(t *testing.T) {

}

func TestSaveComment(t *testing.T) {

}

func TestDeleteComment(t *testing.T) {

}

func TestReplyToComment(t *testing.T) {

}

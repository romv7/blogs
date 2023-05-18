package grpcTest

import (
	"context"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/google/uuid"
	endpoint "github.com/romv7/blogs/endpoints/grpc"
	"github.com/romv7/blogs/internal/pb"
	"github.com/romv7/blogs/internal/storage"
	"github.com/romv7/blogs/internal/store"
	"github.com/romv7/blogs/internal/utils"
	"github.com/romv7/blogs/internal/utils/author"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func saveUser(au map[string]any) (u *store.User, err error) {
	ustore := store.NewUserStore(store.SqlStore)

	u = ustore.NewUser(&pb.User{
		Id:       utils.RandomUniqueId(),
		Uuid:     uuid.NewString(),
		Name:     au["name"].(string),
		FullName: au["full_name"].(string),
		Email:    au["email"].(string),
		Type:     au["type"].(pb.User_Type),
		State: &pb.UserState{
			CreatedAt: timestamppb.Now(),
			UpdatedAt: nil,
			Disabled:  false,
			UVerified: true,
		},
	})

	err = u.Save()

	return
}

func TestNewBlogPost(t *testing.T) {
	b := &endpoint.BlogService{}

	startTestServer(b)
	defer closeTestServer()

	conn, client := createTestBlogClient()
	defer conn.Close()

	for _, tcase := range globalTestCases {
		if tcase.For != BlogServiceTest {
			continue
		}

		u, err := saveUser(tcase.testValues["user"].(map[string]any))
		if err != nil {
			t.Error(err)
		}

		defer u.Delete()

		expectedContent := "# This is an example blog."

		p, err := client.NewBlogPost(context.TODO(), &pb.BlogService_NewBlogPost_Params{
			User:         u.Proto(),
			HeadlineText: tcase.testValues["headline_text"].(string),
			SummaryText:  tcase.testValues["summary_text"].(string),
			Tags:         tcase.testValues["tags"].([]string),
			Images:       tcase.testValues["images"].([]string),
			Refs:         tcase.testValues["refs"].([]string),
			Attachments:  tcase.testValues["attachments"].([]string),
			Content:      expectedContent,
		})

		if err != nil {
			t.Error(err)
		}

		if len(p.Uuid) == 0 {
			t.Error("did not create a uuid for the blog post.")
		}

		if strings.Compare(p.HeadlineText, tcase.testValues["headline_text"].(string)) != 0 {
			t.Error("did not match the expected headline.")
		}

		if strings.Compare(p.Content, expectedContent) != 0 {
			t.Error("did not match the expected content.")
		}

		if p.User.Id != u.Proto().Id {
			t.Error("did not match the expected user id.")
		}

	}

}

func TestSaveBlogPost(t *testing.T) {
	b := &endpoint.BlogService{}

	startTestServer(b)
	defer closeTestServer()

	conn, client := createTestBlogClient()
	defer conn.Close()

	pstore := store.NewPostStore(store.SqlStore)

	for _, tcase := range globalTestCases {
		if tcase.For != BlogServiceTest {
			continue
		}

		u, err := saveUser(tcase.testValues["user"].(map[string]any))
		if err != nil {
			t.Error(err)
		}

		defer u.Delete()

		expectedContent := "# This is an example blog."

		p, err := client.NewBlogPost(context.TODO(), &pb.BlogService_NewBlogPost_Params{
			User:         u.Proto(),
			HeadlineText: tcase.testValues["headline_text"].(string),
			SummaryText:  tcase.testValues["summary_text"].(string),
			Tags:         tcase.testValues["tags"].([]string),
			Images:       tcase.testValues["images"].([]string),
			Refs:         tcase.testValues["refs"].([]string),
			Attachments:  tcase.testValues["attachments"].([]string),
			Content:      expectedContent,
		})

		if err != nil {
			t.Error(err)
		}

		_, err = client.SaveBlogPost(context.TODO(), &pb.BlogService_SaveBlogPost_Params{
			User: u.Proto(),
			Post: p,
		})

		if err != nil {
			t.Error(err)
		}

		defer pstore.NewPost(u.Proto(), p).Delete()

		ah := author.NewAuthorHelper(u.Proto(), storage.Plain)
		rootPath := u.Proto().StoragePath
		blogKey := ah.GetBlogPostFileKey(p.Uuid)
		path := fmt.Sprintf("%s/%s/%s", os.Getenv("STORAGE_DIR"), rootPath, blogKey)

		// Check if there is a file containing the blog content saved in the file system.
		if _, err := os.Stat(path); os.IsNotExist(err) {
			t.Error("did not stored the content of the blog in the file system.")
		}

	}

}

func TestDeleteBlogPost(t *testing.T) {
	b := &endpoint.BlogService{}

	startTestServer(b)
	defer closeTestServer()

	conn, client := createTestBlogClient()
	defer conn.Close()

	for _, tcase := range globalTestCases {
		if tcase.For != BlogServiceTest {
			continue
		}

		u, err := saveUser(tcase.testValues["user"].(map[string]any))
		if err != nil {
			t.Error(err)
		}

		defer u.Delete()

		expectedContent := "# This is an example blog."

		p, err := client.NewBlogPost(context.TODO(), &pb.BlogService_NewBlogPost_Params{
			User:         u.Proto(),
			HeadlineText: tcase.testValues["headline_text"].(string),
			SummaryText:  tcase.testValues["summary_text"].(string),
			Tags:         tcase.testValues["tags"].([]string),
			Images:       tcase.testValues["images"].([]string),
			Refs:         tcase.testValues["refs"].([]string),
			Attachments:  tcase.testValues["attachments"].([]string),
			Content:      expectedContent,
		})

		// Save the post
		_, err = client.SaveBlogPost(context.TODO(), &pb.BlogService_SaveBlogPost_Params{
			User: u.Proto(),
			Post: p,
		})
		if err != nil {
			t.Error(err)
		}

		_, err = client.DeleteBlogPost(context.TODO(), &pb.BlogService_DeleteBlogPost_Params{
			Post: p,
		})

		if err != nil {
			t.Error(err)
		}

		ah := author.NewAuthorHelper(u.Proto(), storage.Plain)
		rootPath := u.Proto().StoragePath
		blogKey := ah.GetBlogPostFileKey(p.Uuid)
		path := fmt.Sprintf("%s/%s/%s", os.Getenv("STORAGE_DIR"), rootPath, blogKey)

		// Check if there is a file containing the blog content saved in the file system.
		if _, err := os.Stat(path); os.IsExist(err) {
			t.Error("did not deleted the content of the blog stored in the file system.")
		}

	}

}

func TestGetBlogPost(t *testing.T) {
	b := &endpoint.BlogService{}

	startTestServer(b)
	defer closeTestServer()

	conn, _ := createTestBlogClient()
	defer conn.Close()

	for _, tcase := range globalTestCases {
		if tcase.For != BlogServiceTest {
			continue
		}

		u, err := saveUser(tcase.testValues["user"].(map[string]any))
		if err != nil {
			t.Error(err)
		}

		defer u.Delete()
	}

}

func getRandomGlobalGrpcTestIndex() uint64 {
	return utils.RandomUniqueId() % uint64(len(globalTestCases))
}

func TestNewComment(t *testing.T) {
	b := &endpoint.BlogService{}

	startTestServer(b)
	defer closeTestServer()

	conn, client := createTestBlogClient()
	defer conn.Close()

	pstore := store.NewPostStore(store.SqlStore)

	for _, tcase := range globalTestCases {
		if tcase.For != BlogServiceTest_Comment {
			continue
		}

		u, err := saveUser(tcase.testValues["user"].(map[string]any))
		if err != nil {
			t.Error(err)
		}

		defer u.Delete()

		// Select a target
		target := globalTestCases[getRandomGlobalGrpcTestIndex()]

		// Save the owner of the selected target. But before that, we must ensure
		// that the owner of the post is not the same as the user who created
		// the comment. Skip this test if it is.
		if target.For == UserServiceTest || target.For == BlogServiceTest_Comment {
			continue
		} else if target.testValues["user"].(map[string]any)["email"] == u.Proto().Email {
			continue
		}

		tu, err := saveUser(target.testValues["user"].(map[string]any))
		if err != nil {
			t.Error(err)
		}

		defer tu.Delete()

		var params *pb.BlogService_NewComment_Params

		switch target.For {
		case BlogServiceTest:
			params = &pb.BlogService_NewComment_Params{
				TargetType: pb.Comment_TT_POST,
			}

			p, err := client.NewBlogPost(context.TODO(), &pb.BlogService_NewBlogPost_Params{
				User:         tu.Proto(),
				HeadlineText: target.testValues["headline_text"].(string),
				SummaryText:  target.testValues["summary_text"].(string),
				Tags:         target.testValues["tags"].([]string),
				Images:       target.testValues["images"].([]string),
				Attachments:  target.testValues["attachments"].([]string),
				Refs:         target.testValues["refs"].([]string),
				Content:      "",
			})

			if err != nil {
				t.Error(err)
			}

			if err = pstore.NewPost(tu.Proto(), p).Save(); err != nil {
				t.Error(err)
			}

			defer pstore.NewPost(tu.Proto(), p).Delete()

			params.TargetUuid = p.Uuid
		}

		params.User = u.Proto()
		params.CommentText = tcase.testValues["comment_text"].(string)

		c, err := client.NewComment(context.TODO(), params)
		if err != nil {
			t.Error(err)
		}

		if c.CommentText.Data != params.CommentText {
			t.Error("did not match the expected comment text.")
		}

	}
}

func TestSaveComment(t *testing.T) {
	b := &endpoint.BlogService{}

	startTestServer(b)
	defer closeTestServer()

	conn, client := createTestBlogClient()
	defer conn.Close()

	pstore := store.NewPostStore(store.SqlStore)

	for _, tcase := range globalTestCases {
		if tcase.For != BlogServiceTest_Comment {
			continue
		}

		u, err := saveUser(tcase.testValues["user"].(map[string]any))
		if err != nil {
			t.Error(err)
		}

		defer u.Delete()

		// Select a target
		target := globalTestCases[getRandomGlobalGrpcTestIndex()]

		// Save the owner of the selected target. But before that, we must ensure
		// that the owner of the post is not the same as the user who created
		// the comment. Skip this test if it is.
		if target.For == UserServiceTest || target.For == BlogServiceTest_Comment {
			continue
		} else if target.testValues["user"].(map[string]any)["email"] == u.Proto().Email {
			continue
		}

		tu, err := saveUser(target.testValues["user"].(map[string]any))
		if err != nil {
			t.Error(err)
		}

		defer tu.Delete()

		var params *pb.BlogService_NewComment_Params

		switch target.For {
		case BlogServiceTest:
			params = &pb.BlogService_NewComment_Params{
				TargetType: pb.Comment_TT_POST,
			}

			p, err := client.NewBlogPost(context.TODO(), &pb.BlogService_NewBlogPost_Params{
				User:         tu.Proto(),
				HeadlineText: target.testValues["headline_text"].(string),
				SummaryText:  target.testValues["summary_text"].(string),
				Images:       target.testValues["images"].([]string),
				Attachments:  target.testValues["attachments"].([]string),
				Refs:         target.testValues["refs"].([]string),
				Content:      "",
			})

			if err != nil {
				t.Error(err)
			}

			if err = pstore.NewPost(tu.Proto(), p).Save(); err != nil {
				t.Error(err)
			}

			defer pstore.NewPost(tu.Proto(), p).Delete()

			params.TargetUuid = p.Uuid
		}

		params.User = u.Proto()
		params.CommentText = tcase.testValues["comment_text"].(string)

		c, err := client.NewComment(context.TODO(), params)
		if err != nil {
			t.Error(err)
		}

		// Save the comment.
		_, err = client.SaveComment(context.TODO(), &pb.BlogService_SaveComment_Params{Comment: c})
		if err != nil {
			t.Error(err)
		}

	}

}

func TestDeleteComment(t *testing.T) {
	b := &endpoint.BlogService{}

	startTestServer(b)
	defer closeTestServer()

	conn, client := createTestBlogClient()
	defer conn.Close()

	pstore := store.NewPostStore(store.SqlStore)

	for _, tcase := range globalTestCases {
		if tcase.For != BlogServiceTest_Comment {
			continue
		}

		u, err := saveUser(tcase.testValues["user"].(map[string]any))
		if err != nil {
			t.Error(err)
		}

		defer u.Delete()

		// Select a target
		target := globalTestCases[getRandomGlobalGrpcTestIndex()]

		// Save the owner of the selected target. But before that, we must ensure
		// that the owner of the post is not the same as the user who created
		// the comment. Skip this test if it is.
		if target.For == UserServiceTest || target.For == BlogServiceTest_Comment {
			continue
		} else if target.testValues["user"].(map[string]any)["email"] == u.Proto().Email {
			continue
		}

		tu, err := saveUser(target.testValues["user"].(map[string]any))
		if err != nil {
			t.Error(err)
		}

		defer tu.Delete()

		var params *pb.BlogService_NewComment_Params

		switch target.For {
		case BlogServiceTest:
			params = &pb.BlogService_NewComment_Params{
				TargetType: pb.Comment_TT_POST,
			}

			p, err := client.NewBlogPost(context.TODO(), &pb.BlogService_NewBlogPost_Params{
				User:         tu.Proto(),
				HeadlineText: target.testValues["headline_text"].(string),
				SummaryText:  target.testValues["summary_text"].(string),
				Images:       target.testValues["images"].([]string),
				Attachments:  target.testValues["attachments"].([]string),
				Refs:         target.testValues["refs"].([]string),
				Content:      "",
			})

			if err != nil {
				t.Error(err)
			}

			if err = pstore.NewPost(tu.Proto(), p).Save(); err != nil {
				t.Error(err)
			}

			defer pstore.NewPost(tu.Proto(), p).Delete()

			params.TargetUuid = p.Uuid
		}

		params.User = u.Proto()
		params.CommentText = tcase.testValues["comment_text"].(string)

		c, err := client.NewComment(context.TODO(), params)
		if err != nil {
			t.Error(err)
		}

		// Save the comment.
		_, err = client.SaveComment(context.TODO(), &pb.BlogService_SaveComment_Params{Comment: c})
		if err != nil {
			t.Error(err)
		}

		// Delete the comment using the gRPC service (DeleteComment)
		_, err = client.DeleteComment(context.TODO(), &pb.BlogService_DeleteComment_Params{Comment: c})
		if err != nil {
			t.Error(err)
		}
	}

}

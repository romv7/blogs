package store_test

import (
	"testing"
	"time"

	"github.com/rommms07/blogs/internal/entities"
	"github.com/rommms07/blogs/internal/store"
	"github.com/rommms07/blogs/pb"
)

var (
	globalStoreErr error
	globalFakeUser = &entities.User{
		User: &pb.User{
			Id:        0,
			Name:      "dustybroom0",
			FullName:  "George Orwell Jr.",
			Disabled:  false,
			CreatedAt: uint64(time.Now().Unix()),
		},
	}
)

var (
	commentStore *store.CommentStore
)

func init() {
	commentStore, globalStoreErr = store.NewCommentStore(globalFakeUser)
}

func Test_theTestCommentStoreMustNowBeDefined(t *testing.T) {
	if commentStore == nil {
		t.Errorf("[fail] commentStore was not defined properly.. (error: %s)", globalStoreErr.Error())
	}
}

func Test_mustInstantiateANewComment(t *testing.T) {
	message := "You should have created a neat service for that before you had proceeded at making one."
	comment := commentStore.New(message, 0, pb.Comment_T_POST)

	if comment.User != globalFakeUser {
		t.Errorf("[fail] comment.New did not matched the expected creator.")
	}

	if comment.CommentText != message {
		t.Errorf("[fail] comment.New content did not matched the expected message.")
	}
}

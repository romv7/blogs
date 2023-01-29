package store_test

import (
	"testing"

	"google.golang.org/protobuf/types/known/timestamppb"
	"github.com/rommms07/blogs/internal/entities"
	"github.com/rommms07/blogs/internal/store"
	"github.com/google/uuid"
	"github.com/rommms07/blogs/pb"
)

var (
	globalStoreErr error
	globalFakeUser = &entities.User{
		User: &pb.User{
			Uuid:      uuid.New().String(),
			Name:      "dustybroom0",	
			FullName:  "George Orwell Jr.",
			State:     &pb.UserState {
				Disabled: false,
				CreatedAt: timestamppb.Now(),
			},
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
	comment := commentStore.New(message, globalFakeUser.Uuid)

	if comment.User != globalFakeUser {
		t.Errorf("[fail] comment.New did not matched the expected creator.")
	}

	if comment.CommentText != message {
		t.Errorf("[fail] comment.New content did not matched the expected message.")
	}
}

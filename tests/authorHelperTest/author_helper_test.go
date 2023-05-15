package authorHelperTest

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/romv7/blogs/internal/pb"
	"github.com/romv7/blogs/internal/store"
	"github.com/romv7/blogs/internal/utils"
	"github.com/romv7/blogs/internal/utils/author"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	mockUser = store.NewUserStore(store.SqlStore).NewUser(&pb.User{
		Id:       utils.RandomUniqueId() + uint64(time.Now().Unix()),
		Uuid:     uuid.NewString(),
		Name:     "rom123",
		FullName: "Rom Vales Villanueva",
		Email:    "romlas@gmail.com",
		Type:     pb.User_T_AUTHOR,
		State: &pb.UserState{
			CreatedAt: timestamppb.Now(),
			UpdatedAt: nil,
			Disabled:  false,
			UVerified: false,
		},
	})
)

func TestShouldCreateNewAuthorHelper(t *testing.T) {
	if err := mockUser.Save(); err != nil {
		t.Error(err)
	}

	defer mockUser.Delete()

	if uh := author.NewAuthorHelper(mockUser.Proto(), author.StoragePlain); uh == nil {
		t.Errorf("unable to create author helper")
	} else {

	}
}

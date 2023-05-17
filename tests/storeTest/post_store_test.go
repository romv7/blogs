package storeTest

import (
	"testing"

	"github.com/romv7/blogs/internal/pb"
	"github.com/romv7/blogs/internal/store"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestNewPostStore(t *testing.T) {
	if store.NewPostStore(store.SqlStore) == nil {
		t.Fatalf("NewPostStore returned an invalid store")
	}
}

func TestPostStore_SqlStore(t *testing.T) {
	pstore := store.NewPostStore(store.SqlStore)
	ustore := store.NewUserStore(store.SqlStore)

	for _, posts := range globalPostTestCases {
		u := ustore.NewUser(posts.u)

		if err := u.Save(); err != nil {
			t.Fatal(err)
		}

		defer u.Delete()

		for _, post := range posts.p {
			post.State = &pb.PostState{
				Stage:       pb.PostState_S_WIP,
				Status:      pb.PostState_S_DRAFT,
				RevisedAt:   timestamppb.Now(),
				PublishedAt: nil,
				ArchivedAt:  nil,
				CreatedAt:   timestamppb.Now(),
				Reacts:      &pb.Reacts{},
			}

			p := pstore.NewPost(u.Proto(), post)

			if err := p.Save(); err != nil {
				t.Fatal(err)
			}

			defer p.Delete()
		}
	}

}

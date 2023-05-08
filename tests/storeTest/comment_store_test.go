package storeTest

import (
	"encoding/json"
	"io/ioutil"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/romv7/blogs/internal/pb"
	"github.com/romv7/blogs/internal/store"
	"github.com/romv7/blogs/internal/utils"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestNewCommentStore(t *testing.T) {
	if store.NewCommentStore(store.SqlStore) == nil {
		t.Fatalf("NewCommentStore returned an invalid store")
	}
}

func createTestUser(fname, name, email string) *pb.User {
	return &pb.User{
		Id:       uint32(utils.RandomUniqueId()) + uint32(time.Now().Unix()),
		Uuid:     uuid.NewString(),
		Name:     name,
		FullName: fname,
		Email:    email,
		Type:     pb.User_T_AUTHOR,
		State: &pb.UserState{
			CreatedAt: timestamppb.Now(),
			UpdatedAt: timestamppb.Now(),
			Disabled:  false,
			UVerified: true,
		},
	}
}

func TestCommentStore_SqlStore(t *testing.T) {
	ustore := store.NewUserStore(store.SqlStore)
	pstore := store.NewPostStore(store.SqlStore)
	cstore := store.NewCommentStore(store.SqlStore)

	// Supposed that there are two users in the server.
	//
	// 1. Rom Vales Villanueva (me)
	// 2. Emil Cioran
	me := ustore.NewUser(createTestUser("Rom Vales Villanueva", "romdevmod", "rommms@gmail.com"))
	cioran := ustore.NewUser(createTestUser("Emil Cioran", "emilcioran", "emil@gmail.com"))

	ustore.Save(me)
	ustore.Save(cioran)
	defer ustore.Delete(me)
	defer ustore.Delete(cioran)

	// Now Emil Cioran decided to create an existenstial blog post in my wonderful
	// server.
	post := pstore.NewPost(cioran.Proto(), &pb.Post{
		Id:           utils.RandomUniqueId() + uint32(time.Now().Unix()),
		Uuid:         uuid.NewString(),
		HeadlineText: "An Overview of my work \"On the Heights of Despair\"",
		SummaryText:  "",
		Tags:         &pb.Tags{Data: []string{"existenstialism", "philosophy", "nihilism"}},
		State: &pb.PostState{
			Stage:       pb.PostState_S_WIP,
			Status:      pb.PostState_S_DRAFT,
			CreatedAt:   timestamppb.Now(),
			PublishedAt: nil,
			RevisedAt:   nil,
			ArchivedAt:  nil,
			Reacts:      &pb.Reacts{},
		},
	})

	pstore.Save(post)
	defer pstore.Delete(post)

	// After a few hours, I decided to open up my blog sharing website and found out
	// that Emil Cioran created a post!! I immediately read his post and astonished
	// by how insightful it was. I decided to comment.
	me_comment := cstore.NewComment(&pb.Comment{
		Id:          utils.RandomUniqueId() + uint32(time.Now().Unix()),
		Uuid:        uuid.NewString(),
		User:        me.Proto(),
		CommentText: &pb.CommentText{Data: "It seems reasonable that existence actually doesn't have any meaning as its essence. In fact, if you would weigh the importance of life to a lifeless rock is futile."},
		State: &pb.CommentState{
			CreatedAt: timestamppb.Now(),
			EditedAt:  nil,
		},
	}, post.Proto().Uuid, pb.Comment_TT_POST)

	me_comment2 := cstore.NewComment(&pb.Comment{
		Id:          utils.RandomUniqueId() + uint32(time.Now().Unix()),
		Uuid:        uuid.NewString(),
		User:        me.Proto(),
		CommentText: &pb.CommentText{Data: "I can't wait for the next part of this blog!"},
		State: &pb.CommentState{
			CreatedAt: timestamppb.Now(),
			EditedAt:  nil,
		},
	}, post.Proto().Uuid, pb.Comment_TT_POST)

	cstore.Save(me_comment)
	cstore.Save(me_comment2)
	defer cstore.Delete(me_comment)
	defer cstore.Delete(me_comment2)

	// I was further astonished when Emil Cioran replied to me with a very
	// philosophical message, that I couldn't even grasp what he means.
	cioran_reply := cstore.NewComment(&pb.Comment{
		Id:          utils.RandomUniqueId() + uint32(time.Now().Unix()),
		Uuid:        uuid.NewString(),
		User:        cioran.Proto(),
		CommentText: &pb.CommentText{Data: "Life is a foul-scented plastic floating in the sea."},
		State: &pb.CommentState{
			CreatedAt: timestamppb.Now(),
			EditedAt:  nil,
		},
	}, me_comment.Proto().Uuid, pb.Comment_TT_COMMENT)

	// And the conversation abruptly ends here....

	cstore.Save(cioran_reply)
	defer cstore.Delete(cioran_reply)

	p, _ := json.Marshal(post.Proto())
	ioutil.WriteFile("/tmp/cioran-blog.json", p, 0644)
}

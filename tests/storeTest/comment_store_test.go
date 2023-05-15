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
		Id:       uint64(utils.RandomUniqueId()) + uint64(time.Now().Unix()),
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
	me := ustore.NewUser(createTestUser("Rom Vales Villanueva", "romdevmod", "rommms12@gmail.com"))
	cioran := ustore.NewUser(createTestUser("Emil Cioran", "emilcioran", "emil@gmail.com"))

	me.Save()
	cioran.Save()
	defer me.Delete()
	defer cioran.Delete()

	// Now Emil Cioran decided to create an existenstial blog post in my wonderful
	// server.
	post := pstore.NewPost(cioran.Proto(), &pb.Post{
		Id:           utils.RandomUniqueId() + uint64(time.Now().Unix()),
		Uuid:         uuid.NewString(),
		HeadlineText: "An Overview of my work \"On the Heights of Despair\"",
		SummaryText:  "\"On the Heights of Despair\" is a philosophical essay written by Romanian philosopher Emil Cioran. In this book, Cioran explores the nature of despair and the human condition, questioning the value of life and the possibility of meaning in a world that is fundamentally absurd. Through a series of aphoristic and poetic reflections, he confronts the most profound and troubling aspects of existence, offering a vision of life that is at once pessimistic and exhilarating.",
		Tags:         &pb.Tags{Data: []string{"existenstialism", "philosophy", "nihilism"}},
		State: &pb.PostState{
			Stage:       pb.PostState_S_WIP,
			Status:      pb.PostState_S_DRAFT,
			CreatedAt:   timestamppb.Now(),
			PublishedAt: nil,
			RevisedAt:   nil,
			ArchivedAt:  nil,
			Reacts:      &pb.Reacts{LikeCount: 1, LoveCount: 102},
		},
	})

	post.Save()
	defer post.Delete()

	// After a few hours, I decided to open up my blog sharing website and found out
	// that Emil Cioran created a post!! I immediately read his post and astonished
	// by how insightful it was. I decided to comment.
	me_comment := cstore.NewComment(&pb.Comment{
		Id:          utils.RandomUniqueId() + uint64(time.Now().Unix()),
		Uuid:        uuid.NewString(),
		User:        me.Proto(),
		CommentText: &pb.CommentText{Data: "It seems reasonable that existence actually doesn't have any meaning as its essence. In fact, if you would weigh the importance of life to a lifeless rock is futile."},
		State: &pb.CommentState{
			CreatedAt: timestamppb.Now(),
			EditedAt:  nil,
		},
	}, post.Proto().Uuid, pb.Comment_TT_POST)

	me_comment2 := cstore.NewComment(&pb.Comment{
		Id:          utils.RandomUniqueId() + uint64(time.Now().Unix()),
		Uuid:        uuid.NewString(),
		User:        me.Proto(),
		CommentText: &pb.CommentText{Data: "I can't wait for the next part of this blog!"},
		State: &pb.CommentState{
			CreatedAt: timestamppb.Now(),
			EditedAt:  nil,
		},
	}, post.Proto().Uuid, pb.Comment_TT_POST)

	me_comment.Save()
	me_comment2.Save()
	defer me_comment.Delete()
	defer me_comment2.Delete()

	// I was further astonished when Emil Cioran replied to me with a very
	// philosophical message, that I couldn't even grasp what he means.
	cioran_reply := cstore.NewComment(&pb.Comment{
		Id:          utils.RandomUniqueId() + uint64(time.Now().Unix()),
		Uuid:        uuid.NewString(),
		User:        cioran.Proto(),
		CommentText: &pb.CommentText{Data: "Life is a foul-scented plastic floating in the sea."},
		State: &pb.CommentState{
			CreatedAt: timestamppb.Now(),
			EditedAt:  nil,
		},
	}, me_comment.Proto().Uuid, pb.Comment_TT_COMMENT)

	// And the conversation abruptly ends here....

	cioran_reply.Save()
	defer cioran_reply.Delete()

	P, err := pstore.GetByUuid(post.Proto().Uuid)
	if err != nil {
		t.Error(err)
	}

	byt, _ := json.Marshal(P.Proto())
	ioutil.WriteFile("/tmp/post.json", byt, 0644)
}

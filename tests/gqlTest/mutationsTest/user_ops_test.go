package mutationsTest

import (
	"context"
	"encoding/json"
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/romv7/blogs/endpoints/gql"
	"github.com/romv7/blogs/internal/pb"
	"github.com/romv7/blogs/internal/store"
	"github.com/romv7/blogs/internal/utils"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

func TestUserOpsCreatePost(t *testing.T) {
	S := gql.DefaultSchema()
	q := `
	mutation UserOpsForCreatePost($post: UserOpsPostMutationsParameter!) {
		userOps {
			createPost(input: $post) {
				op
				code
				message
				uuid
			}
		}
	}
	`

	owner := &pb.User{
		Id:       utils.RandomUniqueId(),
		Uuid:     uuid.NewString(),
		Name:     "quirky34bastard",
		FullName: "Quirky 34 Bastard",
		Email:    "quirk34@gmail.com",
		Type:     pb.User_T_AUTHOR,
		State: &pb.UserState{
			CreatedAt: timestamppb.Now(),
			UpdatedAt: timestamppb.Now(),
			Disabled:  false,
			UVerified: true,
		},
	}

	ustore := store.NewUserStore(store.SqlStore)
	pstore := store.NewPostStore(store.SqlStore)

	if err := ustore.NewUser(owner).Save(); err != nil {
		t.Error(err)
	}

	defer ustore.NewUser(owner).Delete()

	empty := map[string]any{"id": float64(owner.Id)}

	res := S.Exec(context.TODO(), q, "UserOpsForCreatePost", map[string]any{
		"post": empty,
	})

	if len(res.Errors) > 0 {
		t.Errorf("%+v", res.Errors)
	}

	body := map[string]any{}

	if err := json.Unmarshal(res.Data, &body); err != nil {
		t.Fatal(err)
	}

	postUuid := body["userOps"].(map[string]any)["createPost"].(map[string]any)["uuid"].(string)

	var post *store.Post

	if P, err := pstore.GetByUuid(postUuid); err != nil {
		t.Error(err)
	} else if P.Proto().HeadlineText != "" {
		t.Error("did not match the expected headline text.")
	} else {
		post = P
	}

	if post.Proto().Uuid != postUuid {
		t.Error("did not match the expected post uuid.")
	}

	if err := post.Delete(); err != nil {
		t.Error(err)
	}
}

func TestUserOpsUpdatePost(t *testing.T) {
	S := gql.DefaultSchema()
	q := `
	mutation UserOpsForUpdatePost($post: UserOpsPostMutationsParameter!) {
		userOps {
			updatePost(input: $post) {
				op
				code
				message
				uuid
			}
		}
	}
	`

	owner := &pb.User{
		Id:       utils.RandomUniqueId(),
		Uuid:     uuid.NewString(),
		Name:     "quirky36bastard",
		FullName: "Quirky 36 Bastard",
		Email:    "quirk34@gmail.com",
		Type:     pb.User_T_AUTHOR,
		State: &pb.UserState{
			CreatedAt: timestamppb.Now(),
			UpdatedAt: timestamppb.Now(),
			Disabled:  false,
			UVerified: true,
		},
	}

	post := &pb.Post{
		Id:           utils.RandomUniqueId(),
		Uuid:         uuid.NewString(),
		HeadlineText: "Example Blog Post",
		SummaryText:  "Just another example blog post for a unit test.",
		User:         owner,
		Tags:         &pb.Tags{Data: []string{"example", "test"}},
		State: &pb.PostState{
			Stage:       pb.PostState_S_PUB,
			Status:      pb.PostState_S_VISIBLE,
			RevisedAt:   nil,
			ArchivedAt:  nil,
			PublishedAt: timestamppb.Now(),
			CreatedAt:   timestamppb.Now(),
			Reacts:      &pb.Reacts{LikeCount: 13},
		},
	}

	ustore := store.NewUserStore(store.SqlStore)
	pstore := store.NewPostStore(store.SqlStore)

	if err := ustore.NewUser(owner).Save(); err != nil {
		t.Error(err)
	}

	defer ustore.NewUser(owner).Delete()

	if err := pstore.NewPost(owner, post).Save(); err != nil {
		t.Error(err)
	}

	defer pstore.NewPost(owner, post).Delete()

	var (
		expectedHeadlineText = "Just another post that was updated through a test."
	)

	updates := map[string]any{
		"id":           float64(owner.Id),
		"uuid":         post.Uuid,
		"headlineText": expectedHeadlineText,
	}

	res := S.Exec(context.TODO(), q, "UserOpsForUpdatePost", map[string]any{
		"post": updates,
	})

	if len(res.Errors) > 0 {
		t.Errorf("%+v", res.Errors)
	}

	body := map[string]any{}

	if err := json.Unmarshal(res.Data, &body); err != nil {
		t.Fatal(err)
	}

	postUuid := body["userOps"].(map[string]any)["updatePost"].(map[string]any)["uuid"].(string)

	upPost, err := pstore.GetByUuid(postUuid)
	if err != nil {
		t.Error(err)
	}

	if upPost.Proto().HeadlineText != expectedHeadlineText {
		t.Error("did not match the expected headline text.")
	}

}

func TestUserOpsDeletePost(t *testing.T) {
	S := gql.DefaultSchema()
	q := `
	mutation UserOpsForDeletePost($postUuid: String!) {
		userOps {
			deletePost(uuid: $postUuid) {
				op
				code
				message
				uuid
			}
		}
	}
	`

	owner := &pb.User{
		Id:       utils.RandomUniqueId(),
		Uuid:     uuid.NewString(),
		Name:     "quirkybastard",
		FullName: "Quirky Bastard",
		Email:    "quirk@gmail.com",
		Type:     pb.User_T_AUTHOR,
		State: &pb.UserState{
			CreatedAt: timestamppb.Now(),
			UpdatedAt: timestamppb.Now(),
			Disabled:  false,
			UVerified: true,
		},
	}

	post := &pb.Post{
		Id:           utils.RandomUniqueId(),
		Uuid:         uuid.NewString(),
		HeadlineText: "Example Blog Post",
		SummaryText:  "Just another example blog post for a unit test.",
		User:         owner,
		Tags:         &pb.Tags{Data: []string{"example", "test"}},
		State: &pb.PostState{
			Stage:       pb.PostState_S_PUB,
			Status:      pb.PostState_S_VISIBLE,
			RevisedAt:   nil,
			ArchivedAt:  nil,
			PublishedAt: timestamppb.Now(),
			CreatedAt:   timestamppb.Now(),
			Reacts:      &pb.Reacts{LikeCount: 13},
		},
	}

	ustore := store.NewUserStore(store.SqlStore)
	pstore := store.NewPostStore(store.SqlStore)

	if err := ustore.NewUser(owner).Save(); err != nil {
		t.Error(err)
	}

	defer ustore.NewUser(owner).Delete()

	if err := pstore.NewPost(owner, post).Save(); err != nil {
		t.Error(err)
	}

	res := S.Exec(context.TODO(), q, "UserOpsForDeletePost", map[string]any{
		"postUuid": post.Uuid,
	})

	if len(res.Errors) > 0 {
		t.Error(res.Errors)
	}

	body := make(map[string]any)

	if err := json.Unmarshal(res.Data, &body); err != nil {
		t.Fatal(err)
	}

	if P, err := pstore.GetByUuid(post.Uuid); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error(err)
	} else if P != nil {
		t.Error("expected to delete the post.")
	}
}

func TestUserOpsCreateComment(t *testing.T) {
	S := gql.DefaultSchema()
	q := `
	mutation UserOpsForCreateComment($comment: UserOpsCommentMutationsParameter!) {
		userOps {
			createComment(input: $comment) {
				op
				code
				message
				uuid
			}
		}
	}
	`

	owner := &pb.User{
		Id:       utils.RandomUniqueId(),
		Uuid:     uuid.NewString(),
		Name:     "quirkybastard",
		FullName: "Quirky Bastard",
		Email:    "quirk@gmail.com",
		Type:     pb.User_T_AUTHOR,
		State: &pb.UserState{
			CreatedAt: timestamppb.Now(),
			UpdatedAt: timestamppb.Now(),
			Disabled:  false,
			UVerified: true,
		},
	}

	post := &pb.Post{
		Id:           utils.RandomUniqueId(),
		Uuid:         uuid.NewString(),
		HeadlineText: "Example blog post",
		SummaryText:  "Just another example blog post for a unit test.",
		User:         owner,
		Tags:         &pb.Tags{Data: []string{"example", "test"}},
		State: &pb.PostState{
			Stage:       pb.PostState_S_PUB,
			Status:      pb.PostState_S_VISIBLE,
			RevisedAt:   nil,
			ArchivedAt:  nil,
			PublishedAt: timestamppb.Now(),
			CreatedAt:   timestamppb.Now(),
			Reacts:      &pb.Reacts{LikeCount: 103},
		},
	}

	ustore := store.NewUserStore(store.SqlStore)
	pstore := store.NewPostStore(store.SqlStore)
	cstore := store.NewCommentStore(store.SqlStore)

	if err := ustore.NewUser(owner).Save(); err != nil {
		t.Error(err)
	}

	defer ustore.NewUser(owner).Delete()

	if err := pstore.NewPost(owner, post).Save(); err != nil {
		t.Error(err)
	}

	defer pstore.NewPost(owner, post).Delete()

	res := S.Exec(context.TODO(), q, "UserOpsForCreateComment", map[string]any{
		"comment": map[string]any{
			"id":          float64(owner.Id),
			"commentText": "There is a new revision for this post. I'll be publishing it tomorrow!",
			"targetUuid":  post.Uuid,
			"targetType":  pb.Comment_TT_POST.String(),
		},
	})

	if len(res.Errors) > 0 {
		t.Errorf("%+v", res.Errors)
	}

	body := make(map[string]any)

	err := json.Unmarshal(res.Data, &body)
	if err != nil {
		t.Fatal(err)
	}

	commentUuid := body["userOps"].(map[string]any)["createComment"].(map[string]any)["uuid"].(string)

	C, err := cstore.GetByUuid(commentUuid)
	if err != nil {
		t.Error(err)
	}

	defer C.Delete()

	if C.Proto().Uuid != commentUuid {
		t.Error("did not match the expected generated uuid.")
	}

}

func TestUserOpsUpdateComment(t *testing.T) {
	S := gql.DefaultSchema()
	q := `
	mutation UserOpsForUpdateComment($comment: UserOpsCommentMutationsParameter!) {
		userOps {
			updateComment(input: $comment) {
				op
				code
				message
				uuid
			}
		}
	}
	`

	owner := &pb.User{
		Id:       utils.RandomUniqueId(),
		Uuid:     uuid.NewString(),
		Name:     "quirkybastard01",
		FullName: "Quirky Bastard 123",
		Email:    "quir2k@gmail.com",
		Type:     pb.User_T_AUTHOR,
		State: &pb.UserState{
			CreatedAt: timestamppb.Now(),
			UpdatedAt: timestamppb.Now(),
			Disabled:  false,
			UVerified: true,
		},
	}

	post := &pb.Post{
		Id:           utils.RandomUniqueId(),
		Uuid:         uuid.NewString(),
		HeadlineText: "Example blog post",
		SummaryText:  "Just another example blog post for a unit test.",
		User:         owner,
		Tags:         &pb.Tags{Data: []string{"example", "test"}},
		State: &pb.PostState{
			Stage:       pb.PostState_S_PUB,
			Status:      pb.PostState_S_VISIBLE,
			RevisedAt:   nil,
			ArchivedAt:  nil,
			PublishedAt: timestamppb.Now(),
			CreatedAt:   timestamppb.Now(),
			Reacts:      &pb.Reacts{LikeCount: 103},
		},
	}

	comment := &pb.Comment{
		Id:          utils.RandomUniqueId(),
		User:        owner,
		Uuid:        uuid.NewString(),
		CommentText: &pb.CommentText{Data: "There is a new version for this post!"},
		State: &pb.CommentState{
			CreatedAt: timestamppb.Now(),
			EditedAt:  timestamppb.Now(),
			Reacts:    &pb.Reacts{},
		},
	}

	ustore := store.NewUserStore(store.SqlStore)
	pstore := store.NewPostStore(store.SqlStore)
	cstore := store.NewCommentStore(store.SqlStore)

	if err := ustore.NewUser(owner).Save(); err != nil {
		t.Error(err)
	}

	defer ustore.NewUser(owner).Delete()

	if err := pstore.NewPost(owner, post).Save(); err != nil {
		t.Error(err)
	}

	defer pstore.NewPost(owner, post).Delete()

	if err := cstore.NewComment(comment, post.Uuid, pb.Comment_TT_POST).Save(); err != nil {
		t.Error(err)
	}

	defer cstore.NewComment(comment, post.Uuid, pb.Comment_TT_POST).Delete()

	var (
		expectedCommentText = "Updated comment text."
	)

	res := S.Exec(context.TODO(), q, "UserOpsForUpdateComment", map[string]any{
		"comment": map[string]any{
			"id":          float64(owner.Id),
			"uuid":        comment.Uuid,
			"commentText": expectedCommentText,
			"targetUuid":  post.Uuid,
			"targetType":  pb.Comment_TT_POST.String(),
		},
	})

	if len(res.Errors) > 0 {
		t.Errorf("%+v", res.Errors)
	}

	body := make(map[string]any)

	if err := json.Unmarshal(res.Data, &body); err != nil {
		t.Error(err)
	}

	if C, err := cstore.GetById(comment.Id); err != nil {
		t.Error(err)
	} else if C.Proto().CommentText.Data != expectedCommentText {
		t.Error("did not match the expected updated comment text.")
	}
}

func TestUserOpsDeleteComment(t *testing.T) {
	S := gql.DefaultSchema()
	q := `
	mutation UserOpsForDeleteComment($uuid: String!) {
		userOps {
			deleteComment(uuid: $uuid) {
				op
				code
				message
				uuid
			}
		}
	}
	`

	owner := &pb.User{
		Id:       utils.RandomUniqueId(),
		Uuid:     uuid.NewString(),
		Name:     "quirkybastard01",
		FullName: "Quirky Bastard 123",
		Email:    "quir2k@gmail.com",
		Type:     pb.User_T_AUTHOR,
		State: &pb.UserState{
			CreatedAt: timestamppb.Now(),
			UpdatedAt: timestamppb.Now(),
			Disabled:  false,
			UVerified: true,
		},
	}

	post := &pb.Post{
		Id:           utils.RandomUniqueId(),
		Uuid:         uuid.NewString(),
		HeadlineText: "Example blog post",
		SummaryText:  "Just another example blog post for a unit test.",
		User:         owner,
		Tags:         &pb.Tags{Data: []string{"example", "test"}},
		State: &pb.PostState{
			Stage:       pb.PostState_S_PUB,
			Status:      pb.PostState_S_VISIBLE,
			RevisedAt:   nil,
			ArchivedAt:  nil,
			PublishedAt: timestamppb.Now(),
			CreatedAt:   timestamppb.Now(),
			Reacts:      &pb.Reacts{LikeCount: 103},
		},
	}

	comment := &pb.Comment{
		Id:          utils.RandomUniqueId(),
		User:        owner,
		Uuid:        uuid.NewString(),
		CommentText: &pb.CommentText{Data: "There is a new version for this post!"},
		State: &pb.CommentState{
			CreatedAt: timestamppb.Now(),
			EditedAt:  timestamppb.Now(),
			Reacts:    &pb.Reacts{},
		},
	}

	ustore := store.NewUserStore(store.SqlStore)
	pstore := store.NewPostStore(store.SqlStore)
	cstore := store.NewCommentStore(store.SqlStore)

	if err := ustore.NewUser(owner).Save(); err != nil {
		t.Error(err)
	}

	defer ustore.NewUser(owner).Delete()

	if err := pstore.NewPost(owner, post).Save(); err != nil {
		t.Error(err)
	}

	defer pstore.NewPost(owner, post).Delete()

	if err := cstore.NewComment(comment, post.Uuid, pb.Comment_TT_POST).Save(); err != nil {
		t.Error(err)
	}

	res := S.Exec(context.TODO(), q, "UserOpsForDeleteComment", map[string]any{
		"uuid": comment.Uuid,
	})

	if len(res.Errors) > 0 {
		t.Errorf("%+v", res.Errors)
	}

	body := make(map[string]any)

	if err := json.Unmarshal(res.Data, &body); err != nil {
		t.Fatal(err)
	}

	if C, err := cstore.GetByUuid(comment.Uuid); err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Error(err)
	} else if C != nil {
		t.Errorf("unable to delete comment (%s)", comment.Uuid)
	}

}

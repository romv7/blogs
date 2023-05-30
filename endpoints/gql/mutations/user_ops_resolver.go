package mutations

import (
	"context"
	"fmt"
	"net/http"
	"time"

	_errors "errors"

	"github.com/google/uuid"
	"github.com/romv7/blogs/endpoints/gql/models"
	"github.com/romv7/blogs/errors"
	"github.com/romv7/blogs/internal/pb"
	"github.com/romv7/blogs/internal/store"
	"github.com/romv7/blogs/internal/utils"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type UserMutations_Resolver struct{}

func NewUserMutations_Resolver() *UserMutations_Resolver {
	return &UserMutations_Resolver{}
}

type userOpsResolver = *models.GQLModel_UserOpsResultsResolver

func (r *UserMutations_Resolver) CreatePost(ctx context.Context, args *ArgsCreatePost) (userOpsResolver, error) {
	return r.UpdatePost(ctx, &ArgsUpdatePost{Input: args.Input})
}

func (r *UserMutations_Resolver) UpdatePost(ctx context.Context, args *ArgsUpdatePost) (userOpsResolver, error) {

	res := &models.UserOpsResults{Op: models.UserOpsType_UO_UPDATE_POST, StartTime: time.Now()}
	defer func() {
		res.EndTime = time.Now()
	}()

	ustore := store.NewUserStore(store.SqlStore)
	pstore := store.NewPostStore(store.SqlStore)

	var err error
	var post *store.Post
	var isNewPost bool

	// Query the database for any existing post related to the uuid.
	if args.Input.Uuid != nil {
		post, err = pstore.GetByUuid(*args.Input.Uuid)
	}

	if err != nil && !_errors.Is(err, gorm.ErrRecordNotFound) {
		message := err.Error()
		res.Message = &message
		res.Code = http.StatusBadRequest

		return models.NewGQLModel_UserOpsResultsResolver(res), err
	}

	owner, err := ustore.GetById(uint64(args.Input.Id))
	if err != nil {
		message := err.Error()
		res.Message = &message
		res.Code = http.StatusBadRequest

		return models.NewGQLModel_UserOpsResultsResolver(res), nil
	}

	user := owner.Proto()

	// Only executed when the post is nil (indicating that this post is new)
	if post == nil {
		isNewPost = true
		res.Op = models.UserOpsType_UO_CREATE_POST
		post = pstore.NewPost(user, &pb.Post{
			Id:           utils.RandomUniqueId(),
			Uuid:         uuid.NewString(),
			HeadlineText: "",
			SummaryText:  "",
			Content:      "",
			User:         user,
			Comments:     make([]*pb.Comment, 0),
			Tags:         &pb.Tags{},
			Attachments:  make([]string, 0),
			Refs:         make([]string, 0),
			State: &pb.PostState{
				Reacts:      &pb.Reacts{LoveCount: 154},
				CreatedAt:   timestamppb.Now(),
				RevisedAt:   nil,
				ArchivedAt:  nil,
				PublishedAt: nil,
				Stage:       pb.PostState_S_WIP,
				Status:      pb.PostState_S_DRAFT,
			},
		})
	}

	pbpost := post.Proto()

	// Check if the matching post is owned by the current user request.
	if !isNewPost && user.Id != pbpost.User.Id {
		err := fmt.Errorf("user (%d) is not allowed to update this post", user.Id)
		// TODO: If someone is attempting to update a post by someone.
		//       do something to stop this guy.
		message := err.Error()
		res.Message = &message
		res.Code = http.StatusForbidden

		return models.NewGQLModel_UserOpsResultsResolver(res), err
	}

	if !isNewPost {
		pbpost.State.Stage = pb.PostState_S_REV
		pbpost.State.RevisedAt = timestamppb.Now()
	}

	if args.Input.HeadlineText != nil {
		pbpost.HeadlineText = *args.Input.HeadlineText
	}

	if args.Input.SummaryText != nil {
		pbpost.SummaryText = *args.Input.SummaryText
	}

	if args.Input.Tags != nil {
		pbpost.Tags = &pb.Tags{Data: *args.Input.Tags}
	}

	if args.Input.Stage != nil {
		pbpost.State.Stage = pb.PostState_Stage(*args.Input.Stage)

		switch pbpost.State.Stage {
		case pb.PostState_S_PUB:
			pbpost.State.PublishedAt = timestamppb.Now()
		}
	}

	if args.Input.Status != nil {
		pbpost.State.Status = pb.PostState_Status(*args.Input.Status)

		switch pbpost.State.Status {
		case pb.PostState_S_ARCHIVED:
			pbpost.State.ArchivedAt = timestamppb.Now()
		}
	}

	if args.Input.Content != nil {
		pbpost.Content = *args.Input.Content
	}

	if args.Input.Refs != nil {
		pbpost.Refs = *args.Input.Refs
	}

	if args.Input.Attachments != nil {
		pbpost.Attachments = *args.Input.Attachments
	}

	P := pstore.NewPost(user, pbpost)

	if !isNewPost {
		P.ToggleUpdateMode()
	}

	// Propagate the upate to the database.
	if err = P.Save(); err != nil {
		message := err.Error()
		res.Message = &message
		res.Code = http.StatusBadRequest

		return models.NewGQLModel_UserOpsResultsResolver(res), err
	}

	var message string

	if isNewPost {
		message = fmt.Sprintf("post created (%s)", pbpost.Uuid)
		res.Code = http.StatusAccepted
	} else {
		message = fmt.Sprintf("post was updated (%s)", pbpost.Uuid)
		res.Code = http.StatusOK
	}

	res.Message = &message
	res.Uuid = pbpost.Uuid

	return models.NewGQLModel_UserOpsResultsResolver(res), nil
}

func (r *UserMutations_Resolver) DeletePost(ctx context.Context, args *ArgsDeletePost) (userOpsResolver, error) {
	res := &models.UserOpsResults{Op: models.UserOpsType_UO_DELETE_POST, StartTime: time.Now()}
	defer func() {
		res.EndTime = time.Now()
	}()

	if args.Uuid == "" {
		message := errors.ErrInsufficientArguments.Error()
		res.Message = &message
		res.Code = http.StatusBadRequest

		return models.NewGQLModel_UserOpsResultsResolver(res), errors.ErrInsufficientArguments
	}

	pstore := store.NewPostStore(store.SqlStore)

	var delErr error

	if P, err := pstore.GetByUuid(args.Uuid); err != nil {
		message := err.Error()
		res.Message = &message
		res.Code = http.StatusBadRequest

		return models.NewGQLModel_UserOpsResultsResolver(res), errors.ErrInsufficientArguments
	} else {
		delErr = P.Delete()
	}

	if delErr == nil {
		message := fmt.Sprintf("deleted post (%s)", args.Uuid)
		res.Message = &message
		res.Code = http.StatusNoContent
	} else {
		message := delErr.Error()
		res.Message = &message
		res.Code = http.StatusBadRequest
	}

	return models.NewGQLModel_UserOpsResultsResolver(res), delErr
}

func (r *UserMutations_Resolver) CreateComment(ctx context.Context, args *ArgsCreateComment) (userOpsResolver, error) {
	res := &models.UserOpsResults{Op: models.UserOpsType_UO_CREATE_COMMENT, StartTime: time.Now()}
	defer func() {
		res.EndTime = time.Now()
	}()

	if args.Input.Id == 0 ||
		args.Input.CommentText == "" ||
		args.Input.TargetUuid == "" ||
		args.Input.TargetType == "" {

		message := errors.ErrInsufficientArguments.Error()
		res.Message = &message
		res.Code = http.StatusBadRequest

		return models.NewGQLModel_UserOpsResultsResolver(res), errors.ErrInsufficientArguments
	}

	owner, err := store.NewUserStore(store.SqlStore).GetById(uint64(args.Input.Id))
	if err != nil {
		message := err.Error()
		res.Message = &message
		res.Code = http.StatusBadRequest

		return models.NewGQLModel_UserOpsResultsResolver(res), err
	}

	cstore := store.NewCommentStore(store.SqlStore)
	pbcomment := &pb.Comment{
		Id:          utils.RandomUniqueId(),
		Uuid:        uuid.NewString(),
		User:        owner.Proto(),
		CommentText: &pb.CommentText{Data: args.Input.CommentText},
		Replies:     make([]*pb.Comment, 0),
		State: &pb.CommentState{
			CreatedAt: timestamppb.Now(),
			EditedAt:  timestamppb.Now(),
			Reacts:    &pb.Reacts{},
		},
	}

	C := cstore.NewComment(
		pbcomment,
		args.Input.TargetUuid,
		pb.Comment_TargetType(pb.Comment_TargetType_value[args.Input.TargetType]),
	)

	if err := C.Save(); err != nil {
		message := err.Error()
		res.Message = &message
		res.Code = http.StatusBadRequest

		return models.NewGQLModel_UserOpsResultsResolver(res), err
	}

	message := fmt.Sprintf(`created new comment for target (%s)`, args.Input.TargetUuid)
	res.Message = &message
	res.Code = http.StatusCreated
	res.Uuid = pbcomment.Uuid

	return models.NewGQLModel_UserOpsResultsResolver(res), nil
}

func (r *UserMutations_Resolver) UpdateComment(ctx context.Context, args *ArgsUpdateComment) (userOpsResolver, error) {
	res := &models.UserOpsResults{Op: models.UserOpsType_UO_UPDATE_COMMENT, StartTime: time.Now()}
	defer func() {
		res.EndTime = time.Now()
	}()

	if args.Input.Uuid == nil ||
		args.Input.TargetUuid == "" ||
		args.Input.TargetType == "" {

		message := errors.ErrInsufficientArguments.Error()
		res.Message = &message
		res.Code = http.StatusBadRequest

		return models.NewGQLModel_UserOpsResultsResolver(res), errors.ErrInsufficientArguments
	}

	if len(args.Input.CommentText) != 0 {
		cstore := store.NewCommentStore(store.SqlStore)

		C, err := cstore.GetByUuid(*args.Input.Uuid)
		if err != nil {
			message := err.Error()
			res.Message = &message
			res.Code = http.StatusBadRequest

			return models.NewGQLModel_UserOpsResultsResolver(res), err
		}

		pbc := C.Proto()

		pbc.CommentText = &pb.CommentText{Data: args.Input.CommentText}

		err = cstore.NewComment(pbc, args.Input.TargetUuid, pbc.TargetType).Save()
		if err != nil {
			message := err.Error()
			res.Message = &message
			res.Code = http.StatusBadRequest
		}

		message := fmt.Sprintf("updated comment (%s)", *args.Input.Uuid)
		res.Code = http.StatusOK
		res.Message = &message
		res.Uuid = C.Proto().Uuid

	} else {
		res.Code = http.StatusNoContent
	}

	return models.NewGQLModel_UserOpsResultsResolver(res), nil
}

func (r *UserMutations_Resolver) DeleteComment(ctx context.Context, args *ArgsDeleteComment) (userOpsResolver, error) {

	res := &models.UserOpsResults{Op: models.UserOpsType_UO_DELETE_COMMENT, StartTime: time.Now()}
	defer func() {
		res.EndTime = time.Now()
	}()

	if args.Uuid == "" {
		message := errors.ErrInsufficientArguments.Error()
		res.Message = &message
		res.Code = http.StatusBadRequest

		return models.NewGQLModel_UserOpsResultsResolver(res), errors.ErrInsufficientArguments
	}

	cstore := store.NewCommentStore(store.SqlStore)

	var opErr error

	if C, err := cstore.GetByUuid(args.Uuid); err != nil {
		message := err.Error()
		res.Message = &message
		res.Code = http.StatusBadRequest

		return models.NewGQLModel_UserOpsResultsResolver(res), errors.ErrInsufficientArguments
	} else {
		opErr = C.Delete()
	}

	if opErr == nil {
		message := fmt.Sprintf("deleted comment (%s)", args.Uuid)
		res.Message = &message
		res.Code = http.StatusNoContent
	} else {
		message := opErr.Error()
		res.Message = &message
		res.Code = http.StatusBadRequest
	}

	return models.NewGQLModel_UserOpsResultsResolver(res), opErr
}

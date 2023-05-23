package mutations

import (
	"context"
	"net/http"
	"time"

	"github.com/romv7/blogs/endpoints/gql/models"
	"github.com/romv7/blogs/errors"
)

type UserMutations_Resolver struct{}

func NewUserMutations_Resolver() *UserMutations_Resolver {
	return &UserMutations_Resolver{}
}

type userOpsResolver = *models.GQLModel_UserOpsResultsResolver

func (r *UserMutations_Resolver) CreatePost(ctx context.Context, args *ArgsCreatePost) (userOpsResolver, error) {

	res := &models.UserOpsResults{Op: models.UserOpsType_UO_CREATE_POST, StartTime: time.Now()}
	defer func() {
		res.EndTime = time.Now()
	}()

	r.UpdatePost(ctx, &ArgsUpdatePost{Input: args.Input})

	return models.NewGQLModel_UserOpsResultsResolver(res), nil
}

func (r *UserMutations_Resolver) UpdatePost(ctx context.Context, args *ArgsUpdatePost) (userOpsResolver, error) {

	res := &models.UserOpsResults{Op: models.UserOpsType_UO_UPDATE_POST, StartTime: time.Now()}
	defer func() {
		res.EndTime = time.Now()
	}()

	return models.NewGQLModel_UserOpsResultsResolver(res), nil
}

func (r *UserMutations_Resolver) DeletePost(ctx context.Context, args *ArgsDeletePost) (userOpsResolver, error) {

	res := &models.UserOpsResults{Op: models.UserOpsType_UO_DELETE_POST, StartTime: time.Now()}
	defer func() {
		res.EndTime = time.Now()
	}()

	return models.NewGQLModel_UserOpsResultsResolver(res), nil
}

func (r *UserMutations_Resolver) CreateComment(ctx context.Context, args *ArgsCreateComment) (userOpsResolver, error) {

	res := &models.UserOpsResults{Op: models.UserOpsType_UO_CREATE_COMMENT, StartTime: time.Now()}
	defer func() {
		res.EndTime = time.Now()
	}()

	if args.Input.CommentText == "" ||
		args.Input.TargetUuid == "" ||
		args.Input.TargetType == "" {

		message := errors.ErrInsufficientArguments.Error()
		res.Message = &message
		res.Code = http.StatusBadRequest

		return models.NewGQLModel_UserOpsResultsResolver(res), errors.ErrInsufficientArguments
	}

	return models.NewGQLModel_UserOpsResultsResolver(res), nil
}

func (r *UserMutations_Resolver) UpdateComment(ctx context.Context, args *ArgsUpdateComment) (userOpsResolver, error) {

	res := &models.UserOpsResults{Op: models.UserOpsType_UO_UPDATE_COMMENT, StartTime: time.Now()}
	defer func() {
		res.EndTime = time.Now()
	}()

	if args.Input.CommentText == "" ||
		args.Input.TargetUuid == "" ||
		args.Input.TargetType == "" {

		message := errors.ErrInsufficientArguments.Error()
		res.Message = &message
		res.Code = http.StatusBadRequest

		return models.NewGQLModel_UserOpsResultsResolver(res), errors.ErrInsufficientArguments
	}

	return models.NewGQLModel_UserOpsResultsResolver(res), nil
}

func (r *UserMutations_Resolver) DeleteComment(ctx context.Context, args *ArgsDeleteComment) (userOpsResolver, error) {

	res := &models.UserOpsResults{Op: models.UserOpsType_UO_DELETE_COMMENT, StartTime: time.Now()}
	defer func() {
		res.EndTime = time.Now()
	}()

	return models.NewGQLModel_UserOpsResultsResolver(res), nil
}

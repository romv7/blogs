package mutations

import (
	"context"

	"github.com/romv7/blogs/endpoints/gql/models"
)

type UserMutations_Resolver struct{}

func NewUserMutations_Resolver() *UserMutations_Resolver {
	return &UserMutations_Resolver{}
}

func (r *UserMutations_Resolver) CreatePost(ctx context.Context, args *ArgsCreatePost) *models.GQLModel_UserOpsResultsResolver {

	return &models.GQLModel_UserOpsResultsResolver{}
}

func (r *UserMutations_Resolver) UpdatePost(ctx context.Context, args *ArgsUpdatePost) *models.GQLModel_UserOpsResultsResolver {
	return &models.GQLModel_UserOpsResultsResolver{}
}

func (r *UserMutations_Resolver) DeletePost(ctx context.Context, args *ArgsDeletePost) *models.GQLModel_UserOpsResultsResolver {
	return &models.GQLModel_UserOpsResultsResolver{}
}

func (r *UserMutations_Resolver) CreateComment(ctx context.Context, args *ArgsCreateComment) *models.GQLModel_UserOpsResultsResolver {
	return &models.GQLModel_UserOpsResultsResolver{}
}

func (r *UserMutations_Resolver) UpdateComment(ctx context.Context, args *ArgsUpdateComment) *models.GQLModel_UserOpsResultsResolver {
	return &models.GQLModel_UserOpsResultsResolver{}
}

func (r *UserMutations_Resolver) DeleteComment(ctx context.Context, args *ArgsDeleteComment) *models.GQLModel_UserOpsResultsResolver {
	return &models.GQLModel_UserOpsResultsResolver{}
}

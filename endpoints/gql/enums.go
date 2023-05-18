package gql

import (
	gql "github.com/graphql-go/graphql"
	"github.com/romv7/blogs/internal/pb"
)

const (
	EnumPostStageWorkInProgress = "S_WIP"
	EnumPostStageReview         = "S_REVIEW"
	EnumPostStagePublished      = "S_PUB"
	EnumPostStageRevision       = "S_REV"

	EnumPostStatusDraft    = "S_DRAFT"
	EnumPostStatusArchived = "S_ARCHIVED"
	EnumPostStatusTrash    = "S_TRASH"
	EnumPostStatusVisible  = "S_VISIBLE"
	EnumPostStatusHidden   = "S_HIDDEN"

	EnumCommentTargetTypeUser    = "TT_USER"
	EnumCommentTargetTypeComment = "TT_COMMENT"
	EnumCommentTargetTypePost    = "TT_POST"

	EnumUserTypeAuthor = "T_AUTHOR"
	EnumUserTypeNormal = "T_NORMAL"
)

var (
	PostStageEnum = gql.NewEnum(gql.EnumConfig{
		Name:        "PostStage",
		Description: "",
		Values: gql.EnumValueConfigMap{

			// S_WIP
			EnumPostStageWorkInProgress: &gql.EnumValueConfig{
				Description: "",
				Value:       pb.PostState_S_WIP,
			},

			// S_REVIEW
			EnumPostStageReview: &gql.EnumValueConfig{
				Description: "",
				Value:       pb.PostState_S_REVIEW,
			},

			// S_PUB
			EnumPostStagePublished: &gql.EnumValueConfig{
				Description: "",
				Value:       pb.PostState_S_PUB,
			},

			// S_REV
			EnumPostStageRevision: &gql.EnumValueConfig{
				Description: "",
				Value:       pb.PostState_S_REV,
			},
		},
	})

	PostStatusEnum = gql.NewEnum(gql.EnumConfig{
		Name:        "PostStatus",
		Description: "",
		Values: gql.EnumValueConfigMap{

			// S_DRAFT
			EnumPostStatusDraft: &gql.EnumValueConfig{
				Description: "",
				Value:       pb.PostState_S_DRAFT,
			},

			// S_ARCHIVED
			EnumPostStatusArchived: &gql.EnumValueConfig{
				Description: "",
				Value:       pb.PostState_S_ARCHIVED,
			},

			// S_TRASH
			EnumPostStatusTrash: &gql.EnumValueConfig{
				Description: "",
				Value:       pb.PostState_S_TRASH,
			},

			// S_VISIBLE
			EnumPostStatusVisible: &gql.EnumValueConfig{
				Description: "",
				Value:       pb.PostState_S_VISIBLE,
			},

			// S_HIDDEN
			EnumPostStatusHidden: &gql.EnumValueConfig{
				Description: "",
				Value:       pb.PostState_S_HIDDEN,
			},
		},
	})

	CommentTargetTypeEnum = gql.NewEnum(gql.EnumConfig{
		Name:        "CommentTargetType",
		Description: "",
		Values: gql.EnumValueConfigMap{

			// TT_USER
			EnumCommentTargetTypeUser: &gql.EnumValueConfig{},

			// TT_COMMENT
			EnumCommentTargetTypeComment: &gql.EnumValueConfig{},

			// TT_POST
			EnumCommentTargetTypePost: &gql.EnumValueConfig{},
		},
	})

	UserTypeEnum = gql.NewEnum(gql.EnumConfig{
		Name:        "UserTypeEnum",
		Description: "",
		Values: gql.EnumValueConfigMap{

			// T_AUTHOR
			EnumUserTypeAuthor: &gql.EnumValueConfig{
				Description: "",
				Value:       pb.User_T_AUTHOR,
			},

			// T_NORMAL
			EnumUserTypeNormal: &gql.EnumValueConfig{
				Description: "",
				Value:       pb.User_T_NORMAL,
			},
		},
	})
)

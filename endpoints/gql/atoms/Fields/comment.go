package Fields

import (
	gql "github.com/graphql-go/graphql"
	"github.com/romv7/blogs/endpoints/gql/atoms/Enums"
)

var (
	FieldCommentID = &gql.Field{
		Name: "id",
		Type: gql.NewNonNull(gql.ID),
		Resolve: func(P gql.ResolveParams) (out any, err error) {

			return
		},
	}

	FieldCommentUUID = &gql.Field{
		Name: "uuid",
		Type: gql.NewNonNull(gql.String),
		Resolve: func(P gql.ResolveParams) (out any, err error) {

			return
		},
	}

	FieldCommentUser = &gql.Field{
		Name: "user",
		Type: gql.NewNonNull(Enums.ENUM_USER_TYPE),
		Resolve: func(P gql.ResolveParams) (out any, err error) {

			return
		},
	}

	FieldCommentCommentText = &gql.Field{
		Name: "commentText",
		Type: gql.NewNonNull(gql.String),
		Resolve: func(P gql.ResolveParams) (out any, err error) {

			return
		},
	}

	FieldCommentTargetType = &gql.Field{
		Name: "targetType",
		Type: gql.NewNonNull(Enums.ENUM_COMMENT_TARGET_TYPE),
		Resolve: func(P gql.ResolveParams) (out any, err error) {

			return
		},
	}
)

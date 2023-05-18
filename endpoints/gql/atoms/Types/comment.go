package Types

import (
	gql "github.com/graphql-go/graphql"
	"github.com/romv7/blogs/endpoints/gql/atoms/Enums"
	"github.com/romv7/blogs/endpoints/gql/atoms/Fields"
	"github.com/romv7/blogs/internal/utils/gqlutils"
)

var (
	CommentTargetUnion = gql.NewUnion(gql.UnionConfig{
		Name:  "CommentTargetUnion",
		Types: []*gql.Object{PostObject, UserObject, CommentObject},
	})

	CommentObject = gql.NewObject(gql.ObjectConfig{
		Name: "User",
		Fields: gql.Fields{
			"id":          Fields.FieldCommentID,
			"uuid":        Fields.FieldCommentUUID,
			"commentText": Fields.FieldCommentCommentText,
		},
	})

	Comment_MiscFields = gql.Fields{
		"user": &gql.Field{
			Type: gql.NewNonNull(UserObject),
			Resolve: func(P gql.ResolveParams) (out any, err error) {

				return
			},
		},
		"replies": &gql.Field{
			Type: gql.NewNonNull(
				gql.NewList(
					gql.NewNonNull(CommentObject),
				),
			),
			Resolve: func(P gql.ResolveParams) (out any, err error) {

				return
			},
		},
		"targetType": &gql.Field{
			Type: gql.NewNonNull(Enums.ENUM_COMMENT_TARGET_TYPE),
			Resolve: func(P gql.ResolveParams) (out any, err error) {

				return
			},
		},
		"target": &gql.Field{
			Type: gql.NewNonNull(CommentTargetUnion),
			Resolve: func(P gql.ResolveParams) (out any, err error) {

				return
			},
		},
	}
)

func init() {
	gqlutils.AddFields(CommentObject, Comment_MiscFields)
}

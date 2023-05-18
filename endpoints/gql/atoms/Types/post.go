package Types

import (
	gql "github.com/graphql-go/graphql"
	"github.com/romv7/blogs/endpoints/gql/atoms/Fields"
	"github.com/romv7/blogs/internal/utils/gqlutils"
)

var (
	PostObject = gql.NewObject(gql.ObjectConfig{
		Name: "Post",
		Fields: gql.Fields{
			"id":           Fields.FieldPostID,
			"uuid":         Fields.FieldPostUUID,
			"headlineText": Fields.FieldPostHeadlineText,
			"summaryText":  Fields.FieldPostSummaryText,
			"tags":         Fields.FieldPostTags,
			"attachments":  Fields.FieldPostAttachments,
			"refs":         Fields.FieldPostRefs,
			"content":      Fields.FieldPostContent,
			"stage":        Fields.FieldPostStage,
			"status":       Fields.FieldPostStatus,
			"createdAt":    Fields.FieldPostCreatedAt,
			"revisedAt":    Fields.FieldPostRevisedAt,
			"archivedAt":   Fields.FieldPostArchivedAt,
			"publishedAt":  Fields.FieldPostPublishedAt,
		},
	})

	Post_MiscFields = gql.Fields{
		"user": &gql.Field{
			Type: gql.NewNonNull(UserObject),
			Resolve: func(P gql.ResolveParams) (out any, err error) {

				return
			},
		},
		"comments": &gql.Field{
			Type: gql.NewNonNull(
				gql.NewList(
					gql.NewNonNull(CommentObject),
				),
			),
			Resolve: func(P gql.ResolveParams) (out any, err error) {

				return
			},
		},
		"reacts": &gql.Field{
			Type: gql.NewNonNull(ReactsObject),
			Resolve: func(P gql.ResolveParams) (out any, err error) {

				return
			},
		},
	}
)

func init() {
	gqlutils.AddFields(PostObject, Post_MiscFields)
}

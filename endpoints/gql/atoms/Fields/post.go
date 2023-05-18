package Fields

import (
	gql "github.com/graphql-go/graphql"
	"github.com/romv7/blogs/endpoints/gql/atoms/Enums"
)

var (
	FieldPostID = &gql.Field{
		Name: "id",
		Type: gql.NewNonNull(gql.ID),
		Resolve: func(P gql.ResolveParams) (out any, err error) {

			return
		},
	}

	FieldPostUUID = &gql.Field{
		Name: "uuid",
		Type: gql.NewNonNull(gql.String),
		Resolve: func(P gql.ResolveParams) (out any, err error) {

			return
		},
	}

	FieldPostHeadlineText = &gql.Field{
		Name: "headlineText",
		Type: gql.NewNonNull(gql.String),
		Resolve: func(P gql.ResolveParams) (out any, err error) {

			return
		},
	}

	FieldPostSummaryText = &gql.Field{
		Name: "summaryText",
		Type: gql.NewNonNull(gql.String),
		Resolve: func(P gql.ResolveParams) (out any, err error) {

			return
		},
	}

	FieldPostTags = &gql.Field{
		Name: "tags",
		Type: gql.NewNonNull(gql.NewList(gql.String)),
		Resolve: func(P gql.ResolveParams) (out any, err error) {

			return
		},
	}

	FieldPostAttachments = &gql.Field{
		Name: "attachments",
		Type: gql.NewNonNull(gql.NewList(gql.String)),
		Resolve: func(P gql.ResolveParams) (out any, err error) {

			return
		},
	}

	FieldPostRefs = &gql.Field{
		Name: "refs",
		Type: gql.NewNonNull(gql.NewList(gql.String)),
		Resolve: func(P gql.ResolveParams) (out any, err error) {

			return
		},
	}

	FieldPostContent = &gql.Field{
		Name: "content",
		Type: gql.NewNonNull(gql.String),
		Resolve: func(P gql.ResolveParams) (out any, err error) {

			return
		},
	}

	FieldPostStage = &gql.Field{
		Name: "stage",
		Type: Enums.ENUM_POST_STAGE,
		Resolve: func(P gql.ResolveParams) (out any, err error) {

			return
		},
	}

	FieldPostStatus = &gql.Field{
		Name: "status",
		Type: Enums.ENUM_POST_STATUS,
		Resolve: func(P gql.ResolveParams) (out any, err error) {

			return
		},
	}

	FieldPostCreatedAt = &gql.Field{
		Name: "createdAt",
		Type: gql.NewNonNull(gql.DateTime),
		Resolve: func(P gql.ResolveParams) (out any, err error) {

			return
		},
	}

	FieldPostRevisedAt = &gql.Field{
		Name: "revisedAt",
		Type: gql.NewNonNull(gql.DateTime),
		Resolve: func(P gql.ResolveParams) (out any, err error) {

			return
		},
	}

	FieldPostArchivedAt = &gql.Field{
		Name: "archivedAt",
		Type: gql.DateTime,
		Resolve: func(P gql.ResolveParams) (out any, err error) {

			return
		},
	}

	FieldPostPublishedAt = &gql.Field{
		Name: "publishedAt",
		Type: gql.DateTime,
		Resolve: func(P gql.ResolveParams) (out any, err error) {

			return
		},
	}
)

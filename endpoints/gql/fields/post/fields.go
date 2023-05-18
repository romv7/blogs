package post

import (
	gql "github.com/graphql-go/graphql"
	"github.com/romv7/blogs/endpoints/gql/resolvers"
)

var (
	Fields = gql.Fields{
		"id":           Post_Id,
		"uuid":         Post_Uuid,
		"headlineText": Post_HeadlineText,
		"summaryText":  Post_SummaryText,
		"user":         Post_User,
		"comments":     Post_Comments,
		"tags":         Post_Tags,
		"images":       Post_Images,
		"attachments":  Post_Attachments,
		"refs":         Post_Refs,
		"stage":        Post_Stage,
		"status":       Post_Status,
		"reacts":       Post_Reacts,
		"createdAt":    Post_CreatedAt,
		"revisedAt":    Post_RevisedAt,
		"archivedAt":   Post_ArchivedAt,
		"publishedAt":  Post_PublishedAt,
	}

	Post_Id = &gql.Field{
		Name:        "id",
		Description: "",
		Resolve:     resolvers.PostResolveId,
	}

	Post_Uuid = &gql.Field{
		Name:        "uuid",
		Description: "",
		Resolve:     resolvers.PostResolveUuid,
	}

	Post_HeadlineText = &gql.Field{
		Name:        "headlineText",
		Description: "",
		Resolve:     resolvers.PostResolveHeadlineText,
	}

	Post_SummaryText = &gql.Field{
		Name:        "summaryText",
		Description: "",
		Resolve:     resolvers.PostResolveSummaryText,
	}

	Post_User = &gql.Field{
		Name:        "user",
		Description: "",
		Resolve:     resolvers.PostResolveUser,
	}

	Post_Comments = &gql.Field{
		Name:        "comments",
		Description: "",
		Resolve:     resolvers.PostResolveComments,
	}

	Post_Tags = &gql.Field{
		Name:        "tags",
		Description: "",
		Resolve:     resolvers.PostResolveTags,
	}

	Post_Images = &gql.Field{
		Name:        "images",
		Description: "",
		Resolve:     resolvers.PostResolveImages,
	}

	Post_Attachments = &gql.Field{
		Name:        "attachments",
		Description: "",
		Resolve:     resolvers.PostResolveAttachments,
	}

	Post_Refs = &gql.Field{
		Name:        "refs",
		Description: "",
		Resolve:     resolvers.PostResolveRefs,
	}

	Post_Stage = &gql.Field{
		Name:        "stage",
		Description: "",
		Resolve:     resolvers.PostResolveStage,
	}

	Post_Status = &gql.Field{
		Name:        "status",
		Description: "",
		Resolve:     resolvers.PostResolveStatus,
	}

	Post_Reacts = &gql.Field{
		Name:        "reacts",
		Description: "",
		Resolve:     resolvers.PostResolveReacts,
	}

	Post_CreatedAt = &gql.Field{
		Name:        "createdAt",
		Description: "",
		Resolve:     resolvers.PostResolveCreatedAt,
	}

	Post_RevisedAt = &gql.Field{
		Name:        "revisedAt",
		Description: "",
		Resolve:     resolvers.PostResolveRevisedAt,
	}

	Post_ArchivedAt = &gql.Field{
		Name:        "archivedAt",
		Description: "",
		Resolve:     resolvers.PostResolveArchivedAt,
	}

	Post_PublishedAt = &gql.Field{
		Name:        "publishedAt",
		Description: "",
		Resolve:     resolvers.PostResolvePublishedAt,
	}
)

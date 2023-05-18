package comment

import (
	gql "github.com/graphql-go/graphql"
	"github.com/romv7/blogs/endpoints/gql/resolvers"
)

var (
	Fields = gql.Fields{
		"id":          Comment_Id,
		"uuid":        Comment_Uuid,
		"user":        Comment_User,
		"commentText": Comment_Text,
		"replies":     Comment_Replies,
		"reacts":      Comment_Reacts,
		"createdAt":   Comment_CreatedAt,
		"editedAt":    Comment_EditedAt,
	}

	Comment_Id = &gql.Field{
		Name:        "id",
		Description: "",
		Resolve:     resolvers.CommentResolveId,
	}

	Comment_Uuid = &gql.Field{
		Name:        "uuid",
		Description: "",
		Resolve:     resolvers.CommentResolveUuid,
	}

	Comment_User = &gql.Field{
		Name:        "user",
		Description: "",
		Resolve:     resolvers.CommentResolveUser,
	}

	Comment_Text = &gql.Field{
		Name:        "commentText",
		Description: "",
		Resolve:     resolvers.CommentResolveCommentText,
	}

	Comment_Replies = &gql.Field{
		Name:        "replies",
		Description: "",
		Resolve:     resolvers.CommentResolveReplies,
	}

	Comment_Reacts = &gql.Field{
		Name:        "reacts",
		Description: "",
		Resolve:     resolvers.CommentResolveReacts,
	}

	Comment_CreatedAt = &gql.Field{
		Name:        "createdAt",
		Description: "",
		Resolve:     resolvers.CommentResolveCreatedAt,
	}

	Comment_EditedAt = &gql.Field{
		Name:        "editedAt",
		Description: "",
		Resolve:     resolvers.CommentResolveEditedAt,
	}
)

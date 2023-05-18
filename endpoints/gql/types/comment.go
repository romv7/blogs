package types

import (
	gql "github.com/graphql-go/graphql"
)

var (
	CommentObjectType = gql.NewObject(gql.ObjectConfig{
		Name:        "Comment",
		Description: "",
		Fields:      gql.Fields{},
	})
)

package types

import (
	gql "github.com/graphql-go/graphql"
	"github.com/romv7/blogs/endpoints/gql/fields/post"
)

var (
	PostObjectType = gql.NewObject(gql.ObjectConfig{
		Name:        "Post",
		Description: "",
		Fields:      post.Fields,
	})
)

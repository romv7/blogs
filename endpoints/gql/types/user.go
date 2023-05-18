package types

import (
	gql "github.com/graphql-go/graphql"
	"github.com/romv7/blogs/endpoints/gql/fields/user"
)

var (
	UserObjectType = gql.NewObject(gql.ObjectConfig{
		Name:        "User",
		Description: "",
		Fields:      user.Fields,
	})
)

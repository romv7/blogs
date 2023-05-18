package Types

import (
	gql "github.com/graphql-go/graphql"
	"github.com/romv7/blogs/internal/pb"
)

var (
	RootQueryFields = gql.Fields{
		"UserQuery": &gql.Field{
			Type: UserQueryObject,
			Resolve: func(P gql.ResolveParams) (out any, error error) {
				out = P

				return
			},
		},
	}

	UserQueryObject = gql.NewObject(gql.ObjectConfig{
		Name: "UserQuery",
		Fields: gql.Fields{
			"user": &gql.Field{
				Name: "user",
				Type: gql.NewNonNull(UserObject),
				Resolve: func(P gql.ResolveParams) (out any, err error) {
					out = &pb.User{}

					return
				},
			},
		},
	})
)

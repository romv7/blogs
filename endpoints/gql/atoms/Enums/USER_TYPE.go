package Enums

import (
	gql "github.com/graphql-go/graphql"
	"github.com/romv7/blogs/internal/pb"
)

var ENUM_USER_TYPE = gql.NewEnum(gql.EnumConfig{
	Name: "UserType",
	Values: gql.EnumValueConfigMap{
		"T_NORMAL": &gql.EnumValueConfig{
			Value: pb.User_T_NORMAL,
		},

		"T_AUTHOR": &gql.EnumValueConfig{
			Value: pb.User_T_AUTHOR,
		},
	},
})

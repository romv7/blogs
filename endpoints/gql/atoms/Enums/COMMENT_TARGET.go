package Enums

import (
	gql "github.com/graphql-go/graphql"
	"github.com/romv7/blogs/internal/pb"
)

var ENUM_COMMENT_TARGET_TYPE = gql.NewEnum(gql.EnumConfig{
	Name: "PostStatus",
	Values: gql.EnumValueConfigMap{

		"TT_USER": &gql.EnumValueConfig{
			Value: pb.Comment_TT_USER,
		},

		"TT_COMMENT": &gql.EnumValueConfig{
			Value: pb.Comment_TT_COMMENT,
		},

		"TT_POST": &gql.EnumValueConfig{
			Value: pb.Comment_TT_POST,
		},
	},
})

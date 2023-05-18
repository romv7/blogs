package Enums

import (
	gql "github.com/graphql-go/graphql"
	"github.com/romv7/blogs/internal/pb"
)

var ENUM_POST_STAGE = gql.NewEnum(gql.EnumConfig{
	Name: "PostStatus",
	Values: gql.EnumValueConfigMap{

		"S_WIP": &gql.EnumValueConfig{
			Value: pb.PostState_S_WIP,
		},

		"S_REVIEW": &gql.EnumValueConfig{
			Value: pb.PostState_S_REVIEW,
		},

		"S_PUB": &gql.EnumValueConfig{
			Value: pb.PostState_S_PUB,
		},

		"S_REV": &gql.EnumValueConfig{
			Value: pb.PostState_S_REV,
		},
	},
})

package Enums

import (
	gql "github.com/graphql-go/graphql"
	"github.com/romv7/blogs/internal/pb"
)

var ENUM_POST_STATUS = gql.NewEnum(gql.EnumConfig{
	Name: "PostStatus",
	Values: gql.EnumValueConfigMap{

		"S_DRAFT": &gql.EnumValueConfig{
			Value: pb.PostState_S_DRAFT,
		},

		"S_ARCHIVED": &gql.EnumValueConfig{
			Value: pb.PostState_S_ARCHIVED,
		},

		"S_TRASH": &gql.EnumValueConfig{
			Value: pb.PostState_S_TRASH,
		},

		"S_VISIBLE": &gql.EnumValueConfig{
			Value: pb.PostState_S_VISIBLE,
		},

		"S_HIDDEN": &gql.EnumValueConfig{
			Value: pb.PostState_S_HIDDEN,
		},
	},
})

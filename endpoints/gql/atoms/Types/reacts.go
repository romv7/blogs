package Types

import (
	gql "github.com/graphql-go/graphql"
	"github.com/romv7/blogs/endpoints/gql/atoms/Fields"
)

var (
	ReactsObject = gql.NewObject(gql.ObjectConfig{
		Name: "Reacts",
		Fields: gql.Fields{
			"likeCount":     Fields.FieldReactsLikeCount,
			"confusedCount": Fields.FieldReactsConfusedCount,
			"loveCount":     Fields.FieldReactsLoveCount,
			"laughCount":    Fields.FieldReactsLaughCount,
			"sadCount":      Fields.FieldReactsSadCount,
			"careCount":     Fields.FieldReactsCareCount,
		},
	})
)

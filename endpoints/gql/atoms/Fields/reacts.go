package Fields

import gql "github.com/graphql-go/graphql"

var (
	FieldReactsLikeCount = &gql.Field{
		Name: "likeCount",
		Type: gql.Int,
	}

	FieldReactsConfusedCount = &gql.Field{
		Name: "confusedCount",
		Type: gql.Int,
	}

	FieldReactsLoveCount = &gql.Field{
		Name: "loveCount",
		Type: gql.Int,
	}

	FieldReactsLaughCount = &gql.Field{
		Name: "laughCount",
		Type: gql.Int,
	}

	FieldReactsSadCount = &gql.Field{
		Name: "sadCount",
		Type: gql.Int,
	}

	FieldReactsCareCount = &gql.Field{
		Name: "careCount",
		Type: gql.Int,
	}
)

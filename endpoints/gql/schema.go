package gql

import (
	"github.com/graph-gophers/graphql-go"
	"github.com/romv7/blogs/internal/utils/gqlutils"
)

func DefaultSchema() *graphql.Schema {
	return graphql.MustParseSchema(gqlutils.JoinAllSchemaFiles(), NewRootQuery())
}

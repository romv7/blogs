package atoms

import (
	"log"

	gql "github.com/graphql-go/graphql"
	"github.com/romv7/blogs/endpoints/gql/atoms/Types"
)

var (
	DefaultQueryObject = gql.NewObject(gql.ObjectConfig{
		Name:   "RootQuery",
		Fields: Types.RootQueryFields,
	})
)

func init() {

}

func DefaultSchema() (out gql.Schema) {
	var err error

	out, err = gql.NewSchema(gql.SchemaConfig{Query: DefaultQueryObject})
	if err != nil {
		log.Panic(err)
	}

	return
}

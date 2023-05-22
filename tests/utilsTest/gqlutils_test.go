package utilsTest

import (
	"testing"

	"github.com/graph-gophers/graphql-go"
	"github.com/romv7/blogs/internal/utils/gqlutils"
)

func TestJoinAllSchemaFiles(t *testing.T) {
	schemas := gqlutils.JoinAllSchemaFiles()

	if _, err := graphql.ParseSchema(schemas, nil); err != nil {
		t.Error(err)
	}
}

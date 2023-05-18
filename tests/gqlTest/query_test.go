package gqlTest

import (
	"encoding/json"
	"fmt"
	"testing"

	gql "github.com/graphql-go/graphql"
	gqlEndpoint "github.com/romv7/blogs/endpoints/gql"
)

func TestDefaultSchema(t *testing.T) {
	S := gqlEndpoint.DefaultSchema()

	params := gql.Params{Schema: S, RequestString: "{ UserQuery { user { id, name, fullName, email } } }"}
	if r := gql.Do(params); len(r.Errors) > 0 {
		t.Errorf("%+v", r.Errors)
	} else {
		body, _ := json.Marshal(r.Data)
		fmt.Println(string(body))
	}
}

package Types

import (
	gql "github.com/graphql-go/graphql"
	"github.com/romv7/blogs/endpoints/gql/atoms/Fields"
)

var (
	UserObject = gql.NewObject(gql.ObjectConfig{
		Name: "User",
		Fields: gql.Fields{
			"id":         Fields.FieldUserID,
			"uuid":       Fields.FieldUserUUID,
			"name":       Fields.FieldUserName,
			"fullName":   Fields.FieldUserFullName,
			"email":      Fields.FieldUserEmail,
			"type":       Fields.FieldUserType,
			"createdAt":  Fields.FieldUserCreatedAt,
			"updatedAt":  Fields.FieldUserUpdatedAt,
			"isDisabled": Fields.FieldUserIsDisabled,
			"isVerified": Fields.FieldUserIsVerified,
		},
	})
)

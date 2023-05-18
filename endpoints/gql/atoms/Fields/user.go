package Fields

import (
	"time"

	"github.com/google/uuid"
	gql "github.com/graphql-go/graphql"
	"github.com/romv7/blogs/endpoints/gql/atoms/Enums"
	"github.com/romv7/blogs/internal/pb"
	"github.com/romv7/blogs/internal/utils"
)

var (
	FieldUserID = &gql.Field{
		Name: "id",
		Type: gql.NewNonNull(gql.ID),
		Resolve: func(P gql.ResolveParams) (out any, err error) {
			out = utils.RandomUniqueId()

			return
		},
	}

	FieldUserUUID = &gql.Field{
		Name: "uuid",
		Type: gql.NewNonNull(gql.String),
		Resolve: func(P gql.ResolveParams) (out any, err error) {
			out = uuid.NewString()

			return
		},
	}

	FieldUserName = &gql.Field{
		Name: "name",
		Type: gql.NewNonNull(gql.String),
		Resolve: func(P gql.ResolveParams) (out any, err error) {
			out = "rommms"

			return
		},
	}

	FieldUserFullName = &gql.Field{
		Name: "fullName",
		Type: gql.NewNonNull(gql.String),
		Resolve: func(P gql.ResolveParams) (out any, err error) {
			out = "Rom Vales Villanueva"

			return
		},
	}

	FieldUserEmail = &gql.Field{
		Name: "email",
		Type: gql.NewNonNull(gql.String),
		Resolve: func(P gql.ResolveParams) (out any, err error) {
			out = "romdevmod@gmail.com"

			return
		},
	}

	FieldUserType = &gql.Field{
		Name: "type",
		Type: Enums.ENUM_USER_TYPE,
		Resolve: func(P gql.ResolveParams) (out any, err error) {
			out = pb.User_T_AUTHOR

			return
		},
	}

	FieldUserCreatedAt = &gql.Field{
		Name: "createdAt",
		Type: gql.NewNonNull(gql.DateTime),
		Resolve: func(P gql.ResolveParams) (out any, err error) {
			out = time.Now().UTC()

			return
		},
	}

	FieldUserUpdatedAt = &gql.Field{
		Name: "updatedAt",
		Type: gql.DateTime,
		Resolve: func(P gql.ResolveParams) (out any, err error) {
			out = time.Now().UTC()

			return
		},
	}

	FieldUserIsDisabled = &gql.Field{
		Name: "isDisabled",
		Type: gql.Boolean,
		Resolve: func(P gql.ResolveParams) (out any, err error) {
			out = false

			return
		},
	}

	FieldUserIsVerified = &gql.Field{
		Name: "isVerified",
		Type: gql.Boolean,
		Resolve: func(P gql.ResolveParams) (out any, err error) {
			out = true

			return
		},
	}
)

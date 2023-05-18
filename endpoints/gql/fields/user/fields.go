package user

import (
	gql "github.com/graphql-go/graphql"
	"github.com/romv7/blogs/endpoints/gql/resolvers"
)

var (
	Fields = gql.Fields{
		"id":         User_Id,
		"uuid":       User_Uuid,
		"name":       User_Name,
		"fullName":   User_FullName,
		"email":      User_Email,
		"type":       User_Type,
		"createdAt":  User_CreatedAt,
		"updatedAt":  User_UpdatedAt,
		"isDisabled": User_IsDisabled,
		"isVerified": User_IsVerified,
	}

	User_Id = &gql.Field{
		Name:        "id",
		Description: "",
		Resolve:     resolvers.UserResolveId,
	}

	User_Uuid = &gql.Field{
		Name:        "uuid",
		Description: "",
		Resolve:     resolvers.UserResolveUuid,
	}

	User_Name = &gql.Field{
		Name:        "name",
		Description: "",
		Resolve:     resolvers.UserResolveName,
	}

	User_FullName = &gql.Field{
		Name:        "fullName",
		Description: "",
		Resolve:     resolvers.UserResolveFullName,
	}

	User_Email = &gql.Field{
		Name:        "email",
		Description: "",
		Resolve:     resolvers.UserResolveEmail,
	}

	User_Type = &gql.Field{
		Name:        "type",
		Description: "",
		Resolve:     resolvers.UserResolveType,
	}

	User_CreatedAt = &gql.Field{
		Name:        "createdAt",
		Description: "",
		Resolve:     resolvers.UserResolveCreatedAt,
	}

	User_UpdatedAt = &gql.Field{
		Name:        "updatedAt",
		Description: "",
		Resolve:     resolvers.UserResolveUpdatedAt,
	}

	User_IsDisabled = &gql.Field{
		Name:        "isDisabled",
		Description: "",
		Resolve:     resolvers.UserResolveIsDisabled,
	}

	User_IsVerified = &gql.Field{
		Name:        "isVerified",
		Description: "",
		Resolve:     resolvers.UserResolveIsVerified,
	}
)

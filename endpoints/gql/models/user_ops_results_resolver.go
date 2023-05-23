package models

import "github.com/graph-gophers/graphql-go"

type GQLModel_UserOpsResultsResolver struct {
	res *UserOpsResults
}

func NewGQLModel_UserOpsResultsResolver(so *UserOpsResults) *GQLModel_UserOpsResultsResolver {
	return &GQLModel_UserOpsResultsResolver{so}
}

func (so *GQLModel_UserOpsResultsResolver) Op() string {
	return so.res.Op.String()
}

func (so *GQLModel_UserOpsResultsResolver) Message() *string {
	return so.res.Message
}

func (so *GQLModel_UserOpsResultsResolver) Code() *int32 {
	return &so.res.Code
}

func (so *GQLModel_UserOpsResultsResolver) Uuid() string {
	return so.res.Uuid
}

func (so *GQLModel_UserOpsResultsResolver) StartTime() graphql.Time {
	return graphql.Time{Time: so.res.StartTime}
}

func (so *GQLModel_UserOpsResultsResolver) EndTime() graphql.Time {
	return graphql.Time{Time: so.res.EndTime}
}

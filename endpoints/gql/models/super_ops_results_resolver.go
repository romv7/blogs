package models

import "github.com/graph-gophers/graphql-go"

type GQLModel_SuperOpsResultsResolver struct {
	res *SuperOpsResults
}

func NewGQLModel_SuperOpsResultsResolver(so *SuperOpsResults) *GQLModel_SuperOpsResultsResolver {
	return &GQLModel_SuperOpsResultsResolver{so}
}

func (so *GQLModel_SuperOpsResultsResolver) Op() string {
	return so.res.Op.String()
}

func (so *GQLModel_SuperOpsResultsResolver) Message() *string {
	return so.res.Message
}

func (so *GQLModel_SuperOpsResultsResolver) Code() *int32 {
	return &so.res.Code
}

func (so *GQLModel_SuperOpsResultsResolver) Uuid() string {
	return so.res.Uuid
}

func (so *GQLModel_SuperOpsResultsResolver) StartTime() graphql.Time {
	return graphql.Time{Time: so.res.StartTime}
}

func (so *GQLModel_SuperOpsResultsResolver) EndTime() graphql.Time {
	return graphql.Time{Time: so.res.EndTime}
}

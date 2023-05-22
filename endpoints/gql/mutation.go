package gql

import "github.com/romv7/blogs/endpoints/gql/mutations"

func (RootQuery) SuperOps() *mutations.SuperMutations_Resolver {
	return mutations.NewSuperMutations_Resolver()
}

func (RootQuery) UserOps() *mutations.UserMutations_Resolver {
	return mutations.NewUserMutations_Resolver()
}

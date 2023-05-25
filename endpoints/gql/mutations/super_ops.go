package mutations

type UserSuperMutationsParameter struct {
	Id                     *float64
	Name, FullName, Email  *string
	IsDisabled, IsVerified *bool
	Type                   *string

	// Author related metadatas
	Bio, AltName *string
}

type ArgsAddUser struct {
	Input *UserSuperMutationsParameter
}

type ArgsUpdateUser struct {
	Input *UserSuperMutationsParameter
}

type ArgsDeleteUser struct {
	Uuid string
}

// TODO: Make a more meaningful message response. SuperOps

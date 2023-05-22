package mutations

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/romv7/blogs/endpoints/gql/models"
	"github.com/romv7/blogs/errors"
	"github.com/romv7/blogs/internal/pb"
	"github.com/romv7/blogs/internal/store"
	"github.com/romv7/blogs/internal/utils"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type SuperMutations_Resolver struct{}

func NewSuperMutations_Resolver() *SuperMutations_Resolver {
	return &SuperMutations_Resolver{}
}

type outResolver = *models.GQLModel_SuperOpsResultsResolver

func (SuperMutations_Resolver) AddUser(ctx context.Context, args *ArgsAddUser) (outResolver, error) {
	var err error

	res := &models.SuperOpsResults{Op: models.SuperOpsType_SO_NEW, StartTime: time.Now()}
	defer func() {
		res.EndTime = time.Now()
	}()

	if args.Input == nil ||
		args.Input.FullName == nil ||
		args.Input.Name == nil ||
		args.Input.Email == nil {

		res.Code = http.StatusBadRequest

		message := errors.ErrInsufficientArguments.Error()
		res.Message = &message

		return models.NewGQLModel_SuperOpsResultsResolver(res), errors.ErrInsufficientArguments
	}

	ustore := store.NewUserStore(store.SqlStore)
	pbuser := &pb.User{
		Id:       utils.RandomUniqueId(),
		Uuid:     uuid.NewString(),
		Name:     *args.Input.Name,
		FullName: *args.Input.FullName,
		Email:    *args.Input.Email,
		State: &pb.UserState{
			CreatedAt: timestamppb.New(res.StartTime),
			UpdatedAt: timestamppb.New(res.StartTime),
			Disabled:  false,
			UVerified: false,
		},
	}

	if args.Input.Type != nil {
		pbuser.Type = pb.User_Type(pb.User_Type_value[*args.Input.Type])
	} else {
		pbuser.Type = pb.User_T_NORMAL
	}

	if err = ustore.NewUser(pbuser).Save(); err != nil {
		res.Code = http.StatusBadRequest
		message := err.Error()
		res.Message = &message

		return models.NewGQLModel_SuperOpsResultsResolver(res), err
	}

	message := fmt.Sprintf("added user \"%s\" with an id of %d.", pbuser.Name, pbuser.Id)
	res.Message = &message
	res.Code = http.StatusCreated
	res.Uuid = pbuser.Uuid

	return models.NewGQLModel_SuperOpsResultsResolver(res), nil
}

func (SuperMutations_Resolver) UpdateUser(ctx context.Context, args *ArgsUpdateUser) (outResolver, error) {
	var err error

	res := &models.SuperOpsResults{Op: models.SuperOpsType_SO_UPDATE, StartTime: time.Now().UTC()}
	defer func() {
		res.EndTime = time.Now().UTC()
	}()

	if args.Input.Id == nil {
		message := errors.ErrInsufficientArguments.Error()
		res.Message = &message
		res.Code = http.StatusBadRequest

		return models.NewGQLModel_SuperOpsResultsResolver(res), errors.ErrInsufficientArguments
	}

	ustore := store.NewUserStore(store.SqlStore)
	var u *store.User

	if u, err = ustore.GetById(uint64(*args.Input.Id)); err != nil {
		message := err.Error()
		res.Message = &message
		res.Code = http.StatusBadRequest

		return models.NewGQLModel_SuperOpsResultsResolver(res), err
	}

	isUpdated := false
	pbuser := u.Proto()

	if args.Input.Email != nil {
		pbuser.Email = *args.Input.Email
		isUpdated = true
	}

	if args.Input.FullName != nil {
		pbuser.FullName = *args.Input.FullName
		isUpdated = true
	}

	if args.Input.Name != nil {
		pbuser.Name = *args.Input.Name
		isUpdated = true
	}

	if args.Input.Type != nil {
		pbuser.Type = pb.User_Type(pb.User_Type_value[*args.Input.Type])
		isUpdated = true
	}

	if args.Input.IsDisabled != nil {
		pbuser.State.Disabled = *args.Input.IsDisabled
		isUpdated = true
	}

	if args.Input.IsVerified != nil {
		pbuser.State.UVerified = *args.Input.IsVerified
		isUpdated = true
	}

	if isUpdated {
		pbuser.State.UpdatedAt = timestamppb.New(time.Now().UTC())
	}

	if pbuser.Type == pb.User_T_AUTHOR {

		if args.Input.Bio != nil {
			pbuser.Bio = *args.Input.Bio
		}

		if args.Input.AltName != nil {
			pbuser.AltName = *args.Input.Name
		}

		// TODO: Add an ability for authors to add their external social platform links

	}

	if err = ustore.NewUser(pbuser).Save(); err != nil {
		message := err.Error()
		res.Message = &message
		res.Code = http.StatusBadRequest
		res.Uuid = pbuser.Uuid

		return models.NewGQLModel_SuperOpsResultsResolver(res), err
	}

	message := fmt.Sprintf(`updated user "%s" with the new fields.`, pbuser.Name)
	res.Message = &message
	res.Code = http.StatusOK
	res.Uuid = pbuser.Uuid

	return models.NewGQLModel_SuperOpsResultsResolver(res), nil
}

func (SuperMutations_Resolver) DeleteUser(ctx context.Context, args *ArgsDeleteUser) (outResolver, error) {
	var err error

	res := &models.SuperOpsResults{Op: models.SuperOpsType_SO_DELETE, StartTime: time.Now().UTC()}
	defer func() {
		res.EndTime = time.Now().UTC()
	}()

	if _, err := uuid.Parse(args.Uuid); err != nil {
		message := errors.ErrInsufficientArguments.Error()
		res.Message = &message
		res.Code = http.StatusBadRequest

		return models.NewGQLModel_SuperOpsResultsResolver(res), err
	}

	var user *store.User
	ustore := store.NewUserStore(store.SqlStore)

	if u, err := ustore.GetByUuid(args.Uuid); err != nil {
		message := err.Error()
		res.Message = &message
		res.Code = http.StatusBadRequest

		return models.NewGQLModel_SuperOpsResultsResolver(res), err
	} else {
		user = u
	}

	if err = user.Delete(); err != nil {
		message := err.Error()
		res.Message = &message
		res.Code = http.StatusBadRequest

		return models.NewGQLModel_SuperOpsResultsResolver(res), err
	}

	message := fmt.Sprintf("deleted the user with a uuid \"%s\"", args.Uuid)
	res.Message = &message

	res.Code = http.StatusNoContent
	res.Uuid = user.Proto().Uuid

	return models.NewGQLModel_SuperOpsResultsResolver(res), nil
}

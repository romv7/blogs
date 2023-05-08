package grpc

import (
	"context"

	"github.com/google/uuid"
	"github.com/romv7/blogs/internal/pb"
	"github.com/romv7/blogs/internal/store"
	"github.com/romv7/blogs/internal/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type U struct {
	pb.UnimplementedUserServiceServer
}

func (svc *U) New(ctx context.Context, params *pb.UserService_New_Params) (*pb.User, error) {
	if len(params.Name) == 0 ||
		len(params.Email) == 0 ||
		len(params.FullName) == 0 {
		return nil, status.Errorf(codes.InvalidArgument, ErrNotEnoughArgument.Error())
	}

	out := &pb.User{Id: utils.RandomUniqueId(), Uuid: uuid.NewString()}
	out.Email = params.Email
	out.FullName = params.FullName
	out.Name = params.Name
	out.Type = params.Type
	out.State = &pb.UserState{
		CreatedAt: timestamppb.Now(),
		UpdatedAt: nil,
		Disabled:  false,
		UVerified: false,
	}

	return out, nil
}

func (svc *U) Save(ctx context.Context, params *pb.UserService_Save_Params) (*pb.UserService_Save_Response, error) {
	if params.User == nil {
		return nil, status.Errorf(codes.InvalidArgument, ErrNotEnoughArgument.Error())
	}

	ustore := store.NewUserStore(store.SqlStore)
	if err := ustore.Save(ustore.NewUser(params.User)); err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &pb.UserService_Save_Response{}, nil
}

func (svc *U) Delete(ctx context.Context, params *pb.UserService_Delete_Params) (*pb.UserService_Delete_Response, error) {
	return &pb.UserService_Delete_Response{}, nil
}

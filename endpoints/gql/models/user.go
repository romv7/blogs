package models

import (
	"time"

	"github.com/romv7/blogs/internal/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type User struct {
	Id         float64      `json:"id"`
	Uuid       string       `json:"uuid"`
	Name       string       `json:"name"`
	FullName   string       `json:"fullName"`
	Email      string       `json:"email"`
	Type       pb.User_Type `json:"type"`
	CreatedAt  time.Time    `json:"createdAt"`
	UpdatedAt  time.Time    `json:"updatedAt"`
	IsDisabled bool         `json:"isDisabled"`
	IsVerified bool         `json:"isVerified"`
}

func Proto_GQLModelUser(u *pb.User) *User {
	return &User{
		Id:         float64(u.Id),
		Uuid:       u.Uuid,
		Name:       u.Name,
		FullName:   u.FullName,
		Email:      u.Email,
		Type:       u.Type,
		CreatedAt:  u.State.CreatedAt.AsTime(),
		UpdatedAt:  u.State.UpdatedAt.AsTime(),
		IsDisabled: u.State.Disabled,
		IsVerified: u.State.UVerified,
	}
}

func (u *User) Proto() *pb.User {
	return &pb.User{
		Id:       uint64(u.Id),
		Uuid:     u.Uuid,
		Name:     u.Name,
		FullName: u.FullName,
		Email:    u.Email,
		Type:     u.Type,
		State: &pb.UserState{
			CreatedAt: timestamppb.New(u.CreatedAt),
			UpdatedAt: timestamppb.New(u.UpdatedAt),
			Disabled:  u.IsDisabled,
			UVerified: u.IsVerified,
		},
	}
}

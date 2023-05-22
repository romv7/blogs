package models

import (
	"time"

	"github.com/romv7/blogs/internal/constants"
	"github.com/romv7/blogs/internal/pb"
	"github.com/romv7/blogs/internal/utils/author"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type User struct {
	ID        uint64       `gorm:"column:id;autoIncrement:false"`
	Uuid      string       `gorm:"column:uuid;unique"`
	Name      string       `gorm:"column:name"`
	FullName  string       `gorm:"column:fname"`
	Email     string       `gorm:"column:email"`
	Type      pb.User_Type `gorm:"column:type"`
	Disabled  bool         `gorm:"column:is_disabled"`
	Verified  bool         `gorm:"column:is_verified"`
	CreatedAt time.Time    `gorm:"column:created_at"`
	UpdatedAt time.Time    `gorm:"column:updated_at"`

	Bio         string                                `gorm:"-"`
	AltName     string                                `gorm:"-"`
	Stats       *author.AuthorStats                   `gorm:"-"`
	SocialLinks map[constants.SocialLinkType][]string `gorm:"-"`
}

func NewUser(u *pb.User) (uout *User) {
	uout = &User{
		ID:        u.Id,
		Uuid:      u.Uuid,
		Name:      u.Name,
		FullName:  u.FullName,
		Email:     u.Email,
		Type:      u.Type,
		Disabled:  u.State.Disabled,
		Verified:  u.State.UVerified,
		CreatedAt: u.State.CreatedAt.AsTime().UTC(),
		UpdatedAt: u.State.UpdatedAt.AsTime().UTC(),
	}

	return
}

func (u *User) Proto() (out *pb.User) {
	out = &pb.User{
		Id:       u.ID,
		Uuid:     u.Uuid,
		Name:     u.Name,
		FullName: u.FullName,
		Email:    u.Email,
		Type:     u.Type,
		State: &pb.UserState{
			CreatedAt: timestamppb.New(u.CreatedAt),
			UpdatedAt: timestamppb.New(u.UpdatedAt),
			Disabled:  u.Disabled,
			UVerified: u.Verified,
		},
	}

	return
}

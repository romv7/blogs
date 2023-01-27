package entities

import (
	"github.com/rommms07/blogs/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type React struct {
	*User
	*pb.React
}

func NewReact(user *User, targetUuid string, typ pb.React_Type) (react *React) {
	react = &React{
		User: user,
		React: &pb.React{
			UserId:     user.User.Id,
			TargetUuid: targetUuid, 
			Type:       typ,
			ReactAt:    timestamppb.Now(),
		},
	}

	return
}

func (r *React) Remove() error {
	return nil
}

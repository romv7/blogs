package entities

import (
	"github.com/rommms07/blogs/pb"
)

type React struct {
	*User
	*pb.React
}

func NewReact(user *User, targetId uint64, targetType pb.React_TargetType, typ pb.React_Type) (react *React) {
	react = &React{
		User: user,
		React: &pb.React{
			TargetId:   targetId,
			TargetType: targetType,
			Type:       typ,
		},
	}

	return
}

func (r *React) Remove() error {
	return nil
}

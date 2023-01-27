package entities

import (
	"github.com/rommms07/blogs/pb"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Post struct {
	*User
	*pb.Post
}

func NewPost(user *User, keywords ...string) (post *Post) {
	post = &Post{
		User: user,
		Post: &pb.Post{
			UserId:    user.User.Id,
			Keywords:  keywords,
			Uuid:      uuid.New().String(),
			State:	   &pb.PostState {
				Stage: pb.PostState_S_WIP,
				Status: pb.PostState_S_DRAFT,
				CreatedAt: timestamppb.Now(),
			},	
		},
	}

	return
}

func (p *Post) SetHeadlineText(htxt string) *Post {
	p.HeadlineText = htxt
	return p
}

func (p *Post) SetSubjectLine(sub string) *Post {
	p.SubjectText = sub
	return p
}

func (p *Post) ToArchive() *Post {
	return p
}

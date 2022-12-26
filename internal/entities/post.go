package entities

import (
	"time"

	"github.com/rommms07/blogs/pb"
)

type Post struct {
	*User
	*pb.Post
}

func NewPost(user *User, keywords ...string) (post *Post) {
	post = &Post{
		User: user,
		Post: &pb.Post{
			User:      user.User,
			Stage:     pb.Post_S_WIP,
			Status:    pb.Post_S_DRAFT,
			Keywords:  keywords,
			CreatedAt: uint64(time.Now().Unix()),
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

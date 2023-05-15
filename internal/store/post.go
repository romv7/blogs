package store

import (
	"log"

	"github.com/romv7/blogs/internal/pb"
	"github.com/romv7/blogs/internal/storage"
	"github.com/romv7/blogs/internal/utils/author"
)

func (p *Post) Save() (err error) {
	pstore := NewPostStore(p.t)

	if err = pstore.Save(p); err != nil {
		return
	}

	if p.Proto().User != nil && p.Proto().User.Type == pb.User_T_AUTHOR {
		ah := author.NewAuthorHelper(p.Proto().User, storage.Plain)

		if err = ah.SaveAuthorPost(p.Proto()); err != nil {
			return author.ErrNormalUserHasNoResourceId
		}
	}

	return
}

func (p *Post) Delete() (err error) {
	pstore := NewPostStore(p.t)

	if p.Proto().User != nil && p.Proto().User.Type == pb.User_T_AUTHOR {
		ah := author.NewAuthorHelper(p.Proto().User, storage.Plain)
		if err = ah.DeletePostMetadata(p.Proto()); err != nil {
			return
		}
	}

	err = pstore.Delete(p)

	return
}

func (p *Post) SetOwner(u *User) {
	switch p.t {
	case SqlStore:
		p.sqlModel.UserId = u.sqlModel.ID
	default:
		log.Panic(ErrInvalidStore)
	}
}

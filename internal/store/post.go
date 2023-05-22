package store

import (
	"errors"
	"log"
	"os"

	"github.com/romv7/blogs/internal/pb"
	"github.com/romv7/blogs/internal/storage"
	"github.com/romv7/blogs/internal/utils/author"
	"gorm.io/gorm"
)

func (p *Post) Proto() *pb.Post {
	ustore := NewUserStore(SqlStore)
	cstore := NewCommentStore(SqlStore)

	var pout *pb.Post

	switch p.t {
	case SqlStore:
		pout = p.sqlModel.Proto()

		if u, err := ustore.GetById(p.sqlModel.UserId); errors.Is(err, gorm.ErrRecordNotFound) {
			// TODO: Handle post that has no owner.
		} else {
			pout.User = u.Proto()

			if pout.User.Type == pb.User_T_AUTHOR {
				ah := author.NewAuthorHelper(pout.User, p.s)

				// TODO: Get post metadata from the p.s storage.
				switch p.s {
				case storage.Plain:
					m, content, err := ah.GetAuthorPostMetadata(pout)

					var Err error

					// The post still isn't save in the storage.
					if err == nil {
						pout.HeadlineText = m.HeadlineText
						pout.SummaryText = m.SummaryText
						pout.Refs = m.References
						pout.Content = content
					} else if os.IsNotExist(err) {
						Err = ah.SaveAuthorPost(pout)
					} else {
						Err = err
					}

					if Err != nil {
						log.Panic(Err)
					}
				default:
					log.Panic(storage.ErrorInvalidStorageDriver)
				}
			}
		}

		pout.Comments = cstore.TargetCommentProtoTree(pout.Uuid)

		return pout
	default:
		log.Panic(ErrInvalidStore)
	}

	return nil
}

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

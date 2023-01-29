package comment

import (
	"fmt"

	"github.com/rommms07/blogs/internal/entities"
	"github.com/rommms07/blogs/internal/store"
	"github.com/rommms07/blogs/internal/store/source/sql"
)

const (
	db_name = "comments"
)

type CommentStoreSql struct {
	store.UnimplementedStore
}

var (
	commentsDb *sql.SQLDataSource

	initSql = func() {
		store.InitSqlDb(db_name, func(db *sql.SQLDataSource) {
			commentsDb = db
		})
	}
)

func NewSQLCommentStore(db *sql.SQLDataSource) *CommentStoreSql {
	if commentsDb == nil && db != nil {
		commentsDb = db
	}

	return &CommentStoreSql{}
}

func (s *CommentStoreSql) Save(comment *entities.Comment) (err error) {
	initSql()

	tbl := commentsDb.GetDestTableByUnix(db_name, comment.State.CreatedAt.AsTime().Unix())
	query := fmt.Sprintf("insert into %s (uuid, user_uuid, comment_text, replies, edited, created_at, edited_at, target_uuid) values (?, ?, ?, ?, ?, ?, ?, ?) on duplicate key update uuid=values(uuid), user_uuid=values(uuid), comment_text=values(comment_text), replies=values(replies), created_at=values(created_at), edited_at=values(edited_at), target_uuid=values(target_uuid);", tbl)

	_, err = commentsDb.Query(query, 
		comment.Uuid,
		comment.UserUuid,
		comment.CommentText,
		len(comment.Replies),
		comment.State.Edited,
		comment.State.CreatedAt.AsTime(),
		comment.State.EditedAt.AsTime(),
		comment.TargetUuid,
	)


	return
}

func (s *CommentStoreSql) DeleteByUuid(uuid uint64) (err error) {
	initSql()

	return
}

func (s *CommentStoreSql) Read(query string, args ...any) (err error) {
	initSql()

	return
}

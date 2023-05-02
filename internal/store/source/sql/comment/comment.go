package comment

import (
	"fmt"

	"github.com/romv7/blogs/internal/entities"
	"github.com/romv7/blogs/internal/store"
	"github.com/romv7/blogs/internal/store/source/sql"
	"github.com/romv7/blogs/internal/store/source/sql/uuidindexes"
)

const (
	db_name = "comments"
)

type CommentStoreSql struct {
	T int64
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
	query := fmt.Sprintf(`
	insert into %s 
		(
			id,
		 	uuid, 
		 	user_id, 
		 	comment_text, 
		 	edited, 
		 	created_at, 
		 	edited_at, 
		 	target_type,
		 	target_uuid
		) values (?, ?, ?, ?, ?, ?, ?, ?, ?) 
	on duplicate key update 
		uuid=values(uuid), 
		user_id=values(user_id), 
		comment_text=values(comment_text),
		created_at=values(created_at), 
		edited_at=values(edited_at), 
		target_type=values(target_type),
		target_uuid=values(target_uuid);`, tbl)

	uindex := uuidindexes.NewUuidIndex(comment.Id, db_name, tbl, comment.Uuid, comment.UniqueKey())

	if len(comment.GetTargetUuid()) == 0 {
		return fmt.Errorf("comment has no target")
	}

	_, err = commentsDb.Query(query,
		comment.Id,
		comment.Uuid,
		comment.User.Id,
		comment.CommentText,
		comment.State.Edited,
		comment.State.CreatedAt.AsTime(),
		comment.State.EditedAt.AsTime(),
		comment.TargetType,
		comment.GetTargetUuid(),
	)

	if err != nil {
		return err
	}

	return uindex.Save()
}

func (s *CommentStoreSql) Delete(c *entities.Comment) (err error) {

	return
}

// TODO: DeleteByUuid
func (s *CommentStoreSql) DeleteByUuid(uuid uint64) (err error) {
	initSql()

	return
}

func DbName() string {
	return db_name
}

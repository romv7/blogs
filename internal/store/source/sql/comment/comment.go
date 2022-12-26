package comment

import (
	"github.com/rommms07/blogs/internal/entities"
	"github.com/rommms07/blogs/internal/store"
	"github.com/rommms07/blogs/internal/store/source/sql"
)

const (
	db_name = "commentsDb"
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

	commentsDb.Exec("insert into `comments0`")

	return
}

func (s *CommentStoreSql) Delete(id uint64) (err error) {
	initSql()

	return
}

func (s *CommentStoreSql) Read(query string, args ...any) (err error) {
	initSql()

	return
}

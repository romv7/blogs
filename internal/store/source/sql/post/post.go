package post

import (
	"github.com/rommms07/blogs/internal/entities"
	"github.com/rommms07/blogs/internal/store"
	"github.com/rommms07/blogs/internal/store/source/sql"
)

var postsDb *sql.SQLDataSource

const (
	db_name = "posts"
)

type PostStoreSql struct {
	store.UnimplementedStore
}

var initSql = func() {
	store.InitSqlDb(db_name, func(db *sql.SQLDataSource) {
		postsDb = db
	})
}

func (s *PostStoreSql) Save(post *entities.Post) (err error) {
	return
}

func (s *PostStoreSql) Delete(id uint64) (err error) {
	return
}

func (s *PostStoreSql) Read(query string, args ...any) (err error) {
	return
}

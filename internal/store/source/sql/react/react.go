package react

import (
	"github.com/rommms07/blogs/internal/entities"
	"github.com/rommms07/blogs/internal/store"
	"github.com/rommms07/blogs/internal/store/source/sql"
)

var reactsDb *sql.SQLDataSource

const (
	db_name = "reacts"
)

type ReactStoreSql struct {
	store.UnimplementedStore
}

var initSql = func() {
	store.InitSqlDb(db_name, func(db *sql.SQLDataSource) {
		reactsDb = db
	})
}

func (s *ReactStoreSql) Save(react *entities.React) (err error) {
	initSql()

	return
}

func (s *ReactStoreSql) Delete(id uint64) (err error) {
	initSql()

	return
}

func (s *ReactStoreSql) Read(query string, args ...any) (err error) {
	initSql()

	return
}

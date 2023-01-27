package user

import (
	"github.com/rommms07/blogs/internal/entities"
	"github.com/rommms07/blogs/internal/store"
	"github.com/rommms07/blogs/internal/store/source/sql"
)

var usersDb *sql.SQLDataSource

const (
	db_name = "users"
)

type UserStoreSql struct {
	store.UnimplementedStore
}

var initSql = func() {
	println(123)
	store.InitSqlDb(db_name, func(db *sql.SQLDataSource) {
		usersDb = db
	})
}

func (s *UserStoreSql) Save(user *entities.User) (err error) {
	initSql()

	N := usersDb.GetTableIdByUnix(user.State.CreatedAt.AsTime().Unix())
	println(N)

	return
}

func (s *UserStoreSql) Delete(id uint64) (err error) {
	initSql()

	return
}

func (s *UserStoreSql) Read(query string, args ...any) (err error) {
	initSql()

	return
}

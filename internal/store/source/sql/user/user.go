package user

import (
	"fmt"

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
	store.InitSqlDb(db_name, func(db *sql.SQLDataSource) {
		usersDb = db
	})
}

// Save
func (s *UserStoreSql) Save(user *entities.User) (err error) {
	initSql()

	// only works in MariaDB
	tbl := usersDb.GetDestTableByUnix(db_name, user.State.CreatedAt.AsTime().Unix())
		query := fmt.Sprintf(`insert into %s (name, full_name, email, type, uuid, picture_url, created_at, updated_at, disabled) values (?, ?, ?, ?, ?, ?, ?, ?, ?) on duplicate key update name=values(name), full_name=values(full_name), email=values(email), uuid=values(uuid), type=values(type), picture_url=values(picture_url), created_at=values(created_at), updated_at=values(updated_at), disabled=values(disabled);`, tbl)

	_, err = usersDb.Query(query,
		user.Name,
		user.FullName,
		user.Email,
		user.Type,
		user.Uuid,
		user.State.PictureUrl,
		user.State.CreatedAt.AsTime(),
		user.State.UpdatedAt.AsTime(),
		user.State.Disabled,
	)

	return
}

// DeleteById
func (s *UserStoreSql) DeleteByUuid(uuid string, t int64) (err error) {
	initSql()

	tbl := usersDb.GetDestTableByUnix(db_name, t)

	_, err = usersDb.Exec(fmt.Sprintf("delete low_priority from %s where uuid=?", tbl), uuid)

	return
}

// Read
func (s *UserStoreSql) Read(query string, args ...any) (err error) {
	initSql()

	return
}

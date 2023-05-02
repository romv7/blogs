package user

import (
	"fmt"
	"time"

	"github.com/romv7/blogs/internal/entities"
	"github.com/romv7/blogs/internal/store"
	"github.com/romv7/blogs/internal/store/source/sql"
	"github.com/romv7/blogs/internal/store/source/sql/uuidindexes"
	"github.com/romv7/blogs/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var usersDb *sql.SQLDataSource

const (
	db_name = "users"
)

type UserStoreSql struct {
	T int64
	store.UnimplementedStore
}

var initSql = func() {
	store.InitSqlDb(db_name, func(db *sql.SQLDataSource) {
		usersDb = db
	})
}

// A method for saving an entities.User to a database. If the argument record
// passed to the method already exists in the database, it will return an error.
func (s *UserStoreSql) Save(user *entities.User) (err error) {
	initSql()

	// only works in MariaDB
	tbl := usersDb.GetDestTableByUnix(db_name, s.T)
	uindex := uuidindexes.NewUuidIndex(
		user.Id,
		db_name,
		tbl,
		user.Uuid,
		user.Email)

	if uindex.Exists() {
		return
	}

	query := fmt.Sprintf(`
	insert into %s 
		(id, 
		 name, 
		 full_name, 
		 email, 
		 type, 
		 uuid, 
		 created_at, 
		 updated_at, 
		 disabled) values (?, ?, ?, ?, ?, ?, ?, ?, ?)
	on duplicate key update 
		name=values(name), 
		full_name=values(full_name), 
		email=values(email), uuid=values(uuid), 
		type=values(type), created_at=values(created_at), 
		updated_at=values(updated_at), 
		disabled=values(disabled);`, tbl)

	_, err = usersDb.Query(
		query,
		user.Id,
		user.Name,
		user.FullName,
		user.Email,
		user.Type,
		user.Uuid,
		user.State.CreatedAt.AsTime(),
		user.State.UpdatedAt.AsTime(),
		user.State.Disabled,
	)

	if err != nil {
		return
	}

	return uindex.Save()
}

// If you want to delete an entities.User from the database you can call
// this method and it will delete its record from the database. If there
// is an error it will return it.
func (s *UserStoreSql) Delete(u *entities.User) (err error) {
	initSql()

	ui, err := uuidindexes.GetUuidIndex(db_name, u.UniqueKey())
	if err != nil {
		return err
	}

	query := fmt.Sprintf(`delete low_priority from %s where id=? or uuid=? or email=?;`, ui.Ref)
	if _, err := usersDb.Query(query, u.Id, u.Uuid, u.Email); err != nil {
		return err
	}

	err = ui.Revoke()

	return
}

// TODO: DeleteByUuid
func (s *UserStoreSql) DeleteByUuid(uuid string) (err error) {
	initSql()
	tbl := usersDb.GetDestTableByUnix(db_name, s.T)
	_, err = usersDb.Exec(fmt.Sprintf("delete low_priority from %s where uuid=?;", tbl), uuid)
	return
}

// TODO: DeleteById
func (s *UserStoreSql) DeleteById(id uint64) (err error) {
	initSql()
	tbl := usersDb.GetDestTableByUnix(db_name, s.T)
	_, err = usersDb.Exec(fmt.Sprintf("delete low_priority from %s where id=?;", tbl), id)
	return
}

// If you have a uuid of an entites.User you can invoke this method to get
// the record of pointing to the uuid argument.
func (s *UserStoreSql) GetByUuid(uuid string) (u *entities.User, err error) {
	u = &entities.User{
		User: &pb.User{
			State: &pb.UserState{},
		},
	}

	ui, err := uuidindexes.GetUuidIndexByUuid(db_name, uuid)
	if err != nil {
		return nil, err
	}

	query := fmt.Sprintf(`select * from %s where uuid=?;`, ui.Ref)
	row := usersDb.QueryRow(query, uuid)

	var createdAt, updatedAt time.Time

	err = row.Scan(
		&u.Id,
		&u.Uuid,
		&u.Name,
		&u.FullName,
		&u.Email,
		&u.Type,
		&createdAt,
		&updatedAt,
		&u.State.Disabled,
	)

	u.State.CreatedAt = timestamppb.New(createdAt)
	u.State.UpdatedAt = timestamppb.New(updatedAt)

	if err != nil {
		return nil, err
	}

	return
}

func DbName() string {
	return db_name
}

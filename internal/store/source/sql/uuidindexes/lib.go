package uuidindexes

import (
	"encoding/base64"
	"fmt"
	"log"
	"time"

	"github.com/romv7/blogs/internal/store"
	"github.com/romv7/blogs/internal/store/source/sql"
	"github.com/romv7/blogs/internal/utils/rand"

	_sql "database/sql"
)

var (
	uuidIndexesDb *sql.SQLDataSource

	db_name = "uuidIndexes"

	initSql = func() {
		store.InitSqlDb(db_name, func(db *sql.SQLDataSource) {
			uuidIndexesDb = db
		})
	}
)

type UuidIndex struct {
	Id           uint32
	Resource_id  uint32
	Resource_key string
	Uuid         string
	Ref          string
	Pem          byte
	Created_at   time.Time
}

// Creates a new UuidIndex instance by using the arguments passed to this function.
func NewUuidIndex(id uint32, dbName, tbl, uuid, key string) *UuidIndex {
	now := time.Now()
	N, _ := rand.Rand()

	conf, err := store.GetDBConnConfig(dbName)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return &UuidIndex{
		Id:           uint32(now.Unix()) + uint32(N),
		Uuid:         uuid,
		Resource_id:  id,
		Resource_key: fmt.Sprintf("%s:%s", dbName, base64.StdEncoding.EncodeToString([]byte(key))),
		Ref:          fmt.Sprintf("`%s`.%s", conf.DBName, tbl),
		Pem:          0,
		Created_at:   now,
	}
}

// If you are looking for an entity but don't know its uuid you can look it up
// by using this function by passing a dbName and a key as its arguments,
// it will magically out its result, if there is an error err will not be nil.
func GetUuidIndex(dbName, key string) (result *UuidIndex, err error) {
	result, err = getUuidIndexWrapper(
		`select * from %s where resource_key=?;`,
		fmt.Sprintf("%s:%s", dbName, base64.StdEncoding.EncodeToString([]byte(key))))

	return
}

// On the other hand you can use this function if you have a uuid but don't have
// a slight idea of the key of an entity. This is used in the production a lot
// compared to the `GetUuidIndex`
func GetUuidIndexByUuid(dbName, uuid string) (result *UuidIndex, err error) {
	result, err = getUuidIndexWrapper(`select * from %s where uuid=?;`, uuid)
	return
}

func getUuidIndexWrapper(query string, params ...any) (result *UuidIndex, err error) {
	result = &UuidIndex{}

	// TODO: Instead of looping through each table partitions, try making its query mechanism
	// 			 parallel by wrapping this block in goroutines.
	for part := int(store.GetPartCount(db_name) - 1); part >= 0; part-- {
		row := uuidIndexesDb.QueryRow(fmt.Sprintf(query, fmt.Sprintf("%s%d", db_name, part)), params...)

		err = row.Scan(
			&result.Id,
			&result.Resource_id,
			&result.Resource_key,
			&result.Uuid,
			&result.Ref,
			&result.Pem,
			&result.Created_at,
		)

		if err == _sql.ErrNoRows {
			continue
		}

		err = nil
		break
	}

	if err != nil {
		result = nil
	}

	return
}

// Commit a uuidindex to the database. Take note that this does not ensure
// that the saved uuidindex is unique in the database hence make sure that
// what you passed in the database is unique.
func (u *UuidIndex) Save() (err error) {
	initSql()

	N, _ := rand.Rand()
	tbl := uuidIndexesDb.GetDestTableByUnix(db_name, u.Created_at.Unix()+int64(N))
	query := fmt.Sprintf(`
	insert into %s
		(id,
		 resource_id,
		 resource_key,
		 uuid,
		 ref,
		 pem,
		 created_at) values (?, ?, ?, ?, ?, ?, ?);
	`, tbl)

	if u.Exists() {
		return
	}

	_, err = uuidIndexesDb.Query(query,
		u.Id,
		u.Resource_id,
		u.Resource_key,
		u.Uuid,
		u.Ref,
		u.Pem,
		u.Created_at,
	)

	return
}

// Call this function often if you are deleting an entity from the database.
// What this does is it revokes a uuidindex stored in the database by deleting it.
// Consider this as a simple cleanup function for unused object.
func (u *UuidIndex) Revoke() (err error) {
	initSql()

	for part := int(store.GetPartCount(db_name) - 1); part >= 0; part-- {
		query := fmt.Sprintf(`delete low_priority from %s where resource_key=?;`, fmt.Sprintf("%s%d", db_name, part))
		uuidIndexesDb.Query(query, u.Resource_key)
	}

	return nil
}

// To ensure that a new entity does not exist in the database, call this function.
// This lookup the `uuidIndexes` tables for a `resource_key` of a uuidindex object.
func (u *UuidIndex) Exists() bool {
	initSql()

	var res bool

	// TODO: Instead of looping through each table partitions, try making its checking mechanism
	// 			 parallel by wrapping this block in goroutines.
	for part := int(store.GetPartCount(db_name) - 1); part >= 0; part-- {
		query := fmt.Sprintf(`select * from %s where resource_key=?;`, fmt.Sprintf("%s%d", db_name, part))

		row := uuidIndexesDb.QueryRow(query, u.Resource_key)

		if row.Scan(nil) == _sql.ErrNoRows {
			res = false
		} else {
			res = true
			break
		}

	}

	return res
}

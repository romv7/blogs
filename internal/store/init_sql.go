package store

import (
	"log"

	"github.com/romv7/blogs/internal/store/source/sql"
)

var InitSqlDb = func(dbName string, initCb func(*sql.SQLDataSource)) {
	ds, err := OpenSqlDb(dbName, nil)
	if err != nil {
		log.Fatalf(err.Error())
	}

	db, err := ds.Source.(*sql.SQLDataSource).Connect()
	if err != nil || db.DB == nil {
		log.Fatalf(err.Error())
	}

	initCb(db)
}

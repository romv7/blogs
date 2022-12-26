package store

import (
	"log"

	"github.com/rommms07/blogs/internal/store/source/sql"
)

var InitSqlDb = func(connUrlName string, initCb func(*sql.SQLDataSource)) {
	ds, err := OpenSqlDb(connUrlName, nil)
	if err != nil {
		log.Fatalf(err.Error())
	}

	db, err := ds.Source.(*sql.SQLDataSource).Connect()
	if err != nil || db.DB == nil {
		log.Fatalf(err.Error())
	}

	initCb(db)
}

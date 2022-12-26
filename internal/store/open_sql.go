package store

import (
	"log"

	"github.com/rommms07/blogs/internal"
	"github.com/rommms07/blogs/internal/store/source/sql"
)

var ds_err error

func OpenSqlDb(connUrlName string, initCb func(*DataSource[sql.SQLDataSource])) (*DataSource[sql.SQLDataSource], error) {
	config, err := internal.LoadConfig()
	if err != nil {
		return nil, err
	}

	drvName := config.Database.Drv_name

	return NewDataSource(func(ds *DataSource[sql.SQLDataSource]) {
		var prefix string

		if config.IsDev() {
			prefix = config.Main.Environ + "_"
		}

		connInfo, exists := config.Database.Conn_urls[prefix+connUrlName]
		if !exists && config.Main.Environ != "test" {
			log.Fatalf("error: connection info \"%s\" does not exists..", config.Main.Db_prefix+prefix+connUrlName)
		} else if config.Main.Environ == "test" {
			testDb := config.Database.Conn_urls["test_db"]
			connInfo.Url = testDb.Url + config.Main.Db_prefix + prefix + connUrlName
			connInfo.ConnMaxIdleTime = testDb.ConnMaxIdleTime
			connInfo.MaxOpenConns = testDb.MaxOpenConns
		}

		ds.Source = sql.NewSQLDataSource(drvName, connInfo.Url)
		db, err := ds.Source.Connect()
		if err != nil {
			ds.IsValid = false
			ds_err = err
		}

		if ds.IsValid {
			db.SetConnMaxIdleTime(connInfo.ConnMaxIdleTime)
			db.SetMaxOpenConns(int(connInfo.MaxOpenConns))
		}

		if initCb != nil {
			initCb(ds)
		}
	})
}

func OpenSqlDbError() error {
	return ds_err
}

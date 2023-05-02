package store

import (
	"errors"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/romv7/blogs/internal"
	"github.com/romv7/blogs/internal/store/source/sql"
)

var ds_err error
var ErrNoConnUrlFound = errors.New("no connection info found.")

func OpenSqlDb(dbName string, initCb func(*DataSource[sql.SQLDataSource])) (*DataSource[sql.SQLDataSource], error) {
	config, err := internal.LoadConfig()
	if err != nil {
		return nil, err
	}

	drvName := config.Database.Drv_name

	return NewDataSource(func(ds *DataSource[sql.SQLDataSource]) {
		prefix := config.Main.Environ + "_"

		connInfo, exists := config.Database.Conn_urls[prefix+dbName]
		if !exists {
			log.Fatalf("error: connection info \"%s\" does not exists..", config.Main.Db_prefix+dbName)
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

func GetDBConnConfig(dbName string) (out *mysql.Config, err error) {
	config, err := internal.LoadConfig()
	if err != nil {
		return nil, err
	}

	connInfo, exists := config.Database.Conn_urls[config.Main.Environ+"_"+dbName]
	if !exists {
		return nil, ErrNoConnUrlFound
	}

	return mysql.ParseDSN(connInfo.Url)
}

func GetPartCount(dbName string) uint {
	config, err := internal.LoadConfig()
	if err != nil {
		log.Fatalf(err.Error())
	}

	connInfo, exists := config.Database.Conn_urls[config.Main.Environ+"_"+dbName]
	if !exists {
		log.Fatalf(ErrNoConnUrlFound.Error())
	}

	return connInfo.Partitions
}

func OpenSqlDbError() error {
	return ds_err
}

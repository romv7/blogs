package sql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Store() *gorm.DB {
	if db != nil {
		return db
	}

	var err error

	db, err = gorm.Open(mysql.New(sqlDialectorConfig), gormConfig)
	if err != nil {
		panic(err)
	}

	return db
}

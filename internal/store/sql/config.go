package sql

import (
	"os"
	"time"

	"github.com/go-sql-driver/mysql"

	gorm_mysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	sqlConfig = &mysql.Config{
		User:   os.Getenv("MYSQL_USER"),
		Passwd: os.Getenv("MYSQL_PASSWORD"),
		Net:    os.Getenv("MYSQL_NET"),
		Addr:   os.Getenv("MYSQL_ADDR"),
		DBName: os.Getenv("MYSQL_DATABASE"),

		ParseTime:            true,
		AllowNativePasswords: true,
		Loc:                  time.UTC,
	}

	gormConfig = &gorm.Config{}

	sqlDialectorConfig = gorm_mysql.Config{
		DSNConfig:         sqlConfig,
		DefaultStringSize: 256,
	}
)

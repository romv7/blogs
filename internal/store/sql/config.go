package sql

import (
	"fmt"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"

	gorm_mysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type StderrLogger struct {
	f *os.File
}

func NewStderrLogger(f *os.File) *StderrLogger {
	return &StderrLogger{f}
}

func (l *StderrLogger) Printf(fmtstr string, args ...any) {
	fmt.Printf(fmtstr, args...)
	l.f.Write([]byte(fmt.Sprintf(fmtstr, args...)))
}

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

	gormConfig = &gorm.Config{
		PrepareStmt: true,
		Logger:      logger.New(NewStderrLogger(os.Stdout), logger.Config{}),
	}

	sqlDialectorConfig = gorm_mysql.Config{
		DSNConfig:         sqlConfig,
		DefaultStringSize: 256,
	}
)

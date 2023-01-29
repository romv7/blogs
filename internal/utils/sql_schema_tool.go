package utils

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/rommms07/blogs/internal"
	"github.com/rommms07/blogs/internal/store"
	"github.com/rommms07/blogs/internal/store/source/sql"

	dsql "database/sql"
)

type SQLSchemaTool struct{}

var (
	config      *internal.ConfigSchema
	memo_ents   []os.DirEntry
	schema_root string

	loadDbFromDataSource = func(connName string) (*sql.SQLDataSource, error) {
		db, err := store.OpenSqlDb(connName, nil)
		ds_err := store.OpenSqlDbError()

		if err != nil {
			return db.Source.(*sql.SQLDataSource), err
		} else if ds_err != nil {
			return db.Source.(*sql.SQLDataSource), ds_err
		}

		return db.Source.(*sql.SQLDataSource), err
	}

	loopThruEntries = func(entCb func(os.DirEntry) error, rev bool) (r_err error) {
		execCallback := func(entry os.DirEntry) error {
			if r_err != nil {
				return r_err
			}

			if err := entCb(entry); err != nil {
				r_err = err
				return err
			}

			return nil
		}

		for _, entry := range memo_ents {
			if rev {
				defer execCallback(entry)
				continue
			}

			r_err = execCallback(entry)
		}

		return
	}

	removeExt = func(db_file string) string {
		// @reimplement
		return strings.Split(strings.Split(db_file, ".")[0], "-")[1]
	}

	initEntries = func() error {
		err := error(nil)

		if memo_ents == nil {
			memo_ents, err = os.ReadDir(internal.GetSchemaRootDir())
			return err
		}

		return err
	}

	execQuery = func(db *sql.SQLDataSource, query string, args ...any) (dsql.Result, error) {
		return db.Exec(query, args...)
	}
)

func init() {
	if conf, err := internal.LoadConfig(); err == nil {
		config = conf
	} else {
		log.Fatalf("error: %s", err.Error())
	}

	schema_root = internal.GetSchemaRootDir()
}

func NewSQLSchemaTool() *SQLSchemaTool {
	return &SQLSchemaTool{}
}

func (l *SQLSchemaTool) Load(connName, fname string) error {
	db, err := loadDbFromDataSource(connName)
	if err != nil {
		return err
	}

	sqlf, err := os.ReadFile(schema_root + "/" + fname)
	if err != nil {
		return err
	}

	prefix := config.Main.Environ + "_"
	numParts := config.Database.Conn_urls[prefix+connName].Partitions
	if numParts == 0 {
		numParts = 1
	}

	for i := uint(0); i < numParts; i++ {
		_, err = execQuery(db, strings.ReplaceAll(string(sqlf), connName, fmt.Sprintf("%s%d", connName, i)))
		if err != nil {
			return err
		}
	}

	return nil
}

func (l *SQLSchemaTool) Reload(connName, name string) (err error) {
	if err = l.Drop(connName, name); err != nil {
		return
	}

	err = l.Load(connName, name)
	return
}

func (l *SQLSchemaTool) Drop(connName, name string) error {
	db, err := loadDbFromDataSource(connName)
	if err != nil {
		return err
	}

	prefix := config.Main.Environ + "_"
	numParts := config.Database.Conn_urls[prefix+connName].Partitions
	if numParts == 0 {
		numParts = 1
	}

	for i := uint(0); i < numParts; i++ {
		_, err = execQuery(db, fmt.Sprintf("drop table if exists `%s`;", fmt.Sprintf("%s%d", name, i)))
		if err != nil {
			return err
		}
	}

	return err
}

func allOpGeneralFunc(op func(en os.DirEntry) error, rev bool) (err error) {
	if err = initEntries(); err != nil {
		return
	}

	err = loopThruEntries(func(en os.DirEntry) error {
		return op(en)
	}, rev)

	return
}

func (l *SQLSchemaTool) ReloadAll() (err error) {

	err = allOpGeneralFunc(func(en os.DirEntry) error {
		no_ext := removeExt(en.Name())

		err := l.Drop(no_ext, no_ext)
		if err != nil {
			return err
		}

		name := en.Name()
		err = l.Load(no_ext, name)
		if err != nil {
			return err
		}

		return nil
	}, false)

	return
}

func (l *SQLSchemaTool) LoadAll() (err error) {
	err = allOpGeneralFunc(func(en os.DirEntry) error {
		name := en.Name()
		return l.Load(removeExt(name), name)
	}, false)

	return
}

func (l *SQLSchemaTool) DropAll() (err error) {
	err = allOpGeneralFunc(func(en os.DirEntry) error {
		name := removeExt(en.Name())
		return l.Drop(name, name)
	}, true)

	return
}

func initMockDb() {

}

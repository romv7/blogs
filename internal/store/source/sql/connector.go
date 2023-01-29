package sql

import (
	"database/sql"
	"log"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rommms07/blogs/internal"
)

var opened map[string]*sql.DB

type SQLDataSource struct {
	n_conn, is_test        bool
	db_name, drv_name, dsn string
	*sql.DB
}

func mockReqVars() (config *internal.ConfigSchema, prefix, testDbDsn string) {
	config, _ = internal.LoadConfig()
	prefix = config.Main.Db_prefix + config.Main.Environ + "_"
	testDbDsn = config.Database.Conn_urls["test_db"].Url

	return
}

func NewSQLDataSource(drvName string, dataSourceUrl string) *SQLDataSource {
	return &SQLDataSource{drv_name: drvName, dsn: dataSourceUrl}
}

func (s *SQLDataSource) Connect() (*SQLDataSource, error) {
	if db, exists := opened[s.dsn]; exists && s.n_conn != true {
		s.DB = db
		s.n_conn = false
		return s, nil
	}

	db, err := sql.Open(s.drv_name, s.dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	s.DB = db
	return s, nil
}

func (s *SQLDataSource) NConnect() (*SQLDataSource, error) {
	s.n_conn = true
	return s.Connect()
}


// To avoid repeatedly calling internal.LoadConfig, it is wise to memoize the number
// of partitions get from the config to a map. This will reduce the overhead of accessing
// the number of partitions from the config.
var N_PART_CACHE_MAP = make(map[string]uint)

// GetTableIdByUnix returns a remainder of a timestamps based on the
// selected sql data source. For example, [unix ts] % (num_of_partitions) == [0, num_of_partitions-1],
// this utility function, tells us what table our query will point to
// based on a unix timestamp. The drawback of this approach is, of course
// accessing the data.
//
// We can solve the issue by including a histories table, which we can use to store info about a query,
// and whenever we wanted to access a particular data we will reference this histories
// table for information about an action.
func (s *SQLDataSource) GetTableIdByUnix(db_name string, t int64) int64 {	
	// number of partitions the database contains.
	N := uint(0)

	if _, exists := N_PART_CACHE_MAP[db_name]; !exists {
		config, _ := internal.LoadConfig()
		dbconf := config.Database.Conn_urls
		env := config.Main.Environ

		N = dbconf[env+"_"+db_name].Partitions
		N_PART_CACHE_MAP[db_name] = N
	} else {
		N = N_PART_CACHE_MAP[db_name]
	}

	return t%int64(N)
}

func (s *SQLDataSource) GetDestTableByUnix(db_name string, t int64) string {
	N := s.GetTableIdByUnix(db_name, t)
	return fmt.Sprintf("%s%d", db_name, N)
}

func (s *SQLDataSource) InitWithMockDb(dbName string) {
	_, prefix, testDbDsn := mockReqVars()
	s.db_name = prefix + dbName
	s.is_test = true

	// create test database
	if tmpDb, err := sql.Open(s.drv_name, testDbDsn); err != nil {
		log.Fatalf(err.Error())
	} else {
		_, err := tmpDb.Exec(strings.ReplaceAll("create database `$`;", "$", s.db_name))
		if err != nil {
			log.Fatalf(err.Error())
		}

		_, err = tmpDb.Exec(strings.ReplaceAll("use `$`;", "$", s.db_name))
		if err != nil {
			log.Fatalf(err.Error())
		}

		s.DB = tmpDb
	}

	if s.DB == nil {
		log.Fatalf("error: did not properly initialized the test database...")
	}
}

func (s *SQLDataSource) DetachIfMock() (err error) {
	if !s.is_test {
		return
	}

	_, err = s.Exec(strings.ReplaceAll("drop database `$`;", "$", s.db_name))
	return
}

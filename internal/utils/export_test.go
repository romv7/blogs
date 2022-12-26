package utils

import (
	"fmt"
	"os"

	"github.com/rommms07/blogs/internal/store/source/sql"

	dsql "database/sql"
)

func mock_loadDbFromDataSource(connName string) (*sql.SQLDataSource, error) {
	return sql.NewSQLDataSource("mysql", connName), nil
}

func mock_execQuery(db *sql.SQLDataSource, query string, args ...any) (dsql.Result, error) {
	execQuery_in = query
	execQuery_args = args
	return nil, nil
}

func mock_initEntries() error {
	initEntries_isExec = true
	return nil
}

var (
	save_loadDbFromDataSource func(string) (*sql.SQLDataSource, error)
	save_execQuery            func(*sql.SQLDataSource, string, ...any) (dsql.Result, error)
	save_initEntries          func() error
	save_schema_root          string

	initEntries_isExec bool
	execQuery_in       string
	execQuery_args     []any
)

var (
	TestDb    = "tests"
	TestQuery = fmt.Sprintf(`
	create table if not exists %s ( 
		id int UNSIGNED,
		primary key(id),
	);
	`, TestDb)
)

func InitEntriesExecuted() bool {
	return initEntries_isExec
}

func GetExecQueryProps() (query string, args []any) {
	query = execQuery_in
	args = execQuery_args

	return
}

func LoadMockFunctions() {
	save_loadDbFromDataSource = loadDbFromDataSource
	loadDbFromDataSource = mock_loadDbFromDataSource

	save_execQuery = execQuery
	execQuery = mock_execQuery

	save_initEntries = initEntries
	initEntries = mock_initEntries

	save_schema_root = schema_root
	schema_root = "/tmp"
	initEntries_isExec = false
	execQuery_in = ""

	os.WriteFile("/tmp/"+TestDb, []byte(TestQuery), 0650)
}

func Restore() {
	loadDbFromDataSource = save_loadDbFromDataSource
	execQuery = save_execQuery
	initEntries = save_initEntries
	schema_root = save_schema_root

	os.Remove("/tmp/" + TestDb)
}

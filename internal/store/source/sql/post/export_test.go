package post

var (
	InitSql func()
)

func mock_initSql() {

}

func init() {
	InitSql = initSql
	initSql = mock_initSql
}

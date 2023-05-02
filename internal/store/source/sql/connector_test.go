package sql_test

import (
	"testing"

	"github.com/romv7/blogs/internal/store/source/sql"
)

func Test_willReturnAnErrorWhenOffline(t *testing.T) {
	src := sql.NewSQLDataSource("mysql", "rommms:07261999@unix(/var/run/mysqld/mysqld.sock)/rommms.nation")

	_, err := src.Connect()

	if src.DB == nil && err == nil {
		t.Errorf("[fail] src.Connect must return an error when there's is no active SQL connection.")
	} else if src.DB == nil {
		if err == nil {
			t.Errorf("[fail] expected src.Handle() to hold a value after a successful connection.")
		}
	} else {
		src.Close()
	}
}

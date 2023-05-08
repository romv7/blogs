package sql

import "github.com/romv7/blogs/internal/store/sql/models"

var (
	MODELS = []interface{}{
		&models.User{}, &models.Post{}, &models.Comment{},
	}
)

func Migrate() (err error) {
	store := Store()

	for _, m := range MODELS {
		if err = store.AutoMigrate(m); err != nil {
			break
		}
	}

	return
}

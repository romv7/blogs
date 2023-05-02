package post

import (
	"fmt"
	"strings"

	"github.com/romv7/blogs/internal/entities"
	"github.com/romv7/blogs/internal/store"
	"github.com/romv7/blogs/internal/store/source/sql"
	"github.com/romv7/blogs/internal/store/source/sql/uuidindexes"
)

var postsDb *sql.SQLDataSource

const (
	db_name = "posts"
)

type PostStoreSql struct {
	T int64
	store.UnimplementedStore
}

var initSql = func() {
	store.InitSqlDb(db_name, func(db *sql.SQLDataSource) {
		postsDb = db
	})
}

func (s *PostStoreSql) Save(post *entities.Post) (err error) {
	initSql()

	tbl := postsDb.GetDestTableByUnix(
		db_name,
		post.State.CreatedAt.AsTime().Unix(),
	)

	uindex := uuidindexes.NewUuidIndex(post.Id, db_name, tbl, post.Uuid, post.UniqueKey())

	query := fmt.Sprintf(`
	insert into %s
		(
			id,
			user_id,
			uuid,
			headline_text,
			summary_text,
			tags,
			stage,
			status,
			revised_at,
			archived_at,
			published_at,
			created_at,
			original_id
		) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	on duplicate key update
		user_id=values(user_id),
		uuid=values(uuid),
		headline_text=values(headline_text),
		summary_text=values(summary_text),
		tags=values(tags),
		stage=values(stage),
		status=values(status),
		revised_at=values(revised_at),
		archived_at=values(archived_at),
		published_at=values(published_at),
		created_at=values(created_at),
		original_id=values(original_id);`, tbl)

	_, err = postsDb.Query(query,
		post.Id,
		post.User.Id,
		post.Uuid,
		post.HeadlineText,
		post.SummaryText,
		strings.Join(post.Tags, ","),
		post.State.Stage,
		post.State.Status,
		post.State.RevisedAt.AsTime(),
		nil,
		nil,
		post.State.CreatedAt.AsTime(),
		nil,
	)

	if err != nil {
		return err
	}

	return uindex.Save()
}

func (s *PostStoreSql) Delete(p *entities.Post) (err error) {
	initSql()

	ui, err := uuidindexes.GetUuidIndex(db_name, p.UniqueKey())
	if err != nil {
		return err
	}

	query := fmt.Sprintf(`delete low_priority from %s where id=? or uuid=? or headline_text=? and user_id=?`, ui.Ref)
	if _, err := postsDb.Query(query, p.Id, p.Uuid, p.HeadlineText, p.User.Id); err != nil {
		return err
	}

	err = ui.Revoke()

	return
}

func (s *PostStoreSql) GetByUuid(query string, args ...any) (err error) {
	initSql()

	return
}

func DbName() string {
	return db_name
}

package models

import (
	"database/sql"

	"github.com/coopernurse/gorp"
	_ "github.com/mattn/go-sqlite3"
)

var (
	orp *gorp.DbMap
)

func InitDb() (dbmap *gorp.DbMap, err error) {

	db, err := sql.Open("sqlite3", "./blog.db")
	if err != nil {
		return nil, err
	}

	orp = &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}

	orp.AddTableWithName(Post{}, "posts").SetKeys(true, "Id")

	if err = orp.CreateTablesIfNotExists(); err != nil {
		return
	}
	dbmap = orp
	return
}

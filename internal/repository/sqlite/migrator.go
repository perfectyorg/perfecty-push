package sqlite

import (
	"database/sql"
	"github.com/pressly/goose"
	_ "github.com/perfectyorg/perfecty-push/internal/repository/sqlite/migrations"
)

const dialect = "sqlite3"

func Migrate(db *sql.DB) (err error) {
	if err = goose.SetDialect(dialect); err != nil {
		return
	}

	// it will run the migrations registered at _ (/internal/repository/sqlite/migrations)
	return goose.Up(db, ".")
}

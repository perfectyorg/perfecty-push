package sqlite_test

import (
	"database/sql"
	"github.com/perfectyorg/perfecty-push/internal/repository/sqlite"
)

const dsn = "file:test.db?mode=memory"
const driver = "sqlite3"

func setupDB() (db *sql.DB) {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		panic(err)
	}
	if err = sqlite.Migrate(db); err != nil {
		panic(err)
	}
	return
}

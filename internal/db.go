package internal

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func StartDB() (db *sql.DB, err error) {
	driver := Config.Database.Driver
	dsn := Config.Database.DSN

	db, err = sql.Open(driver, dsn)
	if err != nil {
		return
	}
	defer db.Close()

	return
}

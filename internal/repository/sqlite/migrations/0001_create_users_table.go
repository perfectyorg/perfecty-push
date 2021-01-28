package migrations

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up00001, Down00001)
}

func Up00001(tx *sql.Tx) (err error) {
	_, err = tx.Exec("CREATE TABLE users (" +
		"uuid text NOT NULL PRIMARY KEY" +
		");")
	_, err = tx.Exec("INSERT INTO users VALUES('my name');")
	return
}

func Down00001(tx *sql.Tx) (err error) {
	_, err = tx.Exec("DROP TABLE users")
	return
}

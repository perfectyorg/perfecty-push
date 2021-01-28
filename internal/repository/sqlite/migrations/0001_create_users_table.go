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
		"uuid char(36) NOT NULL PRIMARY KEY" +
		");")
	return
}

func Down00001(tx *sql.Tx) (err error) {
	_, err = tx.Exec("DROP TABLE users")
	return
}

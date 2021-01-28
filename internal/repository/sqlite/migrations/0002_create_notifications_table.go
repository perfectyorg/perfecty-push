package migrations

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up00002, Down00002)
}

func Up00002(tx *sql.Tx) (err error) {
	_, err = tx.Exec(`CREATE TABLE notifications (
			uuid char(36) NOT NULL,
			payload varchar(500) NOT NULL,
			total int DEFAULT 0 NOT NULL,
			succeeded int DEFAULT 0 NOT NULL,
			last_cursor int DEFAULT 0 NOT NULL,
			batch_size int DEFAULT 0 NOT NULL,
			status varchar(15) DEFAULT 'scheduled' NOT NULL,
			is_taken tinyint DEFAULT 0 NOT NULL,
			created_at datetime DEFAULT CURRENT_TIMESTAMP NOT NULL,
			finished_at datetime NULL
		);`)

	return
}

func Down00002(tx *sql.Tx) (err error) {
	_, err = tx.Exec(`DROP TABLE notifications;`)

	return
}

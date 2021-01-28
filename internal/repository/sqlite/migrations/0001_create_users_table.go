package migrations

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up00001, Down00001)
}

func Up00001(tx *sql.Tx) (err error) {
	_, err = tx.Exec(`CREATE TABLE users (
			uuid char(36) NOT NULL,
			endpoint varchar(500) NOT NULL UNIQUE,
			remote_ip varchar(46) DEFAULT '',
			key_auth varchar(100) NOT NULL UNIQUE,
			key_p256dh varchar(100) NOT NULL UNIQUE,
			opted_in tinyint(1) DEFAULT 1 NOT NULL,
			enabled tinyint(1) DEFAULT 1 NOT NULL,
			created_at datetime DEFAULT CURRENT_TIMESTAMP NOT NULL,
			disabled_at datetime NULL,
			PRIMARY KEY (uuid)
		);
		CREATE UNIQUE index users_uuid_uk ON users(uuid);
		CREATE UNIQUE index users_endpoint_uk ON users(endpoint);`)

	return
}

func Down00001(tx *sql.Tx) (err error) {
	_, err = tx.Exec(`DROP TABLE users;
		DROP INDEX users_uuid_uk;
		DROP INDEX users_endpoint_uk;`)

	return
}

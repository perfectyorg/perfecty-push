package internal

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/rs/zerolog/log"
	"github.com/rwngallego/perfecty-push/internal/application"
	"github.com/rwngallego/perfecty-push/internal/repository/sqlite"
)

var db *sql.DB

//StartDB Starts the DB and returns the repository implementations according to the driver
func StartDB() (userRepository application.UserRepository, notificationRepository application.NotificationRepository, err error) {
	driver := Config.Database.Driver
	dsn := Config.Database.DSN

	db, err = sql.Open(driver, dsn)
	if err != nil {
		return
	}

	switch driver {
	case "sqlite3":
		if err = sqlite.Migrate(db); err != nil {
			log.Error().Err(err).Msg("Could not run the migrations")
			return
		}
		userRepository = sqlite.NewSqlLiteUserRepository(db)
		notificationRepository = sqlite.NewSqlLiteNotificationRepository(db)
		return
	default:
		err = fmt.Errorf("driver \"%s\" is not supported", driver)
		return
	}
}

func StopDB() (err error) {
	log.Info().Msg("Closing the DB")
	if db != nil {
		return db.Close()
	}
	return
}

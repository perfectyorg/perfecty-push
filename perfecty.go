package perfecty

import (
	"database/sql"
	"github.com/rwngallego/perfecty-push/internal"
	"github.com/rwngallego/perfecty-push/internal/application"
	"github.com/rwngallego/perfecty-push/internal/repository/sqlite"
)

const filePath = "configs/internal.yml"

// Start Setup and start the push server
func Start() (err error) {
	var (
		db *sql.DB
	)

	if err = internal.LoadConfig(filePath); err != nil {
		return
	}

	if err = internal.LoadLogging(); err != nil {
		return
	}

	if db, err = internal.StartDB(); err != nil {
		return
	}

	userRepository := sqlite.NewSqlLiteUserRepository(db)
	registrationService := application.NewRegistrationService(userRepository)
	preferenceService := application.NewPreferenceService(userRepository)

	if err = internal.StartServer(registrationService, preferenceService); err != nil {
		return
	}

	return
}

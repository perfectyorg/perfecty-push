package perfecty

import (
	"github.com/rwngallego/perfecty-push/internal"
	"github.com/rwngallego/perfecty-push/internal/application"
)

const filePath = "configs/internal.yml"

// Start Setup and start the push server
func Start() (err error) {
	if err = internal.LoadConfig(filePath); err != nil {
		return
	}

	if err = internal.LoadLogging(); err != nil {
		return
	}

	userRepository, err := internal.StartDB()
	if err != nil {
		return
	}
	defer internal.StopDB()

	registrationService := application.NewRegistrationService(userRepository)
	preferenceService := application.NewPreferenceService(userRepository)

	if err = internal.StartServer(registrationService, preferenceService); err != nil {
		return
	}

	return
}

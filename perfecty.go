package perfecty

import (
	"github.com/rwngallego/perfecty-push/internal"
	"github.com/rwngallego/perfecty-push/internal/application"
)

const filePath = "configs/perfecty.yml"

// Start Start the push server
func Start() (err error) {
	if err = internal.LoadConfig(filePath); err != nil {
		return
	}

	if err = internal.LoadLogging(); err != nil {
		return
	}

	userRepository, notificationRepository, err := internal.StartDB()
	if err != nil {
		return
	}
	defer internal.StopDB()

	registrationService := application.NewRegistrationService(userRepository)
	preferenceService := application.NewPreferenceService(userRepository)
	scheduleService := application.NewScheduleService(notificationRepository)

	if err = internal.StartServer(registrationService, preferenceService, scheduleService); err != nil {
		return
	}

	return
}

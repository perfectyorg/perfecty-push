package application

import (
	"github.com/google/uuid"
	n "github.com/perfectyorg/perfecty-push/internal/domain/notification"
)

type (
	NotificationRepository interface {
		// Get Get the notifications
		Get(offset int, size int, orderBy string, orderAsc string) (notificationList []*n.Notification, err error)
		// GetById Get the notification by Id
		GetById(id uuid.UUID) (notification *n.Notification, err error)
		// Create Create the notification
		Create(notification *n.Notification) (err error)
		// Update Update the notification
		Update(notification *n.Notification) (err error)
		// Delete Delete the notification
		Delete(id uuid.UUID) (err error)
		// Stats Get the stats
		Stats() (total int, succeeded int, failed int, err error)
	}

	ScheduleService struct {
		notificationRepository NotificationRepository
	}
)

func NewScheduleService(notificationRepository NotificationRepository) (scheduleService *ScheduleService) {
	scheduleService = &ScheduleService{
		notificationRepository: notificationRepository,
	}
	return
}

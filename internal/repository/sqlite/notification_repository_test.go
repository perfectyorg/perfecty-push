package sqlite_test

import (
	n "github.com/rwngallego/perfecty-push/internal/domain/notification"
	"github.com/rwngallego/perfecty-push/internal/repository/sqlite"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	batchSize = 30
)

func TestGetNotifications(t *testing.T) {
	var (
		db               = setupDB()
		notificationRepo = sqlite.NewSqlLiteNotificationRepository(db)
		notification1, _ = n.NewNotification("payload_test_1", 2, batchSize, n.StatusScheduled)
		notification2, _ = n.NewNotification("payload_test_2", 3, batchSize, n.StatusScheduled)
	)
	defer db.Close()
	notification1.Succeeded = 1
	notification1.Take()
	notificationRepo.Create(notification1)
	notificationRepo.Create(notification2)

	notifications, err := notificationRepo.Get(0, 3, "created_at", "asc")

	assert.Equal(t, 2, len(notifications))
	assert.NoError(t, err)
	assert.Equal(t, "payload_test_1", notifications[0].Payload)
	assert.Equal(t, 2, notifications[0].Total)
	assert.Equal(t, batchSize, notifications[0].BatchSize)
	assert.Equal(t, n.StatusScheduled, notifications[0].Status())
	assert.Equal(t, 0, notifications[0].LastCursor)
	assert.Equal(t, 1, notifications[0].Succeeded)
	assert.Equal(t, true, notifications[0].IsTaken())
	assert.Equal(t, false, notifications[0].CreatedAt().IsZero())
}

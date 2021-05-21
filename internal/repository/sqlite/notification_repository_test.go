package sqlite_test

import (
	"github.com/google/uuid"
	n "github.com/perfectyorg/perfecty-push/internal/domain/notification"
	"github.com/perfectyorg/perfecty-push/internal/repository/sqlite"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
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

func TestGetNotificationById(t *testing.T) {
	var (
		db                = setupDB()
		id, _             = uuid.NewRandom()
		timeCreated       = time.Now()
		timeStatusChanged = time.Now()
		notificationRepo  = sqlite.NewSqlLiteNotificationRepository(db)
		notification, _   = n.NewNotificationRaw(id, "payload_test_1", 7, 4, 3, 10, n.StatusCompleted, true, timeCreated, timeStatusChanged)
	)
	defer db.Close()
	notificationRepo.Create(notification)

	created, err := notificationRepo.GetById(notification.Uuid)

	assert.NoError(t, err)
	assert.Equal(t, "payload_test_1", created.Payload)
	assert.Equal(t, 7, created.Total)
	assert.Equal(t, 10, created.BatchSize)
	assert.Equal(t, n.StatusCompleted, created.Status())
	assert.Equal(t, 3, notification.LastCursor)
	assert.Equal(t, 4, notification.Succeeded)
	assert.Equal(t, true, notification.IsTaken())
	assert.Equal(t, false, notification.CreatedAt().IsZero())
	assert.Equal(t, false, notification.StatusChangedAt().IsZero())
}

func TestCreateNotification(t *testing.T) {
	var (
		db               = setupDB()
		notificationRepo = sqlite.NewSqlLiteNotificationRepository(db)
		notification, _  = n.NewNotification("payload_test_1", 2, batchSize, n.StatusScheduled)
	)
	defer db.Close()

	err := notificationRepo.Create(notification)

	created, errGetById := notificationRepo.GetById(notification.Uuid)
	assert.NoError(t, err)
	assert.NoError(t, errGetById)
	assert.NotNil(t, created)
	assert.Equal(t, "payload_test_1", created.Payload)
	assert.Equal(t, 2, created.Total)
	assert.Equal(t, batchSize, created.BatchSize)
	assert.Equal(t, n.StatusScheduled, created.Status())
	assert.Equal(t, 0, notification.LastCursor)
	assert.Equal(t, 0, notification.Succeeded)
	assert.Equal(t, false, notification.IsTaken())
	assert.Equal(t, false, notification.CreatedAt().IsZero())
	assert.Equal(t, false, notification.StatusChangedAt().IsZero())
}
func TestUpdateNotification(t *testing.T) {
	var (
		db               = setupDB()
		notificationRepo = sqlite.NewSqlLiteNotificationRepository(db)
		notification, _  = n.NewNotification("payload_test_1", 2, batchSize, n.StatusScheduled)
	)
	defer db.Close()
	errCreated := notificationRepo.Create(notification)

	notification.Payload = "payload_test_2"
	notification.Total = 5
	notification.BatchSize = 7
	notification.SetStatus(n.StatusFailed)
	notification.Succeeded = 2
	notification.LastCursor = 4
	notification.Take()
	err := notificationRepo.Update(notification)
	updated, errGetById := notificationRepo.GetById(notification.Uuid)

	assert.NoError(t, errCreated)
	assert.NoError(t, err)
	assert.NoError(t, errGetById)
	assert.NotNil(t, updated)
	assert.Equal(t, "payload_test_2", updated.Payload)
	assert.Equal(t, 5, updated.Total)
	assert.Equal(t, 7, updated.BatchSize)
	assert.Equal(t, n.StatusFailed, updated.Status())
	assert.Equal(t, 4, updated.LastCursor)
	assert.Equal(t, 2, updated.Succeeded)
	assert.Equal(t, true, updated.IsTaken())
	assert.Equal(t, false, updated.CreatedAt().IsZero())
	assert.Equal(t, false, updated.StatusChangedAt().IsZero())
}

func TestDeleteNotification(t *testing.T) {
	var (
		db               = setupDB()
		notificationRepo = sqlite.NewSqlLiteNotificationRepository(db)
		notification, _  = n.NewNotification("payload_test_1", 2, batchSize, n.StatusScheduled)
	)
	defer db.Close()
	errCreate := notificationRepo.Create(notification)
	created, errGetById := notificationRepo.GetById(notification.Uuid)

	err := notificationRepo.Delete(notification.Uuid)
	deleted, errGetByIdDeleted := notificationRepo.GetById(notification.Uuid)

	assert.NoError(t, errCreate)
	assert.NotNil(t, created)
	assert.NoError(t, errGetById)
	assert.NoError(t, err)
	assert.Error(t, errGetByIdDeleted)
	assert.Nil(t, deleted)
}

func TestGetNotificationsStats(t *testing.T) {
	var (
		db               = setupDB()
		notificationRepo = sqlite.NewSqlLiteNotificationRepository(db)
		notification1, _ = n.NewNotification("payload_test_1", 5, batchSize, n.StatusScheduled)
		notification2, _ = n.NewNotification("payload_test_2", 7, batchSize, n.StatusScheduled)
		notification3, _ = n.NewNotification("payload_test_3", 10, batchSize, n.StatusScheduled)
	)
	defer db.Close()

	notification1.Succeeded = 3
	notification2.Succeeded = 2
	notification3.Succeeded = 7
	notificationRepo.Create(notification1)
	notificationRepo.Create(notification2)
	notificationRepo.Create(notification3)

	total, succeeded, failed, err := notificationRepo.Stats()

	assert.NoError(t, err)
	assert.Equal(t, 22, total)
	assert.Equal(t, 12, succeeded)
	assert.Equal(t, 10, failed)
}

package notification_test

import (
	n "github.com/perfectyorg/perfecty-push/internal/domain/notification"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewNotification(t *testing.T) {
	var (
		notification *n.Notification
	)

	notification, err := n.NewNotification(`{"name": "test"}`, 30, 200, n.StatusCompleted)

	assert.NoError(t, err)
	assert.Equal(t, `{"name": "test"}`, notification.Payload)
	assert.Equal(t, 30, notification.Total)
	assert.Equal(t, 200, notification.BatchSize)
	assert.Equal(t, n.StatusCompleted, notification.Status())
	assert.Equal(t, false, notification.CreatedAt().IsZero())
}

func TestChangeStatus(t *testing.T) {
	notification := createNotification()
	notification.SetStatus(n.StatusFailed)
	assert.Equal(t, n.StatusFailed, notification.Status())
}

func TestNotificationTakeRelease(t *testing.T) {
	notification := createNotification()
	notification.Take()
	assert.Equal(t, true, notification.IsTaken())
	notification.Release()
	assert.Equal(t, false, notification.IsTaken())
}

func createNotification() (notification *n.Notification) {
	notification, _ = n.NewNotification(
		`{"name": "test"}`,
		30,
		200,
		n.StatusCompleted,
	)
	return
}

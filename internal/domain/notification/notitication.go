package notification

import (
	"github.com/google/uuid"
	"time"
)

const (
	StatusScheduled = iota
	StatusRunning
	StatusFailed
	StatusCompleted
)

type (
	Notification struct {
		Uuid            uuid.UUID
		Payload         string
		Total           int
		Succeeded       int
		LastCursor      int
		BatchSize       int
		status          int
		isTaken         bool
		createdAt       time.Time
		statusChangedAt time.Time
	}
)

func NewNotification(payload string, total int, batchSize int, status int) (notification *Notification, err error) {

	id, err := uuid.NewRandom()
	if err != nil {
		return
	}

	notification = &Notification{
		Uuid:            id,
		Payload:         payload,
		Total:           total,
		Succeeded:       0,
		LastCursor:      0,
		BatchSize:       batchSize,
		status:          status,
		isTaken:         false,
		createdAt:       time.Now(),
		statusChangedAt: time.Now(),
	}
	return
}

func NewNotificationRaw(uuid uuid.UUID, payload string, total int, succeeded int, lastCursor int, batchSize int, status int, isTaken bool, createdAt time.Time, statusChangedAt time.Time) (notification *Notification, err error) {
	notification = &Notification{
		Uuid:            uuid,
		Payload:         payload,
		Total:           total,
		Succeeded:       succeeded,
		LastCursor:      lastCursor,
		BatchSize:       batchSize,
		status:          status,
		isTaken:         isTaken,
		createdAt:       createdAt,
		statusChangedAt: statusChangedAt,
	}
	return
}

// Getters

func (n *Notification) CreatedAt() time.Time {
	return n.createdAt
}

func (n *Notification) StatusChangedAt() time.Time {
	return n.statusChangedAt
}

func (n *Notification) Status() int {
	return n.status
}

func (n *Notification) IsTaken() bool {
	return n.isTaken
}

// Setters

func (n *Notification) SetStatus(status int) {
	n.status = status
	n.statusChangedAt = time.Now()
}

func (n *Notification) Take() {
	n.isTaken = true
}

func (n *Notification) Release() {
	n.isTaken = false
}

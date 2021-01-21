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
		Uuid       uuid.UUID
		Payload    string
		Total      int64
		Succeeded  int64
		LastCursor int64
		BatchSize  int64
		status     int
		IsTaken    bool
		createdAt  time.Time
		FinishedAt *time.Time
	}
)

func NewNotification(payload string, total int64, batchSize int64, status int) (notification *Notification, err error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return
	}

	notification = &Notification{
		Uuid:       id,
		Payload:    payload,
		Total:      total,
		Succeeded:  0,
		LastCursor: 0,
		BatchSize:  batchSize,
		status:     status,
		IsTaken:    false,
		createdAt:  time.Now(),
	}
	return
}

// Getters

func (n *Notification) CreatedAt() time.Time {
	return n.createdAt
}

func (n *Notification) Status() int {
	return n.status
}

// Setters

func (n *Notification) SetStatus(status int) {
	n.status = status
}

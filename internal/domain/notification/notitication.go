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
		isTaken    bool
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
		isTaken:    false,
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

func (n *Notification) IsTaken() bool {
	return n.isTaken
}

// Setters

func (n *Notification) SetStatus(status int) {
	n.status = status
}

func (n *Notification) Take() {
	n.isTaken = true
}

func (n *Notification) Release() {
	n.isTaken = false
}

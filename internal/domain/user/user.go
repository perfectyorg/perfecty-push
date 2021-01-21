package user

import (
	"github.com/google/uuid"
	"time"
)

type (
	User struct {
		Uuid       uuid.UUID
		Endpoint   string
		RemoteIP   string
		KeyAuth    string
		KeyP256DH  string
		optedIn    bool
		enabled    bool
		createdAt  time.Time
		disabledAt *time.Time
	}
)

func NewUser(endpoint string, remoteIP string, keyAuth string, keyP256DH string) (user *User, err error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return
	}

	user = &User{
		Uuid:      id,
		Endpoint:  endpoint,
		RemoteIP:  remoteIP,
		KeyAuth:   keyAuth,
		KeyP256DH: keyP256DH,
		optedIn:   true,
		enabled:   true,
		createdAt: time.Now()}
	return
}

// Access

func (u *User) CreatedAt() time.Time {
	return u.createdAt
}

func (u *User) DisabledAt() *time.Time {
	return u.disabledAt
}

func (u *User) IsEnabled() bool {
	return u.enabled
}

func (u *User) IsOptedIn() bool {
	return u.optedIn
}

// Modifiers

func (u *User) Enable() {
	u.disabledAt = nil
	u.enabled = true
}

func (u *User) Disable() {
	u.enabled = false
	u.disabledAt = new(time.Time)
	*u.disabledAt = time.Now()
}

func (u *User) OptIn() {
	u.optedIn = true
}

func (u *User) OptOut() {
	u.optedIn = false
}

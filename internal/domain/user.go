package domain

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
		IsActive   bool
		Disabled   bool
		CreatedAt  time.Time
		DisabledAt time.Time
	}
)

func NewUser(Uuid uuid.UUID, Endpoint string, RemoteIP string, KeyAuth string, KeyP256DH string, IsActive bool, Disabled bool) (user *User) {
	user = &User{
		Uuid:      Uuid,
		Endpoint:  Endpoint,
		RemoteIP:  RemoteIP,
		KeyAuth:   KeyAuth,
		KeyP256DH: KeyP256DH,
		IsActive:  IsActive,
		Disabled:  Disabled}
	return
}

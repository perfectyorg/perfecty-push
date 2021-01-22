package application

import (
	"github.com/google/uuid"
	u "github.com/rwngallego/perfecty-push/internal/domain/user"
)

//go:generate mockgen -destination=../../mocks/application_mock.go -package=mocks github.com/rwngallego/perfecty-push/internal/application UserRepository

type (
	UserRepository interface {
		Get(offset int, size int, orderBy string, orderAsc string, onlyActive bool) (userList []u.User, err error)
		GetById(id uuid.UUID) (user *u.User, err error)
		GetByEndpoint(endpoint string) (user *u.User, err error)
		Create(user *u.User) (err error)
		Update(user *u.User) (err error)
		Delete(id uuid.UUID) (err error)
		GetTotal() (total int, err error)
		Stats() (total int, active int, inactive int)
	}

	RegistrationService struct {
		repository UserRepository
	}
)

// NewRegistrationService Creates a registration service with the provided repository implementation
func NewRegistrationService(userRepository UserRepository) (registrationService *RegistrationService) {
	registrationService = &RegistrationService{repository: userRepository}
	return
}

func (r *RegistrationService) RegisterUser(endpoint string, remoteIp string, keyAuth string, keyP256DH string) (user *u.User, err error) {
	user, err = r.repository.GetByEndpoint(endpoint)
	if user != nil && err == nil {
		err = r.registerExistingUser(user, remoteIp, keyAuth, keyP256DH)
		return
	} else {
		user, err = r.registerNewUser(endpoint, remoteIp, keyAuth, keyP256DH)
		return
	}
}

func (r *RegistrationService) registerExistingUser(user *u.User, remoteIp string, keyAuth string, keyP256DH string) (err error) {
	user.RemoteIP = remoteIp
	user.KeyAuth = keyAuth
	user.KeyP256DH = keyP256DH
	user.Enable()
	return r.repository.Update(user)
}

func (r *RegistrationService) registerNewUser(endpoint string, remoteIp string, keyAuth string, keyP256DH string) (user *u.User, err error) {
	user, err = u.NewUser(endpoint, remoteIp, keyAuth, keyP256DH)
	if err != nil {
		return
	}

	err = r.repository.Create(user)
	return
}

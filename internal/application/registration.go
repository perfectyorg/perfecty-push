package application

import (
	"github.com/google/uuid"
	u "github.com/perfectyorg/perfecty-push/internal/domain/user"
)

//go:generate mockgen -destination=../../mocks/application_mock.go -package=mocks github.com/perfectyorg/perfecty-push/internal/application UserRepository

type (
	UserRepository interface {
		// Get Get the users
		Get(offset int, size int, orderBy string, orderAsc string, onlyActive bool) (userList []*u.User, err error)
		// GetById Get the user by Id
		GetById(id uuid.UUID) (user *u.User, err error)
		// GetByEndpoint Get the user by endpoint
		GetByEndpoint(endpoint string) (user *u.User, err error)
		// Create Create the user
		Create(user *u.User) (err error)
		// Update Update the user
		Update(user *u.User) (err error)
		// Delete Delete the user by uuid
		Delete(id uuid.UUID) (err error)
		// GetTotal Get the total number of users
		GetTotal(onlyActive bool) (total int, err error)
		// Stats Get the stats
		Stats() (total int, active int, inactive int, err error)
	}

	RegistrationService struct {
		userRepository UserRepository
	}
)

// NewRegistrationService Creates a registration service with the provided userRepository implementation
func NewRegistrationService(userRepository UserRepository) (registrationService *RegistrationService) {
	registrationService = &RegistrationService{userRepository: userRepository}
	return
}

func (r *RegistrationService) RegisterUser(endpoint string, remoteIp string, keyAuth string, keyP256DH string) (user *u.User, err error) {
	user, err = r.userRepository.GetByEndpoint(endpoint)
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
	return r.userRepository.Update(user)
}

func (r *RegistrationService) registerNewUser(endpoint string, remoteIp string, keyAuth string, keyP256DH string) (user *u.User, err error) {
	user, err = u.NewUser(endpoint, remoteIp, keyAuth, keyP256DH)
	if err != nil {
		return
	}

	err = r.userRepository.Create(user)
	return
}

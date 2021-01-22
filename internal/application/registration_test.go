package application_test

import (
	"github.com/golang/mock/gomock"
	"github.com/rwngallego/perfecty-push/internal/application"
	"github.com/rwngallego/perfecty-push/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegisterNewUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockUserRepository(ctrl)
	repo.EXPECT().GetByEndpoint(gomock.Any()).Return(nil, nil)
	repo.EXPECT().Create(gomock.Any()).Return(nil)

	registrationService := application.NewRegistrationService(repo)
	user, err := registrationService.RegisterUser(
		"test_endpoint",
		"test_remote_ip",
		"test_key_auth",
		"test_key_p256_dh",
	)

	assert.NoError(t, err)
	assert.Equal(t, "test_endpoint", user.Endpoint)
	assert.Equal(t, "test_remote_ip", user.RemoteIP)
	assert.Equal(t, "test_key_auth", user.KeyAuth)
	assert.Equal(t, "test_key_p256_dh", user.KeyP256DH)
	assert.Equal(t, true, user.IsEnabled())
	assert.Equal(t, true, user.IsOptedIn())
	assert.Equal(t, false, user.CreatedAt().IsZero())
}

func TestRegisterExistingUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user := createUser()
	repo := mocks.NewMockUserRepository(ctrl)
	repo.EXPECT().GetByEndpoint(gomock.Any()).Return(user, nil)
	repo.EXPECT().Update(gomock.Any()).Return(nil)

	registrationService := application.NewRegistrationService(repo)
	user, err := registrationService.RegisterUser(
		"test_endpoint",
		"test_remote_ip_2",
		"test_key_auth_2",
		"test_key_p256_dh_2",
	)

	assert.NoError(t, err)
	assert.Equal(t, "test_endpoint", user.Endpoint)
	assert.Equal(t, "test_remote_ip_2", user.RemoteIP)
	assert.Equal(t, "test_key_auth_2", user.KeyAuth)
	assert.Equal(t, "test_key_p256_dh_2", user.KeyP256DH)
	assert.Equal(t, true, user.IsEnabled())
	assert.Equal(t, true, user.IsOptedIn())
	assert.Equal(t, false, user.CreatedAt().IsZero())
}

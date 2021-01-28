package application_test

import (
	"github.com/golang/mock/gomock"
	"github.com/rwngallego/perfecty-push/internal/application"
	u "github.com/rwngallego/perfecty-push/internal/domain/user"
	"github.com/rwngallego/perfecty-push/mocks"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpdatePreference(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockUserRepository(ctrl)
	repo.EXPECT().Update(gomock.Any()).Return(nil)

	user := createUser()
	user.OptOut()

	preferenceService := application.NewPreferenceService(repo)
	err := preferenceService.Update(user)

	assert.NoError(t, err)
}

func createUser() (user *u.User) {
	user, _ = u.NewUser(
		"test_endpoint",
		"test_remote_ip",
		"test_key_auth",
		"test_key_p256_dh")
	return
}

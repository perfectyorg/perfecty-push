package user_test

import (
	u "github.com/rwngallego/perfecty-push/internal/domain/user"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUser(t *testing.T) {
	var (
		user *u.User
	)

	user, err := u.NewUser(
		"test_endpoint",
		"test_remote_ip",
		"test_key_auth",
		"test_key_p256_dh")

	assert.NoError(t, err)
	assert.Equal(t, "test_endpoint", user.Endpoint)
	assert.Equal(t, "test_remote_ip", user.RemoteIP)
	assert.Equal(t, "test_key_auth", user.KeyAuth)
	assert.Equal(t, "test_key_p256_dh", user.KeyP256DH)
	assert.Equal(t, false, user.CreatedAt().IsZero())
}

func TestNewUserIsEnabled(t *testing.T) {
	user := createUser()
	assert.Equal(t, true, user.IsEnabled())
}

func TestDisableEnableUser(t *testing.T) {
	user := createUser()
	user.Disable()
	enabledBefore := user.IsEnabled()
	disabledAt := user.DisabledAt()
	user.Enable()
	enabledAfter := user.IsEnabled()
	assert.Equal(t, false, enabledBefore)
	assert.Equal(t, true, enabledAfter)
	assert.NotNil(t, disabledAt)
}

func TestNewUserOptedIn(t *testing.T) {
	user := createUser()
	assert.Equal(t, true, user.IsOptedIn())
}

func TestOptInOptOutUser(t *testing.T) {
	user := createUser()
	user.OptOut()
	optedOut := user.IsOptedIn()
	user.OptIn()
	optedIn := user.IsOptedIn()
	assert.Equal(t, false, optedOut)
	assert.Equal(t, true, optedIn)
}

func createUser() (user *u.User) {
	user, _ = u.NewUser(
		"test_endpoint",
		"test_remote_ip",
		"test_key_auth",
		"test_key_p256_dh")
	return
}

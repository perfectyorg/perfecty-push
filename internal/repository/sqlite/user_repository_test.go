package sqlite_test

import (
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	u "github.com/rwngallego/perfecty-push/internal/domain/user"
	"github.com/rwngallego/perfecty-push/internal/repository/sqlite"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestGetUsers(t *testing.T) {
	var (
		db       = setupDB()
		userRepo = sqlite.NewSqlLiteUserRepository(db)
		user1, _ = u.NewUser("endpoint_test_1", "127.0.0.1", "my_key_auth_1", "my_p256_dh_key_1")
		user2, _ = u.NewUser("endpoint_test_2", "127.0.0.2", "my_key_auth_2", "my_p256_dh_key_2")
	)
	defer db.Close()
	user1.Disable()
	user1.OptOut()
	userRepo.Create(user1)
	userRepo.Create(user2)

	users, err := userRepo.Get(0, 3, "endpoint", "asc", false)

	assert.Equal(t, 2, len(users))
	assert.NoError(t, err)
	assert.Equal(t, "endpoint_test_1", users[0].Endpoint)
	assert.Equal(t, "127.0.0.1", users[0].RemoteIP)
	assert.Equal(t, "my_key_auth_1", users[0].KeyAuth)
	assert.Equal(t, "my_p256_dh_key_1", users[0].KeyP256DH)
	assert.Equal(t, false, users[0].CreatedAt().IsZero())
	assert.Equal(t, false, users[0].IsOptedIn())
	assert.Equal(t, false, users[0].IsEnabled())
	assert.NotNil(t, users[0].DisabledAt())
}

func TestGetUserById(t *testing.T) {
	var (
		db             = setupDB()
		id, _          = uuid.NewRandom()
		timeCreated    = time.Now()
		timeDisabledAt = time.Now()
		userRepo       = sqlite.NewSqlLiteUserRepository(db)
		user, _        = u.NewUserRaw(id, "endpoint_test", "127.0.0.1", "my_key_auth", "my_p256_dh_key", false, false, timeCreated, &timeDisabledAt)
	)
	defer db.Close()
	userRepo.Create(user)

	created, err := userRepo.GetById(user.Uuid)

	assert.NoError(t, err)
	assert.Equal(t, "endpoint_test", created.Endpoint)
	assert.Equal(t, "127.0.0.1", created.RemoteIP)
	assert.Equal(t, "my_key_auth", created.KeyAuth)
	assert.Equal(t, "my_p256_dh_key", created.KeyP256DH)
	assert.Equal(t, false, created.IsOptedIn())
	assert.Equal(t, false, created.IsEnabled())
	assert.Equal(t, false, created.CreatedAt().IsZero())
	assert.Equal(t, false, created.DisabledAt().IsZero())
}

func TestGetUserByEndpoint(t *testing.T) {
	var (
		db       = setupDB()
		userRepo = sqlite.NewSqlLiteUserRepository(db)
		user1, _ = u.NewUser("endpoint_test_1", "127.0.0.1", "my_key_auth_1", "my_p256_dh_key_1")
		user2, _ = u.NewUser("endpoint_test_2", "127.0.0.1", "my_key_auth_2", "my_p256_dh_key_2")
	)
	defer db.Close()
	userRepo.Create(user1)
	userRepo.Create(user2)

	created, err := userRepo.GetByEndpoint(user2.Endpoint)

	assert.NoError(t, err)
	assert.Equal(t, "endpoint_test_2", created.Endpoint)
	assert.Equal(t, "127.0.0.1", created.RemoteIP)
	assert.Equal(t, "my_key_auth_2", created.KeyAuth)
	assert.Equal(t, "my_p256_dh_key_2", created.KeyP256DH)
	assert.Equal(t, false, created.CreatedAt().IsZero())
	assert.Equal(t, true, created.IsOptedIn())
	assert.Equal(t, true, created.IsEnabled())
	assert.Nil(t, created.DisabledAt())
}

func TestCreateUser(t *testing.T) {
	var (
		db       = setupDB()
		userRepo = sqlite.NewSqlLiteUserRepository(db)
		user, _  = u.NewUser("endpoint_test", "127.0.0.1", "my_key_auth", "my_p256_dh_key")
	)
	defer db.Close()
	err := userRepo.Create(user)

	created, errGetById := userRepo.GetById(user.Uuid)

	assert.NoError(t, err)
	assert.NoError(t, errGetById)
	assert.NotNil(t, created)
	assert.Equal(t, "endpoint_test", created.Endpoint)
	assert.Equal(t, "127.0.0.1", created.RemoteIP)
	assert.Equal(t, "my_key_auth", created.KeyAuth)
	assert.Equal(t, "my_p256_dh_key", created.KeyP256DH)
	assert.Equal(t, false, created.CreatedAt().IsZero())
	assert.Equal(t, true, created.IsOptedIn())
	assert.Equal(t, true, created.IsEnabled())
	assert.Nil(t, created.DisabledAt())
}

func TestUpdateUser(t *testing.T) {
	var (
		db       = setupDB()
		userRepo = sqlite.NewSqlLiteUserRepository(db)
		user, _  = u.NewUser("endpoint_test", "127.0.0.1", "my_key_auth", "my_p256_dh_key")
	)
	defer db.Close()
	errCreated := userRepo.Create(user)

	user.Endpoint = "endpoint_test_2"
	user.RemoteIP = "127.0.0.2"
	user.KeyAuth = "my_key_auth_2"
	user.KeyP256DH = "my_p256_dh_key_2"
	user.OptOut()
	user.Disable()
	err := userRepo.Update(user)
	updated, errGetById := userRepo.GetById(user.Uuid)

	assert.NoError(t, errCreated)
	assert.NoError(t, err)
	assert.NoError(t, errGetById)
	assert.NotNil(t, updated)
	assert.Equal(t, "endpoint_test_2", updated.Endpoint)
	assert.Equal(t, "127.0.0.2", updated.RemoteIP)
	assert.Equal(t, "my_key_auth_2", updated.KeyAuth)
	assert.Equal(t, "my_p256_dh_key_2", updated.KeyP256DH)
	assert.Equal(t, false, updated.CreatedAt().IsZero())
	assert.Equal(t, false, updated.IsOptedIn())
	assert.Equal(t, false, updated.IsEnabled())
	assert.Equal(t, false, updated.DisabledAt().IsZero())
}

func TestDeleteUser(t *testing.T) {
	var (
		db       = setupDB()
		userRepo = sqlite.NewSqlLiteUserRepository(db)
		user, _  = u.NewUser("endpoint_test", "127.0.0.1", "my_key_auth", "my_p256_dh_key")
	)
	defer db.Close()
	errCreate := userRepo.Create(user)
	created, errGetById := userRepo.GetById(user.Uuid)

	err := userRepo.Delete(user.Uuid)
	deleted, errGetByIdDeleted := userRepo.GetById(user.Uuid)

	assert.NoError(t, errCreate)
	assert.NotNil(t, created)
	assert.NoError(t, errGetById)
	assert.NoError(t, err)
	assert.Error(t, errGetByIdDeleted)
	assert.Nil(t, deleted)
}

func TestGetTotalUsers(t *testing.T) {
	var (
		db       = setupDB()
		userRepo = sqlite.NewSqlLiteUserRepository(db)
		user1, _ = u.NewUser("endpoint_test_1", "127.0.0.1", "my_key_auth_1", "my_p256_dh_key_1")
		user2, _ = u.NewUser("endpoint_test_2", "127.0.0.2", "my_key_auth_2", "my_p256_dh_key_2")
		user3, _ = u.NewUser("endpoint_test_3", "127.0.0.3", "my_key_auth_3", "my_p256_dh_key_3")
		user4, _ = u.NewUser("endpoint_test_4", "127.0.0.4", "my_key_auth_4", "my_p256_dh_key_4")
	)
	defer db.Close()

	user1.OptOut()
	user2.Disable()
	user3.OptOut()
	user3.Disable()

	userRepo.Create(user1)
	userRepo.Create(user2)
	userRepo.Create(user3)
	userRepo.Create(user4)

	total, err := userRepo.GetTotal(true)

	assert.NoError(t, err)
	assert.Equal(t, 1, total)
}

func TestGetUsersStats(t *testing.T) {
	var (
		db       = setupDB()
		userRepo = sqlite.NewSqlLiteUserRepository(db)
		user1, _ = u.NewUser("endpoint_test_1", "127.0.0.1", "my_key_auth_1", "my_p256_dh_key_1")
		user2, _ = u.NewUser("endpoint_test_2", "127.0.0.2", "my_key_auth_2", "my_p256_dh_key_2")
		user3, _ = u.NewUser("endpoint_test_3", "127.0.0.3", "my_key_auth_3", "my_p256_dh_key_3")
		user4, _ = u.NewUser("endpoint_test_4", "127.0.0.4", "my_key_auth_4", "my_p256_dh_key_4")
	)
	defer db.Close()

	user1.OptOut()
	user2.Disable()
	user3.OptOut()
	user3.Disable()

	userRepo.Create(user1)
	userRepo.Create(user2)
	userRepo.Create(user3)
	userRepo.Create(user4)

	total, active, inactive, err := userRepo.Stats()

	assert.NoError(t, err)
	assert.Equal(t, 4, total)
	assert.Equal(t, 1, active)
	assert.Equal(t, 3, inactive)
}

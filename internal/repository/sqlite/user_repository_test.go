package sqlite_test

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	u "github.com/rwngallego/perfecty-push/internal/domain/user"
	"github.com/rwngallego/perfecty-push/internal/repository/sqlite"
	"github.com/stretchr/testify/assert"
	"testing"
)

const dsn = "file:test.db?mode=memory"
const driver = "sqlite3"

func TestGetUsers(t *testing.T) {
	var (
		db       = setupDB()
		userRepo = sqlite.NewSqlLiteUserRepository(db)
		_, _     = u.NewUser("endpoint_test_1", "127.0.0.1", "my_key_auth_1", "my_p256_dh_key_1")
		_, _     = u.NewUser("endpoint_test_2", "127.0.0.2", "my_key_auth_2", "my_p256_dh_key_2")
	)
	defer db.Close()

	users, err := userRepo.Get(0, 3, "endpoint", "asc", false)

	assert.Equal(t, 2, len(users))
	assert.NoError(t, err)
}

func TestCreateUser(t *testing.T) {
	var (
		db       = setupDB()
		userRepo = sqlite.NewSqlLiteUserRepository(db)
		user, _  = u.NewUser("endpoint_test", "127.0.0.1", "my_key_auth", "my_p256_dh_key")
	)
	defer db.Close()

	err := userRepo.Create(user)
	created, _ := userRepo.GetById(user.Uuid)

	assert.NoError(t, err)
	assert.NotNil(t, created)
}

func TestGetUser(t *testing.T) {
	var (
		db       = setupDB()
		userRepo = sqlite.NewSqlLiteUserRepository(db)
		user, _  = u.NewUser("endpoint_test", "127.0.0.1", "my_key_auth", "my_p256_dh_key")
	)
	defer db.Close()
	userRepo.Create(user)

	created, err := userRepo.GetById(user.Uuid)

	assert.NoError(t, err)
	assert.Equal(t, "endpoint_test", created.Endpoint)
	assert.Equal(t, "127.0.0.1", created.RemoteIP)
	assert.Equal(t, "my_key_auth", created.KeyAuth)
	assert.Equal(t, "my_p256_dh_key", created.KeyP256DH)
	assert.Equal(t, false, created.CreatedAt().IsZero())
	assert.Equal(t, true, created.IsOptedIn())
	assert.Equal(t, true, created.IsEnabled())
	assert.Equal(t, nil, created.DisabledAt())
}

func setupDB() (db *sql.DB) {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		panic(err)
	}
	return
}

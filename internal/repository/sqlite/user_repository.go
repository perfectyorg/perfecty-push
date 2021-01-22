package sqlite

import (
	"database/sql"
	"github.com/google/uuid"
	u "github.com/rwngallego/perfecty-push/internal/domain/user"
)

func NewSqlLiteUserRepository(db *sql.DB) (r *SqlLiteUserRepository) {
	r = &SqlLiteUserRepository{
		db: db,
	}
	return
}

type SqlLiteUserRepository struct {
	db *sql.DB
}

func (r *SqlLiteUserRepository) Get(offset int, size int, orderBy string, orderAsc string, onlyActive bool) (userList []u.User, err error) {
	return
}

func (r *SqlLiteUserRepository) GetById(id uuid.UUID) (user *u.User, err error) {
	return
}

func (r *SqlLiteUserRepository) GetByEndpoint(endpoint string) (user *u.User, err error) {
	return
}

func (r *SqlLiteUserRepository) Create(user *u.User) (err error) {
	return
}

func (r *SqlLiteUserRepository) Update(user *u.User) (err error) {
	return
}

func (r *SqlLiteUserRepository) Delete(id uuid.UUID) (err error) {
	return
}

func (r *SqlLiteUserRepository) GetTotal() (total int, err error) {
	return
}

func (r *SqlLiteUserRepository) Stats() (total int, active int, inactive int) {
	return
}

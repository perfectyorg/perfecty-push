package sqlite

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	u "github.com/rwngallego/perfecty-push/internal/domain/user"
	"time"
)

const (
	fields = "uuid,endpoint,remote_ip,key_auth,key_p256dh,opted_in,enabled,created_at,disabled_at"
)

func NewSqlLiteUserRepository(db *sql.DB) (r *SqlLiteUserRepository) {
	r = &SqlLiteUserRepository{
		db: db,
	}
	return
}

type (
	// We need it to have a general accessor to *sql.Row and *sql.Rows
	rowInterface interface {
		Scan(dest ...interface{}) error
	}

	SqlLiteUserRepository struct {
		db *sql.DB
	}
)

func (r *SqlLiteUserRepository) Get(offset int, size int, orderBy string, orderAsc string, onlyActive bool) (userList []*u.User, err error) {
	stmt, err := r.db.Prepare("SELECT " + fields + " FROM users ORDER BY " + orderBy + " " + orderAsc + " LIMIT ? OFFSET ?")
	if err != nil {
		log.Error().Err(err).Msg("Could not prepare the query")
		return
	}
	defer stmt.Close()

	rows, err := stmt.Query(size, offset)
	if err != nil {
		log.Error().Err(err).Msg("Could not run the query")
		return
	}

	for rows.Next() {
		var user *u.User
		user, err = getUserFromRow(rows)
		if err != nil {
			log.Error().Err(err).Msg("Could not read one of the rows")
			return
		}
		userList = append(userList, user)
	}
	return
}

func (r *SqlLiteUserRepository) GetById(id uuid.UUID) (user *u.User, err error) {
	stmt, err := r.db.Prepare("SELECT " + fields + " FROM users WHERE uuid = ?")
	if err != nil {
		log.Error().Err(err).Msg("Could not prepare the query")
		return
	}
	defer stmt.Close()

	row := stmt.QueryRow(id.String())
	return getUserFromRow(row)
}

func (r *SqlLiteUserRepository) GetByEndpoint(endpoint string) (user *u.User, err error) {
	return
}

func (r *SqlLiteUserRepository) Create(user *u.User) (err error) {
	stmt, err := r.db.Prepare("INSERT INTO users(uuid, endpoint, remote_ip, key_auth, key_p256dh) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		log.Error().Err(err).Msg("Could not prepare the query")
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Uuid, user.Endpoint, user.RemoteIP, user.KeyAuth, user.KeyP256DH)
	if err != nil {
		log.Error().Err(err).Msg("Could not create the user")
		return
	}
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

func getUserFromRow(row rowInterface) (user *u.User, err error) {
	var (
		uuidString string
		endpoint   string
		remoteIP   string
		keyAuth    string
		keyP256DH  string
		optedIn    bool
		enabled    bool
		createdAt  time.Time
		disabledAt *time.Time
	)

	err = row.Scan(&uuidString, &endpoint, &remoteIP, &keyAuth, &keyP256DH, &optedIn, &enabled, &createdAt, &disabledAt)
	if err != nil {
		log.Error().Err(err).Msg("Could not parse the user")
		return
	}

	uuid, err := uuid.Parse(uuidString)
	if err != nil {
		log.Error().Err(err).Msg("Could not parse the uuid")
	}

	return u.NewUserRaw(uuid, endpoint, remoteIP, keyAuth, keyP256DH, optedIn, enabled, createdAt, disabledAt)
}

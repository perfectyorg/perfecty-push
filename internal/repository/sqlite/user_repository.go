package sqlite

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	u "github.com/rwngallego/perfecty-push/internal/domain/user"
	"time"
)

const (
	userFields = "uuid,endpoint,remote_ip,key_auth,key_p256dh,opted_in,enabled,created_at,disabled_at"
)

func NewSqlLiteUserRepository(db *sql.DB) (r *SqlLiteUserRepository) {
	r = &SqlLiteUserRepository{
		db: db,
	}
	return
}

type (
	SqlLiteUserRepository struct {
		db *sql.DB
	}
)

func (r *SqlLiteUserRepository) Get(offset int, size int, orderBy string, orderAsc string, onlyActive bool) (userList []*u.User, err error) {
	stmt, err := r.db.Prepare("SELECT " + userFields + " FROM users ORDER BY " + orderBy + " " + orderAsc + " LIMIT ? OFFSET ?")
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
	stmt, err := r.db.Prepare("SELECT " + userFields + " FROM users WHERE uuid = ?")
	if err != nil {
		log.Error().Err(err).Msg("Could not prepare the query")
		return
	}
	defer stmt.Close()

	row := stmt.QueryRow(id.String())
	return getUserFromRow(row)
}

func (r *SqlLiteUserRepository) GetByEndpoint(endpoint string) (user *u.User, err error) {
	stmt, err := r.db.Prepare("SELECT " + userFields + " FROM users WHERE endpoint = ?")
	if err != nil {
		log.Error().Err(err).Msg("Could not prepare the query")
		return
	}
	defer stmt.Close()

	row := stmt.QueryRow(endpoint)
	return getUserFromRow(row)
}

func (r *SqlLiteUserRepository) Create(user *u.User) (err error) {
	stmt, err := r.db.Prepare("INSERT INTO users(uuid, endpoint, remote_ip, key_auth, key_p256dh, enabled, opted_in, created_at, disabled_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Error().Err(err).Msg("Could not prepare the query")
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Uuid, user.Endpoint, user.RemoteIP, user.KeyAuth, user.KeyP256DH, user.IsEnabled(), user.IsOptedIn(), user.CreatedAt(), user.DisabledAt())
	if err != nil {
		log.Error().Err(err).Msg("Could not create the user")
		return
	}
	return
}

func (r *SqlLiteUserRepository) Update(user *u.User) (err error) {
	stmt, err := r.db.Prepare("UPDATE users SET endpoint=?, remote_ip=?, key_auth=?, key_p256dh=?, opted_in=?, enabled=?, created_at=?, disabled_at=?")
	if err != nil {
		log.Error().Err(err).Msg("Could not prepare the query")
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Endpoint, user.RemoteIP, user.KeyAuth, user.KeyP256DH, user.IsOptedIn(), user.IsEnabled(), user.CreatedAt(), user.DisabledAt())
	if err != nil {
		log.Error().Err(err).Msg("Could not update the user")
		return
	}
	return
}

func (r *SqlLiteUserRepository) Delete(id uuid.UUID) (err error) {
	stmt, err := r.db.Prepare("DELETE FROM users WHERE uuid = ?")
	if err != nil {
		log.Error().Err(err).Msg("Could not prepare the query")
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(id.String())
	if err != nil {
		log.Error().Err(err).Msg("Could not delete the user")
		return
	}
	return
}

func (r *SqlLiteUserRepository) GetTotal(onlyActive bool) (total int, err error) {
	var where string

	if onlyActive {
		where = " WHERE enabled=1 and opted_in=1"
	}

	if err = r.db.QueryRow("SELECT count(*) FROM users " + where).Scan(&total); err != nil {
		log.Error().Err(err).Msg("Could not get the total users")
		return
	}

	return
}

func (r *SqlLiteUserRepository) Stats() (total int, active int, inactive int, err error) {
	if err = r.db.QueryRow("SELECT count(*) FROM users ").Scan(&total); err != nil {
		log.Error().Err(err).Msg("Could not get the total users")
		return
	}

	if err = r.db.QueryRow("SELECT count(*) FROM users WHERE enabled=1 AND opted_in=1").Scan(&active); err != nil {
		log.Error().Err(err).Msg("Could not get the total active users")
		return
	}
	inactive = total - active
	return
}

// Private

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
		log.Error().Err(err).Msg("Could not scan the user")
		return
	}

	id, err := uuid.Parse(uuidString)
	if err != nil {
		log.Error().Err(err).Msg("Could not parse the uuid")
	}

	return u.NewUserRaw(id, endpoint, remoteIP, keyAuth, keyP256DH, optedIn, enabled, createdAt, disabledAt)
}

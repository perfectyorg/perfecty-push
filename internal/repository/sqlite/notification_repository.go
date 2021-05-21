package sqlite

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	n "github.com/perfectyorg/perfecty-push/internal/domain/notification"
	"time"
)

const (
	notificationFields = "uuid,payload,total,succeeded,last_cursor,batch_size,status,is_taken,created_at,status_changed_at"
)

func NewSqlLiteNotificationRepository(db *sql.DB) (r *SqlLiteNotificationRepository) {
	r = &SqlLiteNotificationRepository{
		db: db,
	}
	return
}

type (
	SqlLiteNotificationRepository struct {
		db *sql.DB
	}
)

func (r *SqlLiteNotificationRepository) Get(offset int, size int, orderBy string, orderAsc string) (notificationList []*n.Notification, err error) {
	stmt, err := r.db.Prepare("SELECT " + notificationFields + " FROM notifications ORDER BY " + orderBy + " " + orderAsc + " LIMIT ? OFFSET ?")
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
		var notification *n.Notification
		notification, err = getNotificationFromRow(rows)
		if err != nil {
			log.Error().Err(err).Msg("Could not read one of the rows")
			return
		}
		notificationList = append(notificationList, notification)
	}
	return
}

func (r *SqlLiteNotificationRepository) GetById(id uuid.UUID) (notification *n.Notification, err error) {
	stmt, err := r.db.Prepare("SELECT " + notificationFields + " FROM notifications WHERE uuid = ?")
	if err != nil {
		log.Error().Err(err).Msg("Could not prepare the query")
		return
	}
	defer stmt.Close()

	row := stmt.QueryRow(id.String())
	return getNotificationFromRow(row)
}

func (r *SqlLiteNotificationRepository) Create(notification *n.Notification) (err error) {
	stmt, err := r.db.Prepare("INSERT INTO notifications(uuid, payload, total, succeeded, last_cursor, batch_size, status, is_taken, created_at, status_changed_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Error().Err(err).Msg("Could not prepare the query")
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(notification.Uuid, notification.Payload, notification.Total, notification.Succeeded, notification.LastCursor, notification.BatchSize, notification.Status(), notification.IsTaken(), notification.CreatedAt(), notification.StatusChangedAt())
	if err != nil {
		log.Error().Err(err).Msg("Could not create the notification")
		return
	}
	return
}

func (r *SqlLiteNotificationRepository) Update(notification *n.Notification) (err error) {
	stmt, err := r.db.Prepare("UPDATE notifications SET payload=?, total=?, succeeded=?, last_cursor=?, batch_size=?, status=?, is_taken=?, created_at=?, status_changed_at=?")
	if err != nil {
		log.Error().Err(err).Msg("Could not prepare the query")
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(notification.Payload, notification.Total, notification.Succeeded, notification.LastCursor, notification.BatchSize, notification.Status(), notification.IsTaken(), notification.CreatedAt(), notification.StatusChangedAt())
	if err != nil {
		log.Error().Err(err).Msg("Could not update the notification")
		return
	}
	return
}

func (r *SqlLiteNotificationRepository) Delete(id uuid.UUID) (err error) {
	stmt, err := r.db.Prepare("DELETE FROM notifications WHERE uuid = ?")
	if err != nil {
		log.Error().Err(err).Msg("Could not prepare the query")
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(id.String())
	if err != nil {
		log.Error().Err(err).Msg("Could not delete the notification")
		return
	}
	return
}

func (r *SqlLiteNotificationRepository) Stats() (total int, succeeded int, failed int, err error) {
	stmt, err := r.db.Prepare("SELECT SUM(total), SUM(succeeded) FROM notifications WHERE status != ?")
	if err != nil {
		log.Error().Err(err).Msg("Could not prepare the query")
		return
	}
	defer stmt.Close()

	if err = stmt.QueryRow(n.StatusRunning).Scan(&total, &succeeded); err != nil {
		log.Error().Err(err).Msg("Could not get the notifications stats")
		return
	}
	failed = total - succeeded
	return
}

// Internal

func getNotificationFromRow(row rowInterface) (notification *n.Notification, err error) {
	var (
		uuidString      string
		payload         string
		Total           int
		succeeded       int
		lastCursor      int
		batchSize       int
		status          int
		isTaken         bool
		createdAt       time.Time
		statusChangedAt time.Time
	)

	err = row.Scan(&uuidString, &payload, &Total, &succeeded, &lastCursor, &batchSize, &status, &isTaken, &createdAt, &statusChangedAt)
	if err != nil {
		log.Error().Err(err).Msg("Could not scan the notification")
		return
	}

	id, err := uuid.Parse(uuidString)
	if err != nil {
		log.Error().Err(err).Msg("Could not parse the uuid")
	}

	return n.NewNotificationRaw(id, payload, Total, succeeded, lastCursor, batchSize, status, isTaken, createdAt, statusChangedAt)
}

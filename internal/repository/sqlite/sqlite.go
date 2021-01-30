package sqlite

type (
	// We need it to have a general accessor to *sql.Row and *sql.Rows
	rowInterface interface {
		Scan(dest ...interface{}) error
	}
)

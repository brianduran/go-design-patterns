package repository

import (
	"database/sql"

	"github.com/ltvco/data-eng-go-lib/sqlutil"
)

type User struct {
	Age  int
	Name string
}

// Scan maps the SQL results into the particular tiktokUserToAdd object.
func (U *User) Scan(rows *sql.Rows) error {
	return sqlutil.Scan(rows, U)
}

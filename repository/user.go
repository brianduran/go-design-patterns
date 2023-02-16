package repository

import (
	"database/sql"

	"github.com/ltvco/data-eng-go-lib/sqlutil"
)

// User contains the relevant data to handle an user.
type User struct {
	Age  int
	Name string
}

// Scan maps the SQL results into the particular User object.
func (u *User) Scan(rows *sql.Rows) error {
	return sqlutil.Scan(rows, u)
}

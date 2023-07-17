package repository

import (
	"database/sql"

	"github.com/ltvco/data-eng-go-lib/sqlutil"
)

// ErrUserNotFound is the error returned when a user is not found.
const ErrUserNotFound = "user not found"

// User contains the relevant data to handle a user.
type User struct {
	Age  int    `sql:"age"`
	ID   int    `sql:"id"`
	Name string `sql:"name"`
}

// Scan maps the SQL results into the particular User object.
func (u *User) Scan(rows *sql.Rows) error {
	return sqlutil.Scan(rows, u)
}

package mysql

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

// Open establishes a connection with MySQL.
func Open(config Config) (*sql.DB, error) {
	conf := mysql.Config{
		Net:                  "tcp",
		Addr:                 config.Host,
		User:                 config.User,
		Passwd:               config.Password,
		DBName:               config.Database,
		ParseTime:            true,
		AllowNativePasswords: true,
	}
	fmt.Printf("connection: %v\n", conf.FormatDSN())
	db, err := sql.Open("mysql", conf.FormatDSN())
	if err != nil {
		return nil, errors.New("failed to open MySQL connection")
	}

	if err = db.Ping(); err != nil {
		return nil, errors.New("failed to ping MySQL")
	}

	return db, nil
}

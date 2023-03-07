package repository

import (
	"github.com/ltvco/go-design-patterns/dbconn/mysql"
)

// Config contains all the required configurations for the project.
type Config struct {
	MysqlDB mysql.Config
}

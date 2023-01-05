package repository

import (
	"github.com/ltvco/go_design_patterns/dbconn/mysql"
)

var conf = config{
	mysqlDB: mysql.Config{
		Database: "repository",
		Host:     "localhost",
		Password: "",
		User:     "root",
	},
}

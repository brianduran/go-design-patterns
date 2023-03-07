package main

import (
	"github.com/ltvco/go-design-patterns/dbconn/mysql"
	"github.com/ltvco/go-design-patterns/repository"
)

var conf = repository.Config{
	MysqlDB: mysql.Config{
		Database: "app",
		Host:     "localhost:3307",
		Password: "password",
		User:     "root",
	},
}

package repository

import (
	"context"
	"database/sql"

	"github.com/ltvco/go_design_patterns/dbconn/mysql"

	// required to use go embed feature
	_ "embed"
)

var (
	//go:embed sql/mysql_repo_insert_user_stmt.sql
	insertUserStmt string

	//go:embed sql/mysql_repo_select_user_by_id.sql
	selectUserStmt string

	//go:embed sql/mysql_repo_update_user_by_id.sql
	updateUserStmt string
)

type MysqlRepository struct {
	db *sql.DB
}

func NewMysqlRepository() (*MysqlRepository, error) {
	db, err := mysql.Open(conf.mysqlDB)
	if err != nil {
		return nil, err
	}

	return &MysqlRepository{
		db: db,
	}, nil
}

func (mr *MysqlRepository) CreateUser(ctx context.Context, name string, age int) error {
	_, err := mr.db.ExecContext(ctx, insertUserStmt, name, age)
	return err
}

func (mr *MysqlRepository) GetUser(ctx context.Context, id int) (*User, error) {
	row, err := mr.db.QueryContext(ctx, selectUserStmt, id)
	if err != nil {
		return nil, err
	}

	row.Next()
	user := new(User)

	err = user.Scan(row)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (mr *MysqlRepository) UpdateUser(
	ctx context.Context,
	id int,
	attributes []interface{},
) error {
	attributes = append(attributes, id)
	_, err := mr.db.ExecContext(ctx, updateUserStmt, attributes...)
	return err
}

func (*MysqlRepository) DeleteUser(id int) error {
	return nil
}

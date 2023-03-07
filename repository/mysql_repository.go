package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/ltvco/go-design-patterns/dbconn/mysql"

	// required to use go embed feature
	_ "embed"
)

var (
	//go:embed sql/mysql_repo_delete_user_by_id.sql
	deleteUserStmt string

	//go:embed sql/mysql_repo_insert_user.sql
	insertUserStmt string

	//go:embed sql/mysql_repo_select_user_by_id.sql
	selectUserStmt string

	//go:embed sql/mysql_repo_update_user_by_id.sql
	updateUserStmt string
)

// MysqlRepository contains the methods to handle a user stored in a MySQL database.
type MysqlRepository struct {
	db *sql.DB
}

// NewMysqlRepository creates a new NewMysqlRepository.
func NewMysqlRepository(config mysql.Config) (*MysqlRepository, error) {
	db, err := mysql.Open(config)
	if err != nil {
		return nil, err
	}

	return &MysqlRepository{
		db: db,
	}, nil
}

// CreateUser executes the SQL statement to create a user.
func (mr *MysqlRepository) CreateUser(ctx context.Context, name string, age int) error {
	_, err := mr.db.ExecContext(ctx, insertUserStmt, name, age)
	return err
}

// DeleteUser executes the SQL statement to delete a user.
func (mr *MysqlRepository) DeleteUser(ctx context.Context, id int) error {
	_, err := mr.db.ExecContext(ctx, deleteUserStmt, id)
	return err
}

// GetUser executes the SQL statement to retrieve a user's data.
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

// UpdateUser executes the SQL statement to update a user.
func (mr *MysqlRepository) UpdateUser(
	ctx context.Context,
	id int,
	attributes map[string]interface{},
) error {
	var values []interface{}
	var setAttributes []string
	for key, value := range attributes {
		setAttributes = append(setAttributes, fmt.Sprintf("`%s` = ?", key))
		values = append(values, value)
	}
	values = append(values, id)
	setStatement := fmt.Sprintf("SET %s", strings.Join(setAttributes, ", "))

	query := fmt.Sprintf(updateUserStmt, setStatement)
	_, err := mr.db.ExecContext(ctx, query, values...)
	return err
}

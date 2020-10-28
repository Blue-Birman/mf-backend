package customers

import (
	"github.com/rvalessandro/go-bookstore_users-api/datasources/mysql"
	"github.com/rvalessandro/go-bookstore_users-api/utils/errors"
	"github.com/rvalessandro/go-bookstore_users-api/utils/mysql_util"
)

const (
	queryGetCustomer = "SELECT id, name, email, address, created_at, updated_at FROM users WHERE id=?"
)

var (
	usersDB = make(map[int64]*Customer)
)

func (user *Customer) Get() *errors.RestErr {
	stmt, err := mysql.Client.Prepare(queryGetCustomer)
	if err != nil {
		return errors.NewErrInternalServer(err.Error())
	}
	defer stmt.Close()

	result := stmt.QueryRow(user.ID)
	if getErr := result.Scan(&user.ID, &user.Name, &user.Email, &user.Address, &user.CreatedAt, &user.UpdatedAt); getErr != nil {
		return mysql_util.ParseError(getErr)
	}

	return nil
}

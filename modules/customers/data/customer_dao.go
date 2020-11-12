package customers

import (
	"github.com/rvalessandro/mf-backend/datasources/mysql"
	"github.com/rvalessandro/mf-backend/utils/errors"
	"github.com/rvalessandro/mf-backend/utils/mysql_util"
)

const (
	queryGetCustomer = "SELECT id, name, email, address, created_at, updated_at FROM users WHERE id=?"
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

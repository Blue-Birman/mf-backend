package mysql_util

import (
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/rvalessandro/mf-backend/utils/errors"
)

const (
	errorNoRows = "no rows in result set"
)

func ParseError(err error) *errors.RestErr {
	sqlErr, ok := err.(*mysql.MySQLError)
	if !ok {
		if strings.Contains(err.Error(), errorNoRows) {
			return errors.NewErrNotFound("no record matching given id")
		}
		return errors.NewErrInternalServer("error parsing database response")
	}

	switch sqlErr.Number {
	case 1062:
		return errors.NewErrBadRequest("invalid data")
	}

	return errors.NewErrInternalServer("error processing request")
}

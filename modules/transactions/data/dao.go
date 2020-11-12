package data

import (
	"github.com/rvalessandro/mf-backend/modules/transactions/domain"
	"github.com/rvalessandro/mf-backend/utils/errors"
)

func Get() (*domain.Transaction, *errors.RestErr) {
	return nil, errors.NewErrInternalServer("Method not yet implemented")
}

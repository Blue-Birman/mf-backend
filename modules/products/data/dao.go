package data

import (
	"github.com/rvalessandro/mf-backend/modules/products/domain"
	"github.com/rvalessandro/mf-backend/utils/errors"
)

func Get() (*domain.Product, *errors.RestErr) {
	return nil, errors.NewErrInternalServer("Method not yet implemented")
}

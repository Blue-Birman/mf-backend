package products

import "github.com/rvalessandro/mf-backend/utils/errors"

func (product *Product) Get() *errors.RestErr {
	return errors.NewErrInternalServer("Method not yet implemented")
}

package transactions

import "github.com/rvalessandro/mf-backend/utils/errors"

func (transaction *Transaction) Get() *errors.RestErr {
	return errors.NewErrInternalServer("Method not yet implemented")
}

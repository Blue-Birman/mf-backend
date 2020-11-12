package services

import (
	"github.com/rvalessandro/mf-backend/modules/transactions/data"
	"github.com/rvalessandro/mf-backend/modules/transactions/domain"
	"github.com/rvalessandro/mf-backend/utils/errors"
)

func GetTransaction(transactionID int64) (*domain.Transaction, *errors.RestErr) {
	transaction, err := data.Get()
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

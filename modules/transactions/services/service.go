package services

import (
	"github.com/rvalessandro/mf-backend/modules/transactions/data"
	"github.com/rvalessandro/mf-backend/modules/transactions/domain"
	"github.com/rvalessandro/mf-backend/utils/errors"
	"time"
)

func FindTransaction() ([]domain.Transaction, *errors.RestErr) {
	transactions, err := data.Find()
	if err != nil {
		return nil, err
	}

	return transactions, nil
}

func GetTransaction(transactionID int64) (*domain.Transaction, *errors.RestErr) {
	transaction, err := data.Get(transactionID)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func CreateTransaction(transactionParam domain.CreateTransactionParams) (*domain.Transaction, *errors.RestErr) {
	currentTime := time.Now()
	transactionParam.CreatedAt = currentTime
	transactionParam.UpdatedAt = currentTime

	transaction, err := data.Create(transactionParam)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func UpdateTransaction(id int64, transactionParam domain.CreateTransactionParams) (*domain.Transaction, *errors.RestErr) {
	transactionParam.UpdatedAt = time.Now()

	transaction, err := data.Update(id, transactionParam)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func DeleteTransaction(transactionID int64) (*domain.Transaction, *errors.RestErr) {
	transaction, err := data.Get(transactionID)
	if err != nil {
		return nil, err
	}

	err = data.Delete(transactionID)
	if err != nil {
		return nil, err
	}

	return transaction, err
}

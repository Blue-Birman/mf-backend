package domain

import (
	"github.com/rvalessandro/mf-backend/modules/products/domain"
	"time"
)

type Transaction struct {
	ID         int64            `json:"id"`
	CustomerID int64            `json:"customer_id"`
	Date       time.Time        `json:"date"`
	Products   []domain.Product `json:"products"`
	CreatedAt  time.Time        `json:"created_at"`
	UpdatedAt  time.Time        `json:"updated_at"`
}

type CreateTransactionParams struct {
	CustomerID int64                      `json:"customer_id"`
	Date       time.Time                  `json:"date"`
	Products   []TransactionProductParams `json:"products"`
	CreatedAt  time.Time                  `json:"created_at"`
	UpdatedAt  time.Time                  `json:"updated_at"`
}

type TransactionProductParams struct {
	ProductID     int64 `json:"product_id"`
	TransactionID int64 `json:"transaction_id"`
	Qty           int64 `json:"qty"`
}

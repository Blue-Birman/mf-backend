package domain

import (
	"github.com/rvalessandro/mf-backend/modules/products/domain"
	"time"
)

type Transaction struct {
	ID         int64                       `json:"id"`
	CustomerID int64                       `json:"customer_id"`
	Date       time.Time                   `json:"date"`
	Products   []domain.TransactionProduct `json:"products"`
	Total      int64                       `json:"total"`
	CreatedAt  time.Time                   `json:"created_at"`
	UpdatedAt  time.Time                   `json:"updated_at"`
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

func (t *Transaction) CalculateTotal() {
	var total int64
	total = 0

	for _, product := range t.Products {
		total = total + product.Total
	}

	t.Total = total
}

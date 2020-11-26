package domain

import (
	"time"
)

type Transaction struct {
	ID         int64     `json:"id"`
	CustomerID int64     `json:"customer_id"`
	Date       time.Time `json:"date"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type CreateTransactionParams struct {
	CustomerID int64     `json:"id"`
	Date       time.Time `json:"date"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

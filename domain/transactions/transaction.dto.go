package transactions

import (
	"time"
)

type Transaction struct {
	ID           int64     `json:"id"`
	CustomerID   int64     `json:"customer_id"`
	Date         time.Time `json:"date"`
	TotalPayment int64     `json:"total_payment"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

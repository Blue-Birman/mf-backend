package domain

import (
	"time"
)

type Cart struct {
	CustomerID int64     `json:"customer_id"`
	ProductID  int64     `json:"product_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type CreateCartParams struct {
	CustomerID int64     `json:"customer_id"`
	ProductID  int64     `json:"product_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func NewCart() *CreateCartParams {
	currentTime := time.Now()
	return &CreateCartParams{
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}
}

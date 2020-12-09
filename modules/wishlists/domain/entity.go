package domain

import (
	"time"
)

type Wishlist struct {
	CustomerID int64     `json:"customer_id"`
	ProductID  int64     `json:"product_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type CreateWishlistParams struct {
	CustomerID int64     `json:"customer_id"`
	ProductID  int64     `json:"product_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func NewWishlist() *CreateWishlistParams {
	currentTime := time.Now()
	return &CreateWishlistParams{
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}
}

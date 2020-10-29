package wishlists

import (
	"time"
)

type Wishlist struct {
	ID         int64     `json:"id"`
	CustomerID int64     `json:"customer_id"`
	ProductID  int64     `json:"product_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

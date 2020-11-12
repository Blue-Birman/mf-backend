package domain

import (
	"time"
)

type Product struct {
	ID        int64     `json:"id"`
	Name      int64     `json:"name"`
	ImageURL  time.Time `json:"image_url"`
	Price     int64     `json:"price"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

package domain

import (
	"database/sql"
	"time"
)

type Product struct {
	ID        int64        `json:"id"`
	Name      int64        `json:"name"`
	ImageURL  string       `json:"image_url"`
	Price     int64        `json:"price"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

type CreateProductParams struct {
	Name      int64        `json:"name"`
	ImageURL  string       `json:"image_url"`
	Price     int64        `json:"price"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

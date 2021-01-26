package domain

import (
	"time"
)

type Product struct {
	ID          int64     `json:"id"`
	CategoryID  int64     `json:"category_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ImageURL    string    `json:"image_url"`
	PreviewURL  string    `json:"preview_url"`
	Price       int64     `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateProductParams struct {
	Name        string    `json:"name"`
	CategoryID  int64     `json:"category_id"`
	Description string    `json:"description"`
	ImageURL    string    `json:"image_url"`
	PreviewURL  string    `json:"preview_url"`
	Price       int64     `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TransactionProduct struct {
	ID          int64     `json:"id"`
	CategoryID  int64     `json:"category_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ImageURL    string    `json:"image_url"`
	PreviewURL  string    `json:"preview_url"`
	Price       int64     `json:"price"`
	Qty         int8      `json:"qty"`
	Total       int64     `json:"total"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

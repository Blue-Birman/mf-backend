package domain

import (
	"time"
)

type Category struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateCategoryParams struct {
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewCategory() *CreateCategoryParams {
	currentTime := time.Now()
	return &CreateCategoryParams{
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}
}

package domain

import (
	"database/sql"
	"time"
)

type Customer struct {
	ID        int64        `json:"id"`
	Name      string       `json:"name"`
	Email     string       `json:"email"`
	Password  string       `json:"password"`
	Address   string       `json:"address"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

type CreateCustomerParams struct {
	Name      string       `json:"name"`
	Email     string       `json:"email"`
	Password  string       `json:"password"`
	Address   string       `json:"address"`
	CreatedAt time.Time    `json:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at"`
}

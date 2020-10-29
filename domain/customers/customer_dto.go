package customers

import (
	"strings"
	"time"

	"github.com/rvalessandro/mf-backend/utils/errors"
)

type Customer struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (customer *Customer) Validate() *errors.RestErr {
	customer.Name = strings.TrimSpace(strings.ToLower(customer.Name))

	customer.Email = strings.TrimSpace(strings.ToLower(customer.Email))
	if customer.Email == "" {
		return errors.NewErrBadRequest("Email must not be empty")
	}

	customer.Password = strings.TrimSpace(customer.Password)
	if customer.Password == "" {
		return errors.NewErrBadRequest("Password must not be empty")
	}

	return nil
}

package customers

import (
	"strings"

	"github.com/rvalessandro/go-bookstore_users-api/utils/errors"
)

type Customer struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Address   string `json:"address"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (user *Customer) Validate() *errors.RestErr {
	user.Name = strings.TrimSpace(strings.ToLower(user.Name))

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewErrBadRequest("Email must not be empty")
	}

	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return errors.NewErrBadRequest("Password must not be empty")
	}

	return nil
}

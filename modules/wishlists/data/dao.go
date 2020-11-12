package data

import (
	"github.com/rvalessandro/mf-backend/modules/wishlists/domain"
	"github.com/rvalessandro/mf-backend/utils/errors"
)

func Get() (*domain.Wishlist, *errors.RestErr) {
	return nil, errors.NewErrInternalServer("Method not yet implemented")
}

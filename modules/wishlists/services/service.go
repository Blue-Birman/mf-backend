package services

import (
	"github.com/rvalessandro/mf-backend/modules/wishlists/data"
	"github.com/rvalessandro/mf-backend/modules/wishlists/domain"
	"github.com/rvalessandro/mf-backend/utils/errors"
)

func GetWishlist(wishlistID int64) (*domain.Wishlist, *errors.RestErr) {
	wishlist, err := data.Get()
	if err != nil {
		return nil, err
	}
	return wishlist, nil
}

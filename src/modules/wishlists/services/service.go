package services

import (
	"time"

	productDomain "github.com/rvalessandro/mf-backend/modules/products/domain"
	"github.com/rvalessandro/mf-backend/modules/wishlists/data"
	"github.com/rvalessandro/mf-backend/modules/wishlists/domain"
	"github.com/rvalessandro/mf-backend/utils/errors"
)

func GetWishlistByCustomerID(customerID int64) ([]productDomain.Product, *errors.RestErr) {
	wishlists, err := data.Find(customerID)
	if err != nil {
		return nil, err
	}

	return wishlists, nil
}

func CreateWishlist(wishlistParam domain.CreateWishlistParams) (*domain.Wishlist, *errors.RestErr) {
	currentTime := time.Now()
	wishlistParam.CreatedAt = currentTime
	wishlistParam.UpdatedAt = currentTime

	wishlist, err := data.Create(wishlistParam)
	if err != nil {
		return nil, err
	}

	return wishlist, nil
}

func DeleteWishlist(customerID int64, productID int64) (*domain.Wishlist, *errors.RestErr) {
	wishlist, err := data.Get(customerID, productID)
	if err != nil {
		return nil, err
	}

	err = data.Delete(customerID, productID)
	if err != nil {
		return nil, err
	}

	return wishlist, err
}

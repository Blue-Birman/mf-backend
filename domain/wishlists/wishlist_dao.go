package wishlists

import "github.com/rvalessandro/mf-backend/utils/errors"

func (wishlist *Wishlist) Get() *errors.RestErr {
	return errors.NewErrInternalServer("Method not yet implemented")
}

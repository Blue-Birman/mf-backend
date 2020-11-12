package app

import (
	customerPresentation "github.com/rvalessandro/mf-backend/modules/customers/presentation"
	productPresentation "github.com/rvalessandro/mf-backend/modules/products/presentation"
	transactionPresentation "github.com/rvalessandro/mf-backend/modules/transactions/presentation"
	wishlistPresentation "github.com/rvalessandro/mf-backend/modules/wishlists/presentation"
)

func mapURLs() {
	customerPresentation.MapURLs(*router)
	productPresentation.MapURLs(*router)
	transactionPresentation.MapURLs(*router)
	wishlistPresentation.MapURLs(*router)
}

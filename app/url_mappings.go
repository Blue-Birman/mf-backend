package app

import (
	cartPresentation "github.com/rvalessandro/mf-backend/modules/carts/presentation"
	categoryPresentation "github.com/rvalessandro/mf-backend/modules/categories/presentation"
	customerPresentation "github.com/rvalessandro/mf-backend/modules/customers/presentation"
	productPresentation "github.com/rvalessandro/mf-backend/modules/products/presentation"
	transactionPresentation "github.com/rvalessandro/mf-backend/modules/transactions/presentation"
)

func mapURLs() {
	customerPresentation.MapURLs(router)
	productPresentation.MapURLs(router)
	transactionPresentation.MapURLs(router)
	categoryPresentation.MapURLs(router)
	cartPresentation.MapURLs(router)
}

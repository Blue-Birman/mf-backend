package app

import (
	customerPresentation "github.com/rvalessandro/mf-backend/modules/customers/presentation"
	productPresentation "github.com/rvalessandro/mf-backend/modules/products/presentation"
)

func mapURLs() {
	customerPresentation.MapURLs(*router)
	productPresentation.MapURLs(*router)
}

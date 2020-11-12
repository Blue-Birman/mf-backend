package app

import "github.com/rvalessandro/mf-backend/modules/customers/presentation"

func mapURLs() {
	presentation.MapURLs(*router)
}

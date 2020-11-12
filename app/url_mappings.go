package app

import "github.com/rvalessandro/mf-backend/controllers/customers"

func mapUrls() {
	router.GET("/customers/:customer_id", customers.Get)
}

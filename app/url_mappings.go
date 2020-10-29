package app

import "github.com/rvalessandro/mf-backend/controllers/customers"

func mapUrls() {
	router.GET("/customers/:customer_id", customers.Get)
	// router.GET("/users", users.Search)
	// router.POST("/users", users.Create)
	// router.PUT("/users/:user_id", users.Update)
	// router.PATCH("/users/:user_id", users.Update)
	// router.DELETE("/users/:user_id", users.Delete)
}

package app

import (
	"github.com/buddabic2/bookstore_users-api/controllers/ping"
	"github.com/buddabic2/bookstore_users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users/:user_id", users.GetUser)
	router.GET("/users/search", users.CreateUser)
	router.POST("/users", users.CreateUser)
}

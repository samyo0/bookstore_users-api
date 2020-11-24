package app

import (
	"github.com/samyo0/bookstore_users-api/controllers/ping"
	"github.com/samyo0/bookstore_users-api/controllers/users"
)

func mapURL() {
	router.GET("/ping", ping.Ping)
	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
	router.POST("/search/:user_id", users.SearchUser)
}

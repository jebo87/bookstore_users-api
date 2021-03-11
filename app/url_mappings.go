package app

import (
	"github.com/jebo87/bookstore_users-api/controllers/ping"
	"github.com/jebo87/bookstore_users-api/controllers/users"
)

//MapURLs sets up the handlers for our web server.
func MapURLs() {
	router.GET("/ping", ping.Ping)
	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)

}

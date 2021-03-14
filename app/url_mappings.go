package app

import (
	"github.com/jebo87/bookstore_users-api/controllers/ping"
	"github.com/jebo87/bookstore_users-api/controllers/users"
)

//MapURLs sets up the handlers for our web server.
func MapURLs() {
	router.GET("/ping", ping.Ping)
	router.GET("/users/:user_id", users.Get)
	router.POST("/users", users.Create)
	router.PUT("/users/:user_id", users.Update)
	router.PATCH("/users/:user_id", users.Update)
	router.DELETE("/users/:user_id", users.Delete)
	router.GET("/internal/users/search", users.Search)

}

package app

import "github.com/jebo87/bookstore_users-api/controllers"

func MapURLs() {
	router.GET("/ping", controllers.Ping)
}

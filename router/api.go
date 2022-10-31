package router

import (
	"api_standard/controllers"
)

func Routes() {
	routes := ApiRouter()

	routes.Handle("GET", "/users", controllers.Index)
	routes.Handle("POST", "/users", controllers.Store)

	routes.InitRoute()
}

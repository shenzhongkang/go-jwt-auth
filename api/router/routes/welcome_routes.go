package routes

import (
	"gome/api/controllers"
	"net/http"
)

var welcomeRoutes = []Route{
	{
		URI: "/welcome",
		Method: http.MethodGet,
		Handler: controllers.Welcome,
		AuthRequired: false,
	},
}
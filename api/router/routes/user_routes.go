package routes

import (
	"gome/api/controllers"
	"net/http"
)

var usersRoutes = []Route{
	{
		URI: "/users",
		Method: http.MethodGet,
		Handler: controllers.GetUsers,
		AuthRequired: false,
	},
	{
		URI: "/users/{id}",
		Method: http.MethodGet,
		Handler: controllers.GetUser,
		AuthRequired: false,
	},
	{
		URI: "/users",
		Method: http.MethodPost,
		Handler: controllers.CreateUser,
		AuthRequired: false,
	},
	{
		URI: "/users/{id}",
		Method: http.MethodPut,
		Handler: controllers.UpdateUser,
		AuthRequired: true,
	},
	{
		URI: "/users/{id}",
		Method: http.MethodDelete,
		Handler: controllers.DeleteUser,
		AuthRequired: true,
	},
}
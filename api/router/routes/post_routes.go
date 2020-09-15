package routes

import (
	"gome/api/controllers"
	"net/http"
)

var postRoutes = []Route{
	{
		URI: "/posts",
		Method: http.MethodGet,
		Handler: controllers.GetPosts,
		AuthRequired: false,
	},
	{
		URI: "/posts",
		Method: http.MethodGet,
		Handler: controllers.CreatePost,
		AuthRequired: true,
	},
	{
		URI: "/posts/{id}",
		Method: http.MethodGet,
		Handler: controllers.GetPost,
		AuthRequired: false,
	},
	{
		URI: "/posts/{id}",
		Method: http.MethodPut,
		Handler: controllers.UpdatePost,
		AuthRequired: true,
	},
	{
		URI: "/posts/{id}",
		Method: http.MethodDelete,
		Handler: controllers.DeletePost,
		AuthRequired: true,
	},
}
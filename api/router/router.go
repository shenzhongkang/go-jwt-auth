package router

import (
	"github.com/gorilla/mux"
	"gome/api/router/routes"
)

func New() *mux.Router  {
	r := mux.NewRouter().StrictSlash(true)
	r = r.PathPrefix("/api/v1").Subrouter()
	return routes.SetupRoutesWithMiddlewares(r)
}
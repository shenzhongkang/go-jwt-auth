package controllers

import (
	"gome/api/responses"
	"net/http"
)

// Welcome to gome
func Welcome(w http.ResponseWriter, r *http.Request)  {
	responses.JSON(w, http.StatusOK, "Hello Gome!")
}
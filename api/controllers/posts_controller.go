package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"gome/api/database"
	"gome/api/models"
	"gome/api/repository"
	"gome/api/repository/crud"
	"gome/api/responses"
	"gome/api/utils/types"
	"net/http"
	"strconv"
)

// GetPosts from the DB
func GetPosts(w http.ResponseWriter, r *http.Request)  {
	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := crud.NewRepositoryPostsCRUD(db)

	func(postRepository repository.PostRepository) {
		posts, err := postRepository.FindAll()
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		responses.JSON(w, http.StatusOK, posts)
	}(repo)
}

// CreatePost from the DB
func CreatePost(w http.ResponseWriter, r *http.Request)  {
	user := models.User{}
	user = r.Context().Value(types.UserKey("users")).(models.User)

	post := models.Post{}
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	post.AuthorID = user.ID
	post.Prepare()
	err = post.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := crud.NewRepositoryPostsCRUD(db)

	func(postRepository repository.PostRepository) {
		post, err := postRepository.Save(post)
		if err != nil {
			responses.ERROR(w, http.StatusUnprocessableEntity, err)
			return
		}
		w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, post.ID))
		responses.JSON(w, http.StatusCreated, post)
	}(repo)
}

// GetPost from the DB
func GetPost(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	uid, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := crud.NewRepositoryPostsCRUD(db)
	func(postRepository repository.PostRepository) {
		post, err := postRepository.FindByID(uint32(uid))
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}
		responses.JSON(w, http.StatusOK, post)
	}(repo)
}

// UpdatePost from the DB
func UpdatePost(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	uid, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	post := models.Post{}
	err = json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user := models.User{}
	user = r.Context().Value(types.UserKey("user")).(models.User)
	post.AuthorID = user.ID

	err = post.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := crud.NewRepositoryPostsCRUD(db)

	func(postRepository repository.PostRepository) {
		rows, err := postRepository.Update(uint32(uid), post)
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}
		responses.JSON(w, http.StatusOK, rows)
	}(repo)
}

// DeletePost from the DB
func DeletePost(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	uid, err := strconv.ParseInt(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}

	user := models.User{}
	user = r.Context().Value(types.UserKey("user")).(models.User)

	db, err := database.Connect()
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := crud.NewRepositoryPostsCRUD(db)

	func(postRepository repository.PostRepository) {
		_, err := postRepository.Delete(uint32(uid), user.ID)
		if err != nil {
			responses.ERROR(w, http.StatusBadRequest, err)
			return
		}
		w.Header().Set("Entity", fmt.Sprintf("%d", uid))
		responses.JSON(w, http.StatusNoContent, "")
	}(repo)
}
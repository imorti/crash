package controller

import (
	"encoding/json"
	"net/http"

	"github.com/imorti/crash/rest-api/entity"
	"github.com/imorti/crash/rest-api/service"
)

// Post - to be used in post requests
var (
	postService service.PostService = service.NewPostService()
)

// PostContoller - handles incoming requests for Get/Add of posts
type PostContoller interface {
	GetPosts(resp http.ResponseWriter, req *http.Request)
	AddPosts(resp http.ResponseWriter, req *http.Request)
}

// GetPosts - get all posts from the repository
func GetPosts(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	posts, err := repo.FindAll()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error getting the posts"})

	}
	postService.FindAll
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(posts)
}

// AddPosts - Adds posts to the data store
func AddPosts(resp http.ResponseWriter, req *http.Request) {
	var post entity.Post
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error decoding the post"})
		return
	}
	// validate the post
	errInvalid := postService.Validate(&post)
	if errInvalid != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: errInvalid.Error()})
		return
	}
	// create post now that we're validated
	errCreate := postService.Create(&post)
	if errCreate != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error creating the post"})
		return
	}
	result, errCreate := postService.Create(&post)

	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(result)
}

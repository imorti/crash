package controller

import (
	"encoding/json"
	"net/http"

	"github.com/imorti/crash/rest-api/entity"
	"github.com/imorti/crash/rest-api/errors"
	"github.com/imorti/crash/rest-api/service"
)

type controller struct{}

// Post - to be used in post requests
var (
	postService service.PostService
)

// NewPostController - Constructor for new PostController
func NewPostController(service service.PostService) PostController {
	postService = service
	return &controller{}
}

// PostController - handles incoming requests for Get/Add of posts
type PostController interface {
	GetPosts(resp http.ResponseWriter, req *http.Request)
	AddPosts(resp http.ResponseWriter, req *http.Request)
}

// GetPosts - get all posts from the repository
func (*controller) GetPosts(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	posts, err := postService.FindAll()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error getting the posts"})

	}

	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(posts)
}

// AddPosts - Adds posts to the data store
func (*controller) AddPosts(resp http.ResponseWriter, req *http.Request) {
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
	result, errCreate := postService.Create(&post)
	if errCreate != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error creating the post"})
		return
	}

	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(result)
}

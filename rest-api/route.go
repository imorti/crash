package main

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/imorti/crash/rest-api/entity"
	"github.com/imorti/crash/rest-api/repository"
)

// Post - to be used in post requests
var (
	repo repository.PostRepository = repository.NewPostRepository()
)

func getPosts(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	posts, err := repo.FindAll()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error getting the posts array}`))
		return
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(posts)
}

func addPosts(resp http.ResponseWriter, req *http.Request) {
	var post entity.Post
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error marshalling the request"}`))
		return
	}
	post.ID = rand.Int63()
	repo.Save(&post)
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(post)
}

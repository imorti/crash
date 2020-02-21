package main

import (
	"encoding/json"
	"net/http"
)

// Post - to be used in post requests
type Post struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

var (
	posts []Post
)

func init() {
	posts = []Post{Post{Id: 1, Title: "Title 1", Text: "Text 1"}}
}

func getPosts(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	result, error := json.Marshal(posts)
	if error != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error marshalling the posts array}`))
		return
	}
	resp.WriteHeader(http.StatusOK)
	resp.Write((result))
}

func addPosts(resp http.ResponseWriter, req *http.Request) {
	var post Post
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error marshalling the request"}`))
		return
	}
	post.Id = len(posts) + 1
	posts = append(posts, post)
	resp.Header().Set("Content-type", "application/json")
	resp.WriteHeader(http.StatusOK)
	result, err := json.Marshal(posts)
	resp.Write(result)
}

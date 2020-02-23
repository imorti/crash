package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

var (
	chiDispatcher = chi.NewRouter()
)

type chiRouter struct{}

// NewChiRouter - constructor
func NewChiRouter() Router {
	return &chiRouter{}
}

// GET - handles GET requests
func (*chiRouter) GET(uri string, f func(resp http.ResponseWriter, req *http.Request)) {
	chiDispatcher.Get(uri, f)
}

// POST - handles post requests
func (*chiRouter) POST(uri string, f func(resp http.ResponseWriter, req *http.Request)) {
	chiDispatcher.Post(uri, f)
}

// SERVE - serves up our application
func (*chiRouter) SERVE(port string) {
	fmt.Printf("Chi HTTP server running on port %v ", port)
	http.ListenAndServe(port, chiDispatcher)
}

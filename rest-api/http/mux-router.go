package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type muxRouter struct {
}

var (
	muxDispatcher = mux.NewRouter()
)

// NewMuxRouter - Constructor for mux router
func NewMuxRouter() Router {
	return &muxRouter{}
}

// GET - handles get HTTP requests
func (*muxRouter) GET(uri string, f func(resp http.ResponseWriter, req *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("GET")
}

// POST - handles post HTTP requests
func (*muxRouter) POST(uri string, f func(resp http.ResponseWriter, req *http.Request)) {

	muxDispatcher.HandleFunc(uri, f).Methods("POST")
}

// SERVE - serves up our application
func (*muxRouter) SERVE(port string) {
	fmt.Printf("Mux HTTP server running on port %v", port)
	http.ListenAndServe(port, muxDispatcher)
}

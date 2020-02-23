package main

import (
	"fmt"
	"net/http"

	"github.com/imorti/crash/rest-api/http"
)

var (
	httpRouter router.Router = router.NewMuxRouter()
)

func main() {
	const port string = ":8080"

	httpRouter.GET("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Server up and running. Well done!")
	})

	// httpRouter.GET("/posts", getPosts)
	// httpRouter.POST("/posts", addPosts)

	httpRouter.SERVE(port)

}

package main

import (
	"fmt"
	"net/http"

	"github.com/imorti/crash/rest-api/controller"
	"github.com/imorti/crash/rest-api/http"
)

var (
	httpRouter     router.Router             = router.NewMuxRouter()
	postController controller.PostController = controller.NewPostController()
)

func main() {
	const port string = ":8000"

	httpRouter.GET("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Server up and running. Well done!")
	})

	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPosts)

	httpRouter.SERVE(port)

}

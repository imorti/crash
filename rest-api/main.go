package main

import (
	"fmt"
	"net/http"

	"github.com/imorti/crash/rest-api/controller"
	"github.com/imorti/crash/rest-api/http"
	"github.com/imorti/crash/rest-api/repository"
	"github.com/imorti/crash/rest-api/service"
)

var (
	postRepository repository.PostRepository = repository.NewFirestoreRepository()
	postService    service.PostService       = service.NewPostService(postRepository)
	postController controller.PostController = controller.NewPostController(postService)
	httpRouter     router.Router             = router.NewChiRouter()
)

func main() {
	const port string = ":8001"

	httpRouter.GET("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Server up and running. Well done!")
	})

	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPosts)

	httpRouter.SERVE(port)

}

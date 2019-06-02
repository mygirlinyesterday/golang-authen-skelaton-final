package routes

import (
	"net/http"

	"github.com/playground/api/controllers"
)

var postsRoutes = []Route{
	Route{
		Uri:     "/posts",
		Method:  http.MethodPost,
		Handler: controllers.CreatePost,
	},
}

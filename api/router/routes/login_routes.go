package routes

import (
	"net/http"

	"github.com/playground/api/controllers"
)

var loginRoutes = []Route{
	Route{
		Uri:          "/login",
		Method:       http.MethodPost,
		Handler:      controllers.Login,
		AuthRequired: false,
	},
}

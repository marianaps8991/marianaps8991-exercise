package routes

import (
	"net/http"

	"api/controllers"
)

var loginRoutes = []Route{
	Route{
		Uri:           "/login",
		Method:        http.MethodPost,
		Handler:       controllers.LogIn,
		AuthRequired:  false,
		AdminRequired: false,
	},
}

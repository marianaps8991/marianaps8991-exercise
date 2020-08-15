package routes

import (
	"net/http"

	"api/controllers"
)

var usersRoutes = []Route{
	Route{
		Uri:           "/users",
		Method:        http.MethodGet,
		Handler:       controllers.GetUsers,
		AuthRequired:  true,
		AdminRequired: false,
	},
	Route{
		Uri:           "/user",
		Method:        http.MethodPost,
		Handler:       controllers.CreateUsers,
		AuthRequired:  true,
		AdminRequired: true,
	},
	Route{
		Uri:           "/user/{id}",
		Method:        http.MethodGet,
		Handler:       controllers.GetUser,
		AuthRequired:  true,
		AdminRequired: false,
	},
	Route{
		Uri:           "/user/{id}",
		Method:        http.MethodPut,
		Handler:       controllers.UpdateUsers,
		AuthRequired:  true,
		AdminRequired: true,
	},
	Route{
		Uri:           "/user/{id}",
		Method:        http.MethodDelete,
		Handler:       controllers.DeleteUsers,
		AuthRequired:  true,
		AdminRequired: true,
	},
}

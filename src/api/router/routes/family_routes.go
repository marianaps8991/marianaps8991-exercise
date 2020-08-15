package routes

import (
	"net/http"

	"api/controllers"
)

var familyRoutes = []Route{
	Route{
		Uri:           "/family",
		Method:        http.MethodPost,
		Handler:       controllers.CreateFamily,
		AuthRequired:  true,
		AdminRequired: false,
	},
	Route{
		Uri:           "/families",
		Method:        http.MethodGet,
		Handler:       controllers.GetFamilies,
		AuthRequired:  true,
		AdminRequired: false,
	},
	Route{
		Uri:           "/family/{id}",
		Method:        http.MethodGet,
		Handler:       controllers.GetFamily,
		AuthRequired:  true,
		AdminRequired: false,
	},
	Route{
		Uri:           "/family/{id}",
		Method:        http.MethodPut,
		Handler:       controllers.UpdateFamily,
		AuthRequired:  true,
		AdminRequired: false,
	},
	Route{
		Uri:           "/family/{id}",
		Method:        http.MethodDelete,
		Handler:       controllers.DeleteFamily,
		AuthRequired:  true,
		AdminRequired: false,
	},
}

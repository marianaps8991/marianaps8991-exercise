package routes

import (
	"net/http"

	"api/middlewares"

	"github.com/gorilla/mux"
)

type Route struct {
	Uri           string
	Method        string
	Handler       func(http.ResponseWriter, *http.Request)
	AuthRequired  bool
	AdminRequired bool
}

func Load() []Route {
	routes := usersRoutes
	routes = append(routes, familyRoutes...)
	routes = append(routes, loginRoutes...)
	return routes
}

func SetUpRoutesWithMiddleWare(r *mux.Router) *mux.Router {

	for _, routes := range Load() {

		if routes.AuthRequired {
			if routes.AdminRequired {
				r.HandleFunc(routes.Uri,
					middlewares.SetMiddlewareCors(
						middlewares.SetMiddlewareLogger(
							middlewares.SetMiddlewareJSON(
								middlewares.SetMiddlewareAdminAuth(
									routes.Handler)))),
				).Methods(routes.Method, http.MethodOptions)

			} else {
				r.HandleFunc(routes.Uri,
					middlewares.SetMiddlewareCors(
						middlewares.SetMiddlewareLogger(
							middlewares.SetMiddlewareJSON(
								middlewares.SetMiddlewareAuthentication(routes.Handler)))),
				).Methods(routes.Method, http.MethodOptions)
			}
		} else {
			r.HandleFunc(routes.Uri,
				middlewares.SetMiddlewareCors(
					middlewares.SetMiddlewareLogger(
						middlewares.SetMiddlewareJSON(routes.Handler))),
			).Methods(routes.Method, http.MethodOptions)
		}
	}
	return r
}

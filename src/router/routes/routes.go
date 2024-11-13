package routes

import (
	"net/http"
	"web-app-go/src/middleware"

	"github.com/gorilla/mux"
)

// Route representa todas as rotas da aplicação
type Route struct {
	URI            string
	Method         string
	Function       func(http.ResponseWriter, *http.Request)
	Authentication bool
}

// Config coloca todas as rotas dentro do router
func Config(router *mux.Router) *mux.Router {
	routes := routesLogin
	routes = append(routes, userRoutes...)
	routes = append(routes, homePageRoute)
	routes = append(routes, publicationRoutes...)
	for _, route := range routes {
		if route.Authentication {
			router.HandleFunc(
				route.URI,
				middleware.Logger(middleware.Authenticate(route.Function)),
			).Methods(route.Method)
		} else {
			router.HandleFunc(
				route.URI,
				middleware.Logger(route.Function),
			).Methods(route.Method)
		}
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}

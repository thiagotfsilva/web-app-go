package routes

import (
	"net/http"

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

	for _, route := range routes {
		router.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}

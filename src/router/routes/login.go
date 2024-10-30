package routes

import (
	"net/http"
	"web-app-go/src/controllers"
)

var routesLogin = []Route{
	{
		URI:            "/",
		Method:         http.MethodGet,
		Function:       controllers.LoadLoginView,
		Authentication: false,
	},
	{
		URI:            "/login",
		Method:         http.MethodGet,
		Function:       controllers.LoadLoginView,
		Authentication: false,
	},
	{
		URI:            "/login",
		Method:         http.MethodPost,
		Function:       controllers.Login,
		Authentication: false,
	},
}

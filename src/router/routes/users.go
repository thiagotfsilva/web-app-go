package routes

import (
	"net/http"
	"web-app-go/src/controllers"
)

var userRoutes = []Route{
	{
		URI:            "/create",
		Method:         http.MethodGet,
		Function:       controllers.LoadUserRegisterView,
		Authentication: false,
	},
}

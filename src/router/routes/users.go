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
	{
		URI:            "/users",
		Method:         http.MethodPost,
		Function:       controllers.Register,
		Authentication: false,
	},
	{
		URI:            "/find-users",
		Method:         http.MethodGet,
		Function:       controllers.LoadUserPage,
		Authentication: true,
	},
}

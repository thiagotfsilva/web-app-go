package routes

import (
	"net/http"
	"web-app-go/src/controllers"
)

var homePageRoute = Route{

	URI:            "/home",
	Method:         http.MethodGet,
	Function:       controllers.LoadHomePage,
	Authentication: true,
}

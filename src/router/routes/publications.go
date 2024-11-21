package routes

import (
	"net/http"
	"web-app-go/src/controllers"
)

var publicationRoutes = []Route{
	{
		URI:            "/publications",
		Method:         http.MethodPost,
		Function:       controllers.CreatePublication,
		Authentication: true,
	},
	{
		URI:            "/publications/{publicationId}/likes",
		Method:         http.MethodPost,
		Function:       controllers.LikePublication,
		Authentication: true,
	},
}

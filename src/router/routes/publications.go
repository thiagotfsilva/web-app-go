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
		URI:            "/publications/{publicationId}/like",
		Method:         http.MethodPost,
		Function:       controllers.LikePublication,
		Authentication: true,
	},
	{
		URI:            "/publications/{publicationId}/dislike",
		Method:         http.MethodPost,
		Function:       controllers.DislikePublication,
		Authentication: true,
	},
}

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
	{
		URI:            "/publications/{publicationId}/edit",
		Method:         http.MethodGet,
		Function:       controllers.LoadEditPage,
		Authentication: true,
	},
	{
		URI:            "/publications/{publicationId}",
		Method:         http.MethodPut,
		Function:       controllers.EditPublication,
		Authentication: true,
	},
	{
		URI:            "/publications/{publicationId}",
		Method:         http.MethodDelete,
		Function:       controllers.DeletePublication,
		Authentication: true,
	},
}

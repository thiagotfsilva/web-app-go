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
		Function:       controllers.LoadUserFindView,
		Authentication: true,
	},
	{
		URI:            "/users/{userId}",
		Method:         http.MethodGet,
		Function:       controllers.LoadUserProfileView,
		Authentication: true,
	},
	{
		URI:            "/users/{userId}/unfollow",
		Method:         http.MethodPost,
		Function:       controllers.UnfollowUser,
		Authentication: true,
	},
	{
		URI:            "/users/{userId}/follow",
		Method:         http.MethodPost,
		Function:       controllers.FollowUser,
		Authentication: true,
	},
}

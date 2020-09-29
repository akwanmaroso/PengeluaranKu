package routes

import (
	"net/http"

	"github.com/akwanmaroso/PengeluaranKu/api/controllers"
)

var allRoutes = []Route{
	{
		Uri:          "/index",
		Method:       http.MethodGet,
		Handler:      controllers.Index,
		AuthRequired: false,
	},
	{
		Uri:          "/users",
		Method:       http.MethodPost,
		Handler:      controllers.CreateUser,
		AuthRequired: false,
	},
}

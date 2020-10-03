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
	{
		Uri:          "/login",
		Method:       http.MethodPost,
		Handler:      controllers.Login,
		AuthRequired: false,
	},
	{
		Uri:          "/transactions",
		Method:       http.MethodGet,
		Handler:      controllers.GetTransactions,
		AuthRequired: false,
	},
	{
		Uri:          "/transactions",
		Method:       http.MethodPost,
		Handler:      controllers.CreateTransaction,
		AuthRequired: false,
	},
	{
		Uri:          "/categories",
		Method:       http.MethodPost,
		Handler:      controllers.CreateCategory,
		AuthRequired: false,
	},
	{
		Uri:          "/categories",
		Method:       http.MethodGet,
		Handler:      controllers.GetCategories,
		AuthRequired: false,
	},
}

package routes

import (
	"net/http"

	"github.com/akwanmaroso/PengeluaranKu/api/controllers"
)

var allRoutes = []Route{
	{
		URI:          "/index",
		Method:       http.MethodGet,
		Handler:      controllers.Index,
		AuthRequired: false,
	},
	{
		URI:          "/users",
		Method:       http.MethodPost,
		Handler:      controllers.CreateUser,
		AuthRequired: false,
	},
	{
		URI:          "/login",
		Method:       http.MethodPost,
		Handler:      controllers.Login,
		AuthRequired: false,
	},
	{
		URI:          "/transactions",
		Method:       http.MethodGet,
		Handler:      controllers.GetTransactions,
		AuthRequired: true,
	},
	{
		URI:          "/transactions",
		Method:       http.MethodPost,
		Handler:      controllers.CreateTransaction,
		AuthRequired: true,
	},
	{
		URI:          "/transactions/{id}",
		Method:       http.MethodDelete,
		Handler:      controllers.DeleteTransaction,
		AuthRequired: false,
	},
	{
		URI:          "/categories",
		Method:       http.MethodPost,
		Handler:      controllers.CreateCategory,
		AuthRequired: false,
	},
	{
		URI:          "/categories",
		Method:       http.MethodGet,
		Handler:      controllers.GetCategories,
		AuthRequired: false,
	},
}

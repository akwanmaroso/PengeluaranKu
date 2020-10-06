package routes

import (
	"fmt"
	"net/http"

	"github.com/akwanmaroso/PengeluaranKu/api/middlewares"
	"github.com/gorilla/mux"
)

type Route struct {
	Uri          string
	Method       string
	Handler      func(http.ResponseWriter, *http.Request)
	AuthRequired bool
}

func Load() []Route {
	routes := allRoutes
	return routes
}

func SetupRoutes(r *mux.Router) *mux.Router {
	for _, route := range Load() {
		r.HandleFunc(route.Uri,
			middlewares.SetMiddlewareLogger(
				middlewares.SetMiddlewareJSON(route.Handler)),
		).Methods(route.Method)
		fmt.Println(route.Uri, route.Method)
	}
	return r
}

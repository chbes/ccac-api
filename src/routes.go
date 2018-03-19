package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Route - API route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes - Array of route
type Routes []Route

const apiName = "/ccac-api"

var routes = Routes{

	Route{
		"Welcome",
		"GET",
		apiName + "/",
		Welcome,
	},

	Route{
		"GetTransactionsHandler",
		"GET",
		apiName + "/transactions",
		GetTransactionsHandler,
	},

	Route{
		"CreateTransactionHandler",
		"POST",
		apiName + "/transactions",
		CreateTransactionHandler,
	},

	Route{
		"GetUsersHandler",
		"GET",
		apiName + "/users",
		GetUsersHandler,
	},

	Route{
		"WebSockectConnection",
		"GET",
		apiName + "/ws",
		WebSockectConnection,
	},
}

// NewRouter - return a new router
func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}

	return router
}

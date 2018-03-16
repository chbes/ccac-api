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

var routes = Routes{

	Route{
		"Welcome",
		"GET",
		"/",
		Welcome,
	},

	Route{
		"GetTransactionsHandler",
		"GET",
		"/transactions",
		GetTransactionsHandler,
	},

	Route{
		"CreateTransactionHandler",
		"POST",
		"/transactions",
		CreateTransactionHandler,
	},

	Route{
		"GetUsersHandler",
		"GET",
		"/users",
		GetUsersHandler,
	},

	Route{
		"WebSockectConnection",
		"GET",
		"/ws",
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

package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{

	Route{
		"Is Node Online?",
		"Get",
		"/api/online",
		NodeOnline,
	},
	Route{
		"Debug get Number",
		"Get",
		"/api/getNumber",
		NewNumber,
	},
}

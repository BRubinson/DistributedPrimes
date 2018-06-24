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
		"Is Overlord Alive",
		"Get",
		"/api/available",
		Available,
	},
	Route{
		"new Node Connection",
		"Post",
		"/api/connect",
		ConnectNewNode,
	},
	Route{
		"master Portal",
		"Post",
		"/api/master",
		MasterPortal,
	},
	Route{
		"NUMBY",
		"Get",
		"/api/getnum",
		GetCount,
	},
}

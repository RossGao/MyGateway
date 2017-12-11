/*
Setup routers whose requests are listened by the service.*/

package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Router struct { // Router is defined as hadler of httrp request. Because its has implemented ServeHttp func.
	Name, Method, Pattern string
	HandlerFunc           http.HandlerFunc
}

type routers []Router

var iniRouters = routers{
	Router{
		Name:        "User authentication",
		Method:      "POST",
		Pattern:     "/authenticate",
		HandlerFunc: Authenticate,
	},
	Router{
		Name:        "User authorization",
		Method:      "GET",
		Pattern:     "/authorization",
		HandlerFunc: Authorize,
	},
}

// NewRouter return router which contains pathes that will be listened by the service.
func NewRouter() *mux.Router {
	baseRouter := mux.NewRouter().StrictSlash(true)
	loginRouter := baseRouter.PathPrefix("/{productID:[a-zA-Z]{3}}/").Subrouter() // 所有产品线的简称都是三个字母
	for _, router := range iniRouters {
		loginRouter.Methods(router.Method).
			Path(router.Pattern).
			Name(router.Name).
			Handler(router.HandlerFunc)
	}

	return loginRouter
}

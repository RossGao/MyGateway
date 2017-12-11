/*
Handlers that are defined to handle login operations such as authentication and authorization.*/

package main

import (
	registry "Registry"
	token "Token"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type UserInfo struct {
	Name string
	Id   int64
}

type AuthenticateInfo struct {
	UserInfo
	HasLoggedIn       bool
	ExpireIn          int
	AuthorizationCode string
}

func Authenticate(w http.ResponseWriter, r *http.Request) {
	// TODO: redirect call to authenticate user in different product's servers.
	parameters := mux.Vars(r)
	productID := parameters["productID"]
	productHost := registry.HostAddress(productID)
	productRoute := mux.NewRouter()
	productRoute.Host("{hostAddress}").Path("/Authenticate").Methods("POST").Name("Authenticate") // The actual url may different from this...(maybe /auth/)
	targetURL, err := productRoute.Get("Authenticate").URL(productHost)
	if err == nil {
		response, err := http.Post(targetURL.RequestURI(), "application/json", r.Body)
		if err == nil {
			var authInfo AuthenticateInfo
			err := json.NewDecoder(response.Body).Decode(&authInfo)
			if err == nil {
				// TODO: check from cache if the user has already logged in and set the authorization code and its expiration time ticks.

			}
		}
	}
}

type FunctionalPermission struct {
	Id, ParentId, Order         int64
	Name, Type, Icon, Url, Rank string
}

type RangePermission struct {
	CustomerId, CustomerName, ParentId string
}

type AuthorizationInfo struct {
	UserInfo
	token.TokenInfo
	FunctionalPermission
	RangePermission
}

func Authorize(w http.ResponseWriter, r *http.Request) {
	//TODO: redirect to the specified host to do authorization, then call authorization server to get token info.
	result := AuthorizationInfo{}
	fmt.Println(result)
}

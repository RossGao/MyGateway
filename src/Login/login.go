package main

/*
	The entrance of client login request which contains authentication and authorization requests through gateway. This file is responsible for
	1. Setup requests listeners and handlers for login requests.
	2. Redirect authentication request.
	3. Before redirect authorization request, get client identity then call authorization service to validate client credential and issue token.
	4. Cache users login info for login status control.
*/
import (
	"log"
)

func main() {
	log.Fatal(":8080", NewRouter())
}

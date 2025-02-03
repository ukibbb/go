package main

import "net/http"

func onlyAdmin(fn ApiFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		// get user from db
		// check it's role
		// return unauthorized if not correct role

	}

}

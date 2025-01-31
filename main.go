package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Repository interface {
	Get()
	Create()
	Update()
	Delete()
}

type UserRegisterRequest struct {
	username        string
	email           string
	password        string
	passwordConfirm string
}

func (r *UserRegisterRequest) validate() []error {
	var errors []error
	if userRegister.password != userRegister.passwordConfirm {
		errors = append(errors, ApiError{
			status: http.StatusBadRequest,
			msg:    "Password doesn't match confirmed password",
		})
	}
	return nil
	// check for username uniqnes
	// check for email uniqnes

}

// ctrl + d | ctrl + u scrolling
type ApiError struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

func (e ApiError) Error() string {
	return fmt.Sprintf("error: %s status code: %d\n", e.msg, e.status)
}

func main() {
	const listenAddr string = ":3000"

	router := http.NewServeMux()

	rh := http.RedirectHandler("https://github.com", 307)

	router.Handle("/foo", rh)
	// mux.Handle("/register", http.HandlerFunc(handleRegister))
	router.HandleFunc("GET /register", handler(handleRegister))

	log.Printf("Listening on %s\n", listenAddr)

	http.ListenAndServe(listenAddr, router)
}

func handleRegister(w http.ResponseWriter, r *http.Request) error {
	request := new(UserRegisterRequest)
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return ApiError{
			Status: http.StatusBadRequest,
			Msg:    "Wrong register payload",
		}
	}
	if errors := request.validate(); errors != nil {
		return ApiError{
			Status: http.StatusUnprocessableEntity,
			Msg:    "",
		}
	}
	return nil
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

func handler(fn apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			if e, ok := err.(ApiError); ok {
				// write json with error
				fmt.Println(e)
			}
		}
	}
}

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/mail"
)


const minUserNameLength int = 3

type Repository interface {
	Get()
	Create()
	Update()
	Delete()
}

type UserRegisterRequest struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"passwordConfirm"`
}

func (r *UserRegisterRequest) validate() map[string]string {
	var errors = map[string]string{}
	if r.Password != r.PasswordConfirm {
		errors["password"] = "password and password confirm doesn't match"
	}
	if len(r.Username) < minUserNameLength {
		errors["username"] = "username minimum length equals 3"
	}
	_, err :=  mail.ParseAddress(r.Email)
	if err != nil {
		errors["email"] = "invalid email"
	}
	// if username is not unique
	errors["username"] = fmt.Sprintf("%s and username is not unique", errors["username"])
	errors["email"] = fmt.Sprintf("%s and email is not unique", errors["email"])

	return errors

}

// ctrl + d | ctrl + u scrolling
type ApiError struct {
	Status int    `json:"status"`
	Msg    interface{} `json:"msg"`
}

func (e ApiError) Error() string {
	return fmt.Sprintf("error: %s status code: %d\n", e.Msg, e.Status)
}

func main() {
	const listenAddr string = ":3000"

	router := http.NewServeMux()

	rh := http.RedirectHandler("https://github.com", 307)

	router.Handle("/foo", rh)
	// mux.Handle("/register", http.HandlerFunc(handleRegister))
	router.HandleFunc("POST /register", handler(handleRegister))

	log.Printf("Listening on %s\n", listenAddr)

	http.ListenAndServe(listenAddr, router)
}

func handleRegister(w http.ResponseWriter, r *http.Request) error {
	request := new(UserRegisterRequest)
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		return ApiError{
			Status: http.StatusBadRequest,
			Msg:    "Wrong register payload",
		}
	}
	log.Printf("register user request: %+v error", request)
	if errors := request.validate(); len(errors) > 0 {
		fmt.Printf("errors occured: %+v\n", errors)
		return ApiError{
			Status: http.StatusUnprocessableEntity,
			Msg:    errors,
		}
	}
	return nil
}

type apiFunc func(w http.ResponseWriter, r *http.Request) error

func handler(fn apiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			if e, ok := err.(ApiError); ok {
				w.WriteHeader(e.Status)
				json.NewEncoder(w).Encode(e)

			}
		}
	}
}

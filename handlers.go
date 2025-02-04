package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

type ApiError struct {
	Status int         `json:"status"`
	Msg    interface{} `json:"msg"`
}

type ApiResponse struct {
	Status int         `json:"status"`
	Msg    interface{} `json:"msg"`
}

func (e ApiError) Error() string {
	return fmt.Sprintf("error: %s status code: %d\n", e.Msg, e.Status)
}

type ApiFunc func(w http.ResponseWriter, r *http.Request) error

func handler(fn ApiFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			if e, ok := err.(ApiError); ok {
				w.WriteHeader(e.Status)
				json.NewEncoder(w).Encode(e)
			}

		}
	}
}

type UserHandlers struct {
	r *Repository[User]
}

func NewUserHandlers(db Storage[User]) *UserHandlers {
	return &UserHandlers{
		r: NewRepository[User](db),
	}
}

func (h *UserHandlers) handleLogin(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (h *UserHandlers) handleRegister(w http.ResponseWriter, r *http.Request) error {
	ur := new(UserRegisterRequest)
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(ur); err != nil {
		return ApiError{
			Status: http.StatusBadRequest,
			Msg:    "Wrong register payload",
		}
	}

	if errors := ur.validate(); len(errors) > 0 {
		log.Printf("errors occured: %+v\n", errors)
		return ApiError{
			Status: http.StatusUnprocessableEntity,
			Msg:    errors,
		}
	}

	user := User{
		Id:        uuid.New(),
		Username:  ur.Username,
		Email:     ur.Email,
		Password:  ur.Password,
		CreatedAt: time.Now(),
		IsActive:  false,
	}

	e, err := h.r.Create(&user)
	if err != nil {
		return ApiError{
			Status: http.StatusInternalServerError,
			Msg:    err.Error(),
		}
	}
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(ApiResponse{
		Status: http.StatusCreated,
		Msg:    fmt.Sprintf("User %s has been successfully created activation email has been sent to %s", e.Username, e.Email),
	})
	return nil
}

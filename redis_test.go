package main

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestRedisHelpersCreateValues(t *testing.T) {
	user := User{
		Id:        uuid.New(),
		Username:  "TestUser",
		Email:     "test@email.com",
		Password:  "UserPassword",
		CreatedAt: time.Now(),
		IsActive:  true,
		Role:      "role",
	}
	h := NewRedisHelpers[User]()
	v, _ := h.CreateValues(user)

	h.RetriveStruct(v)

}

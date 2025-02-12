package main

import (
	"reflect"
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
		CreatedAt: time.Now().Local().Format(time.RFC822),
		IsActive:  true,
		Role:      "role",
	}
	h := NewRedisHelpers[User]()
	v, _ := h.CreateValues(user)

	s, _ := h.RetriveStruct(v)

	if s != user {
		t.Error("stuct retrieved from created values is not the same")
	}

	if !reflect.DeepEqual(s, user) {
		t.Error("stuct retrieved from created values is not the same")
	}
}

package main

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
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

type RedisUser struct {
	Username string `redis:"username"` // only for how keys are stored
	Email    string `redis:"email"`
	Role     int    `redis:"role"`
	IsActive bool   `redis:"isActive"`
}

func TestRedis(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr: ":5000",
	})

	user := RedisUser{
		Username: "user",
		Email:    "uki@bb.pl",
		Role:     0,
		IsActive: true,
	}
	id := uuid.NewString()
	if err := rdb.HSet(context.Background(), id, user).Err(); err != nil {
		t.Error("failed to set user")
	}

	var nUser RedisUser

	// u, err := rdb.HGetAll(context.Background(), id).Result() // field-value
	if err := rdb.HGetAll(context.Background(), id).Scan(&nUser); err != nil {
		t.Error("failed to get user")
	}

	if user != nUser {
		t.Error("failed to retrive proper redis user")
	}

	if !reflect.DeepEqual(user, nUser) {
		t.Error("Shit happens")
	}
}

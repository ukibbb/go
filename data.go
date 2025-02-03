package main

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id        uuid.UUID
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	IsActive  bool
	Role      string
}

type Entity = User

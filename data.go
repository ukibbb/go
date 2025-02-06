package main

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Storable interface {
	GetKeyForRedis() string
}

type User struct {
	Id        uuid.UUID
	Username  string
	Email     string
	Password  string
	CreatedAt time.Time
	IsActive  bool
	Role      string
}

func (u User) GetKeyForRedis() string {
	return fmt.Sprintf("user:%s", u.Id)
}

type Order struct {
	Id uuid.UUID
}

func (o Order) GetKeyForRedis() string {
	return fmt.Sprintf("order:%s", o.Id)
}

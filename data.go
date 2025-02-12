package main

import (
	"fmt"

	"github.com/google/uuid"
)

type Storable interface {
	GetKeyForRedis() string
}

type User struct {
	Id        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt string    `json:"createdAt"`
	IsActive  bool      `json:"isActive"`
	Role      string    `json:"role"`
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

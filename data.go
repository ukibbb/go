package main

import (
	"fmt"

	"github.com/google/uuid"
)

type Storable interface {
	GetKeyForRedis() string
	Namespace() string
}

type User struct {
	Id        string `json:"id" redis:"id"`
	Username  string `json:"username" redis:"username"`
	Email     string `json:"email" redis:"email"`
	Password  string `json:"password"  redis:"password"`
	CreatedAt string `json:"createdAt" redis:"createdAt"`
	IsActive  bool   `json:"isActive" redis:"isActive"`
	Role      string `json:"role" redis:"role"`
}

func (u User) GetKeyForRedis() string {
	return fmt.Sprintf("%s:%s", u.Namespace(), u.Id)
}

func (u User) Namespace() string {
	return "user"
}

type Order struct {
	Id uuid.UUID
}

func (o Order) GetKeyForRedis() string {
	return fmt.Sprintf("%s:%s", o.Namespace(), o.Id)
}
func (o Order) Namespace() string {
	return "order"
}

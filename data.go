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
	Id        uuid.UUID `json:"id" redis:"Id"`
	Username  string    `json:"username" redis:"Username"`
	Email     string    `json:"email" redis:"Email"`
	Password  string    `json:"password" redis:"Password"`
	CreatedAt time.Time `json:"createdAt" redis:"CreatedAt"`
	IsActive  bool      `json:"isActive" redis:"IsActive"`
	Role      string    `json:"role" redis:"Role"`
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

package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

func main() {
	// s := NewServer(":3000")

	// go s.Start()
	//
	ds := NewRedisDataStore[User](&redis.Options{
		Addr:     "localhost:5000",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	u, err := ds.Create(User{
		Id:       uuid.New(),
		Username: "ukasz",
		Email:    "ukasz.bulinski.com",
	})
	fmt.Println(u, "user", err, "error")

	select {}

}

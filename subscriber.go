package main

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func main() {
	c := NewRedisClient(&redis.Options{
		Addr:     "localhost:5000",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ctx := context.Background()
	c.Subscribe(ctx, "worker")
	if err != nil {
		panic(err)
	}
}

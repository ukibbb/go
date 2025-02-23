package main

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient(opts *redis.Options) *redis.Client {
	return redis.NewClient(opts)
}

func main() {
	c :=
		NewRedisClient(&redis.Options{
			Addr:     "localhost:5000",
			Password: "", // no password set
			DB:       0,  // use default DB
		})
	ctx := context.Background()
	err := c.Publish(ctx, "worker", "hello from publisher").Err()
	if err != nil {
		panic(err)
	}
}

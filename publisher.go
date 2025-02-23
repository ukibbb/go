package main

import (
	"context"

	"github.com/redis/go-redis/v9"
)

func main() {
	c :=
		redis.NewClient(&redis.Options{
			Addr:     "localhost:5000",
			Password: "", // no password set
			DB:       0,  // use default DB
		})
	ctx := context.Background()
	err := rdb.Publish(ctx, "worker").Err()
	if err != nil {
		panic(err)
	}
}

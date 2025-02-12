package main

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type RedisDataStore[T Storable] struct {
	h *RedisHelpers[T]
	c *redis.Client
}

func NewRedisDataStore[T Storable](opts *redis.Options) *RedisDataStore[T] {
	return &RedisDataStore[T]{
		h: NewRedisHelpers[T](),
		c: redis.NewClient(opts),
	}
}

func (r *RedisDataStore[T]) GetAll() (T, error) {
	var cursor uint64

	batch, cursor, err := r.c.Scan(
		context.Background(),
		cursor, "user:*",
		10).Result()

	re, err := r.c.HGetAll(context.Background(), batch[0]).Result()

	e, err := r.h.RetriveStruct(re)

	fmt.Println(e, "eeee")

	if err != nil {
		fmt.Println(err)
		return e, err
	}

	return e, nil
}

func (r *RedisDataStore[T]) Get(key uuid.UUID) (e T, err error) {
	return e, err
}

func (r *RedisDataStore[T]) Create(e T) (T, error) {
	var s T
	id := e.GetKeyForRedis()
	values, err := r.h.CreateValues(e)

	if err != nil {
		return s, err
	}

	// r.c.HMSet(context.Background(), id, e).Result()
	if err := r.c.HSet(context.Background(), id, values).Err(); err != nil {
		return s, err
	}

	return e, nil
}

func (r *RedisDataStore[T]) Update(e T) (T, error) {
	return e, nil
}

func (r *RedisDataStore[T]) Delete(id uuid.UUID) (e T, err error) {
	return e, err
}

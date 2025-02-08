package main

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

// Contains actual logic needed to retrive data from storage.
type DataStore[T Storable] interface {
	Get(id string) (T, error)
	GetAll() error
	Create(T) (T, error)
	Update(T) (T, error)
	Delete(id uuid.UUID) (T, error)
}

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

func (r *RedisDataStore[T]) GetAll() error {
	var cursor uint64
	batch, cursor, err := r.c.Scan(
		context.Background(),
		cursor, "user:*",
		10).Result()

	if err != nil {
		return err
	}

	var user User

	err = r.c.HGetAll(context.Background(), batch[0]).Scan(&user)

	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("re", user)

	return nil
}

func (r *RedisDataStore[T]) Get(key string) (e T, err error) {
	return e, err
}

func (r *RedisDataStore[T]) Create(e T) (T, error) {
	var s T
	id := e.GetKeyForRedis()
	values, err := r.h.CreateValues(e)

	if err != nil {
		return s, err
	}

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

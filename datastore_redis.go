package main

import (
	"context"
	"fmt"

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

func (r *RedisDataStore[T]) GetAll() ([]T, error) {
	var (
		cursor   uint64
		e        T
		entities []T
	)

	ctx := context.Background()

	namespace := fmt.Sprintf("%s:*", e.Namespace())

	for {
		batch, cursor, err := r.c.Scan(
			ctx,
			cursor,
			namespace,
			10).Result()

		if err != nil {
			break
		}

		for _, key := range batch {
			if err := r.c.HGetAll(ctx, key).Scan(&e); err != nil {
				fmt.Println(err)
				continue
			}
			entities = append(entities, e)
		}

		if cursor == 0 {
			break
		}
	}

	return entities, nil
}

func (r *RedisDataStore[T]) Get(key string) (e T, err error) {
	return e, err
}

func (r *RedisDataStore[T]) Create(e T) (T, error) {
	var s T
	id := e.GetKeyForRedis()

	if err := r.c.HSet(context.Background(), id, e).Err(); err != nil {
		return s, err
	}

	return e, nil
}

func (r *RedisDataStore[T]) Update(e T) (T, error) {
	return e, nil
}

func (r *RedisDataStore[T]) Delete(id string) (e T, err error) {
	return e, err
}

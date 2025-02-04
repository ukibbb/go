package main

import (
	"context"
	"fmt"
	"reflect"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

// Contains actual logic needed to retrive data from storage.
type DataStore[T Entity] interface {
	Get(id uuid.UUID) (*T, error)
	Create(*T) (*T, error)
	Update(*T) (*T, error)
	Delete(id uuid.UUID) (*T, error)
}

type InMemoryDataStore[T Entity] struct {
	storage map[uuid.UUID]*T
}

func NewInMemoryDataStore[T Entity]() *InMemoryDataStore[T] {
	return &InMemoryDataStore[T]{
		storage: make(map[uuid.UUID]*T),
	}
}

func (db *InMemoryDataStore[Entity]) Get(id uuid.UUID) (e *Entity, err error) {
	e, ok := db.storage[id]
	if !ok {
		return e, fmt.Errorf("error: entity %s not found in Storage", id)
	}
	return e, err

}
func (db *InMemoryDataStore[Entity]) Create(e *Entity) (*Entity, error) {
	uuid := uuid.New()
	db.storage[uuid] = e
	return e, nil
}

func (db *InMemoryDataStore[Entity]) Update(e *Entity) (*Entity, error) {
	return e, nil
}
func (db *InMemoryDataStore[Entity]) Delete(id uuid.UUID) (e *Entity, err error) {
	return e, err
}

type RedisDataStore[T Entity] struct {
	h *RedisHelpers
	c *redis.Client
}

func NewRedisStorage[T Entity](opts *redis.Options) *RedisDataStore[T] {
	return &RedisDataStore[T]{
		h: NewRedisHelpers(),
		c: redis.NewClient(opts),
	}
}

func (r *RedisDataStore[Entity]) GetAll() {
	t := reflect.TypeOf(Entity{}).Name()
}

func (r *RedisDataStore[Entity]) Get(id uuid.UUID) (*Entity, error) {
	return nil, nil
}

func (r *RedisDataStore[Entity]) Create(e *Entity) (*Entity, error) {

	if err := r.c.HSet(context.Background(), "", map[string]interface{}{}).Err(); err != nil {
		return nil, err
	}

	return nil, nil
}

// func (r *RedisStorage[Entity]) Update(e *Entity) (*Entity, error)    {}
// func (r *RedisStorage[Entity]) Delete(id uuid.UUID) (*Entity, error) {}

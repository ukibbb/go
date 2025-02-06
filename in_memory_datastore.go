package main

import (
	"fmt"

	"github.com/google/uuid"
)

type InMemoryDataStore[T Storable] struct {
	storage map[uuid.UUID]*T
}

func NewInMemoryDataStore[T Storable]() *InMemoryDataStore[T] {
	return &InMemoryDataStore[T]{
		storage: make(map[uuid.UUID]*T),
	}
}

func (db *InMemoryDataStore[T]) Get(id uuid.UUID) (e *T, err error) {
	e, ok := db.storage[id]
	if !ok {
		return e, fmt.Errorf("error: entity %s not found in Storage", id)
	}
	return e, err

}
func (db *InMemoryDataStore[T]) Create(e *T) (*T, error) {
	uuid := uuid.New()
	db.storage[uuid] = e
	return e, nil
}

func (db *InMemoryDataStore[T]) Update(e *T) (*T, error) {
	return e, nil
}
func (db *InMemoryDataStore[T]) Delete(id uuid.UUID) (e *T, err error) {
	return e, err
}

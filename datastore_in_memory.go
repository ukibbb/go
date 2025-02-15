package main

import (
	"fmt"

	"github.com/google/uuid"
)

type InMemoryDataStore[T Storable] struct {
	storage map[string]T
}

func NewInMemoryDataStore[T Storable]() *InMemoryDataStore[T] {
	return &InMemoryDataStore[T]{
		storage: make(map[string]T),
	}
}

func (db *InMemoryDataStore[T]) Get(id string) (e T, err error) {
	e, ok := db.storage[id]
	if !ok {
		return e, fmt.Errorf("error: entity %s not found in Storage", id)
	}
	return e, err

}

func (db *InMemoryDataStore[T]) GetAll() (t T, err error) {
	return t, err
}
func (db *InMemoryDataStore[T]) Create(e T) (T, error) {
	uuid := uuid.NewString()
	db.storage[uuid] = e
	return e, nil
}

func (db *InMemoryDataStore[T]) Update(e T) (T, error) {
	return e, nil
}
func (db *InMemoryDataStore[T]) Delete(id string) (e T, err error) {
	return e, err
}

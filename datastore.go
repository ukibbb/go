package main

import (
	"github.com/google/uuid"
)

// Contains actual logic needed to retrive data from storage.
type DataStore[T Storable] interface {
	Get(id uuid.UUID) (T, error)
	GetAll() (T, error)
	Create(T) (T, error)
	Update(T) (T, error)
	Delete(id uuid.UUID) (T, error)
}

package main

// Contains actual logic needed to retrive data from storage.
type DataStore[T Storable] interface {
	Get(id string) (T, error)
	GetAll() ([]T, error)
	Create(T) (T, error)
	Update(T) (T, error)
	Delete(id string) (T, error)
}

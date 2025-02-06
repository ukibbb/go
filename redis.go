package main

import (
	"errors"
	"reflect"
)

func NewRedisHelpers[T Storable]() *RedisHelpers[T] {
	return &RedisHelpers[T]{}
}

type RedisHelpers[T Storable] struct{}

func (h *RedisHelpers[T]) CreateValues(e T) (map[string]interface{}, error) {
	t := reflect.TypeOf(e)
	v := reflect.ValueOf(e)
	if t.Kind() != reflect.Struct {
		return nil, errors.New("values can be created from struct only")
	}
	values := make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		// fooT.Field(i) // struct field metadata
		values[t.Field(i).Name] = v.Field(i).String()
	}

	return values, nil

}

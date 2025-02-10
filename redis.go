package main

import (
	"errors"
	"fmt"
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

func (*RedisHelpers[T]) RetrieveValues(values map[string]string) (T, error) {
	var e T

	t := reflect.TypeOf(e)
	if t.Kind() != reflect.Struct {
		return e, errors.New("type in retrive values is not struct")
	}

	fv := reflect.ValueOf(&e).Elem()

	for k, v := range values {
		f := fv.FieldByName(k)
		fmt.Println("fv:", fv, "f:", f, "k:", k, "v:", v)

	}

	return e, nil

}

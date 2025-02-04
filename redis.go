package main

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/google/uuid"
)

func NewRedisHelpers() *RedisHelpers {
	return &RedisHelpers{}
}

type RedisHelpers struct{}

func (h *RedisHelpers) CreateKey(e *Entity) (string, error) {
	t := reflect.TypeOf(e)
	name := t.Name()
	uuid, err := uuid.NewUUID()
	if err != nil {
		return "", errors.New("failed to create new uuid")
	}
	return fmt.Sprintf("%s:%s", uuid, name), nil

}
func (h *RedisHelpers) CreateValues(e *Entity) (map[string]interface{}, error) {
	t := reflect.TypeOf(e)
	v := reflect.ValueOf(e)
	if t.Kind() != reflect.Struct {
		return nil, errors.New("values can be created from struct only")
	}
	values := make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		// fooT.Field(i) // struct field metadata
		values[t.Field(i).Name] = v.Field(i)
	}

	return values, nil

}

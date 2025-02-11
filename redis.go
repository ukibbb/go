package main

import (
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/google/uuid"
)

func NewRedisHelpers[T Storable]() *RedisHelpers[T] {
	return &RedisHelpers[T]{}
}

type RedisHelpers[T Storable] struct{}

func (h *RedisHelpers[T]) CreateValues(e T) (map[string]string, error) {
	t := reflect.TypeOf(e)
	v := reflect.ValueOf(e)

	// If it's a pointer, get the element it points to
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if t.Kind() != reflect.Struct {
		return nil, errors.New("values can be created from struct only")
	}
	values := make(map[string]string)
	for i := 0; i < t.NumField(); i++ {
		ft := t.Field(i)
		fv := v.Field(i)
		// Check if this field is time.Time
		fName := ft.Name
		if fv.Type() == reflect.TypeOf(time.Time{}) {
			// Get the time value and format it
			timeValue := fv.Interface().(time.Time)
			values[fName] = timeValue.Format(time.RFC3339)
			continue
		}
		if fv.Kind() == reflect.Bool {
			if fv.Bool() {
				values[fName] = "true"
			} else {
				values[fName] = "false"
			}
			continue
		}
		if fv.Type() == reflect.TypeOf(uuid.UUID{}) {
			values[fName] = fv.Interface().(uuid.UUID).String()
			continue
		}
		if fv.Kind() == reflect.String {
			values[fName] = fv.String()
			continue
		}
		values[fName] = "type not supported"
	}

	return values, nil
}

func (*RedisHelpers[T]) RetriveStruct(values map[string]string) (T, error) {
	var e T
	t := reflect.TypeOf(e)
	v := reflect.ValueOf(e)

	if t.Kind() != reflect.Struct {
		return e, errors.New("type in retrive values is not struct")
	}

	for i := 0; i < t.NumField(); i++ {
		ft := t.Field(i)
		fv := v.Field(i)

		name := ft.Name
		v, ok := values[name]
		if !ok {
			continue
		}
		kind := fv.Kind()
		if kind == reflect.String {
			fv.SetString(v)
		}
		if kind == reflect.Bool {
			if v == "true" {
				fv.SetBool(true)
			} else {
				fv.SetBool(false)
			}
		}

		if kind == reflect.Struct {
			if fv.Type() == reflect.TypeOf(time.Time{}) {

			}
			if fv.Type() == reflect.TypeOf(uuid.UUID{}) {

			}
		}

	}

	fmt.Println("entity", e)

	return e, nil

}

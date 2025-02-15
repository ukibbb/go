package main

import (
	"errors"
	"reflect"

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
		//if fv.Type() == reflect.TypeOf(time.Time{}) {
		//	// Get the time value and format it
		//	timeValue := fv.Interface().(time.Time)
		//	values[fName] = timeValue.Format(time.RFC3339)
		//	continue
		//}
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
	// panic: reflect: reflect.Value.SetString using unaddressable value
	// occurs because you are trying to modify a value
	// that is not addressable. In Go,
	// reflection can only modify values that are addressable,
	// meaning you can take their address (e.g., via a pointer).
	//
	var e T // this is not original struct
	t := reflect.TypeOf(e)
	// v := reflect.ValueOf(e) // it returns reflect.Value representing not addrassable copy of e
	// Attempting to modify this copy using fv.SetString or fv.SetBool will result in the panic because the value is not addressable
	// To fix this issue, you need to work with
	// a pointer to the struct instead of the struct itself.
	// This way, the reflect.Value will be addressable, and you can modify its fields.

	v := reflect.ValueOf(&e).Elem()
	// reflect.ValueOf(&e) gets the value of the pointer to e.
	// .Elem() dereferences the pointer to get the underlying struct value, which is now addressable.

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

		// if kind == reflect.Struct {
		//	// if fv.Type() == reflect.TypeOf(time.Time{}) {
		//	// 	pT, err := time.Parse(time.RFC3339, v)
		//	// 	if err != nil {
		//	// 		fv.SetString("Time failed to parse")
		//	// 		continue
		//	// 	}
		//	// 	fv.Set(reflect.ValueOf(pT))
		//	// }
		//}
		if fv.Type() == reflect.TypeOf(uuid.UUID{}) {
			pU, err := uuid.Parse(v)
			if err != nil {
				fv.SetString("UUID failed to parse")
				continue
			}
			fv.Set(reflect.ValueOf(pU))
		}
	}

	return e, nil

}

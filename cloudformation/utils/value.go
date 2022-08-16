package utils

import (
	"encoding/json"
)

type Value[T any] struct {
	Value  T
	Custom string
}

func NewValue[T any](v T) Value[T] {
	return Value[T]{Value: v}
}

func NewPtrValue[T any](v T) *Value[T] {
	return &Value[T]{Value: v}
}

func NewFunc[T any](fn string) Value[T] {
	return Value[T]{Custom: fn}
}

func NewPtrFunc[T any](fn string) *Value[T] {
	return &Value[T]{Custom: fn}
}

func (v *Value[T]) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	if v.Custom != "" {
		return json.Marshal(v.Custom)
	}

	return json.Marshal(v.Value)
}

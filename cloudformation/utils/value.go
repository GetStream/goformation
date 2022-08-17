package utils

import (
	"encoding/json"
)

type Value[T any] struct {
	Value T
	Func  string
}

func NewValue[T any](val T) Value[T] {
	return Value[T]{Value: val}
}

func NewFunction[T any](fun string) Value[T] {
	return Value[T]{Func: fun}
}

func (v *Value[T]) MarshalJSON() ([]byte, error) {
	if v == nil {
		return []byte("null"), nil
	}

	if v.Func != "" {
		return json.Marshal(v.Func)
	}

	return json.Marshal(v.Value)
}

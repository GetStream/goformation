package utils

import (
	"encoding/json"
)

type Value[T any] struct {
	Val  T
	Func string
}

func NewValue[T any](val T) *Value[T] {
	return &Value[T]{Val: val}
}

func NewFunction[T any](fun string) *Value[T] {
	return &Value[T]{Func: fun}
}

func (v *Value[T]) Value() T {
	if v != nil {
		return v.Val
	}

	var t T
	return t
}

func (v *Value[T]) MarshalJSON() ([]byte, error) {
	if v.Func != "" {
		return json.Marshal(v.Func)
	}

	return json.Marshal(v.Val)
}

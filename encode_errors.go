package jsonapi

import "fmt"

type EncodeError struct {
	v       interface{}
	message string
}

func NewEncodeError(v interface{}, message string) error {
	return EncodeError{v: v, message: message}
}

func (e EncodeError) Error() string {
	return fmt.Sprintf("cannot encode %T %v: %s", e.v, e.v, e.message)
}

func NewNotEncodableError(v interface{}) error {
	return NewEncodeError(v, "does not implement jsonapi.Encodable")
}

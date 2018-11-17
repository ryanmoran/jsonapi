package jsonapi

import "fmt"

type DecodeError struct {
	v       interface{}
	message string
}

func NewDecodeError(v interface{}, message string) error {
	return DecodeError{v: v, message: message}
}

func (e DecodeError) Error() string {
	return fmt.Sprintf("cannot decode %T %v: %s", e.v, e.v, e.message)
}

func NewNotDecodableError(v interface{}) error {
	return NewDecodeError(v, "does not implement jsonapi.Decodable")
}

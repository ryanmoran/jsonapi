package jsonapi

import (
	"encoding/json"
	"errors"
	"reflect"
)

type EncodePayload struct {
	m interface{}
}

func NewEncodePayload(m interface{}) EncodePayload {
	return EncodePayload{m}
}

func (ep EncodePayload) MarshalJSON() ([]byte, error) {
	var (
		document EncodeDocument
		err      error
	)

	switch reflect.TypeOf(ep.m).Kind() {
	case reflect.Struct:
		document, err = NewSingularEncodeDocument(ep.m)

	case reflect.Slice:
		if errs, ok := ep.m.(Errors); ok {
			document, err = NewErrorsEncodeDocument(errs)
		} else {
			document, err = NewMultipleEncodeDocument(ep.m)
		}

	default:
		err = errors.New("unknown kind")
	}
	if err != nil {
		panic(err)
	}

	return json.Marshal(document)
}

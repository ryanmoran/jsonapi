package jsonapi

import (
	"encoding/json"
	"reflect"
)

type EncodeDocument struct {
	Data   json.Marshaler `json:"data,omitempty"`
	Errors Errors         `json:"errors,omitempty"`
}

func NewSingularEncodeDocument(m interface{}) (EncodeDocument, error) {
	encodable, err := ToEncodable(m)
	if err != nil {
		panic(err)
	}

	return EncodeDocument{
		Data: NewEncodeResourceObject(encodable),
	}, nil
}

func NewMultipleEncodeDocument(m interface{}) (EncodeDocument, error) {
	var resourceObjects EncodeResourceObjects

	value := reflect.ValueOf(m)
	for i := 0; i < value.Len(); i++ {
		elem := value.Index(i)
		encodable, err := ToEncodable(elem.Interface())
		if err != nil {
			panic(err)
		}

		resourceObjects = append(resourceObjects, NewEncodeResourceObject(encodable))
	}

	return EncodeDocument{
		Data: resourceObjects,
	}, nil
}

func NewErrorsEncodeDocument(errors Errors) (EncodeDocument, error) {
	return EncodeDocument{
		Errors: errors,
	}, nil
}

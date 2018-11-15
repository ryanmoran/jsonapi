package jsonapi

import (
	"encoding/json"
)

type DecodePayload struct {
	v interface{}
}

func NewDecodePayload(v interface{}) DecodePayload {
	return DecodePayload{v}
}

func (dp DecodePayload) UnmarshalJSON(data []byte) error {
	decodable, ok := dp.v.(Decodable)
	if !ok {
		panic("not decodable")
	}

	document := NewDecodeDocument(decodable)
	err := json.Unmarshal(data, &document)
	if err != nil {
		panic(err)
	}

	return nil
}

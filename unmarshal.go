package jsonapi

import "encoding/json"

func Unmarshal(data []byte, v interface{}) error {
	payload := NewDecodePayload(v)
	return json.Unmarshal(data, &payload)
}

package jsonapi

import (
	"encoding/json"
)

func Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(NewEncodePayload(v))
}

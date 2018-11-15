package jsonapi

import "encoding/json"

type EncodeResourceLinkage struct {
	Type string
	ID   string
}

func NewEncodeResourceLinkage(e Encodable) EncodeResourceLinkage {
	return EncodeResourceLinkage{
		Type: e.Type(),
		ID:   e.Primary(),
	}
}

func (erl EncodeResourceLinkage) MarshalJSON() ([]byte, error) {
	linkage := struct {
		Type string `json:"type"`
		ID   string `json:"id"`
	}{
		Type: erl.Type,
		ID:   erl.ID,
	}

	return json.Marshal(linkage)
}

type EncodeResourceLinkages []EncodeResourceLinkage

func (erl EncodeResourceLinkages) MarshalJSON() ([]byte, error) {
	var marshaledObjects []json.RawMessage

	for _, linkage := range erl {
		marshaledObject, err := json.Marshal(linkage)
		if err != nil {
			panic(err)
		}

		marshaledObjects = append(marshaledObjects, marshaledObject)
	}

	return json.Marshal(marshaledObjects)
}

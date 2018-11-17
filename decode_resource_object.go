package jsonapi

import (
	"encoding/json"
	"fmt"
)

type DecodeResourceObject struct {
	d Decodable
}

func NewDecodeResourceObject(d Decodable) DecodeResourceObject {
	return DecodeResourceObject{d}
}

func (dro DecodeResourceObject) UnmarshalJSON(data []byte) error {
	object := struct {
		Type          string              `json:"type"`
		ID            string              `json:"id"`
		Attributes    DecodeAttributes    `json:"attributes"`
		Relationships DecodeRelationships `json:"relationships"`
	}{
		Attributes:    NewDecodeAttributes(dro.d),
		Relationships: NewDecodeRelationships(dro.d),
	}

	err := json.Unmarshal(data, &object)
	if err != nil {
		return err
	}

	if dro.d.Type() != object.Type {
		return NewDecodeError(dro.d, fmt.Sprintf("types %q and %q do not match", dro.d.Type(), object.Type))
	}

	dro.d.SetPrimary(object.ID)

	return nil
}

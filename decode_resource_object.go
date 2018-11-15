package jsonapi

import "encoding/json"

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
		panic(err)
	}

	if dro.d.Type() != object.Type {
		panic("types don't match")
	}

	dro.d.SetPrimary(object.ID)

	return nil
}

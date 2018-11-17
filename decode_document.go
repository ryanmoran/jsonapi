package jsonapi

import "encoding/json"

type DecodeDocument struct {
	d Decodable
}

func NewDecodeDocument(d Decodable) DecodeDocument {
	return DecodeDocument{d}
}

func (dd DecodeDocument) UnmarshalJSON(data []byte) error {
	document := struct {
		Data DecodeResourceObject `json:"data"`
	}{
		Data: NewDecodeResourceObject(dd.d),
	}

	return json.Unmarshal(data, &document)
}

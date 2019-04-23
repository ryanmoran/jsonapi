package jsonapi

import (
	"encoding/json"
)

type EncodeResourceObject struct {
	Type          string
	ID            string
	Attributes    EncodeAttributes
	Links         EncodeLinks
	Relationships EncodeRelationships
	Meta          EncodeMeta
}

func NewEncodeResourceObject(encodable Encodable) EncodeResourceObject {
	return EncodeResourceObject{
		Type:          encodable.Type(),
		ID:            encodable.Primary(),
		Attributes:    NewEncodeAttributes(encodable),
		Links:         NewEncodeLinks(encodable),
		Relationships: NewEncodeRelationships(encodable),
		Meta:          NewEncodeMeta(encodable),
	}
}

func (ero EncodeResourceObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Type          string              `json:"type"`
		ID            string              `json:"id"`
		Attributes    EncodeAttributes    `json:"attributes,omitempty"`
		Links         EncodeLinks         `json:"links,omitempty"`
		Relationships EncodeRelationships `json:"relationships,omitempty"`
		Meta          EncodeMeta          `json:"meta,omitempty"`
	}{
		Type:          ero.Type,
		ID:            ero.ID,
		Attributes:    ero.Attributes,
		Links:         ero.Links,
		Relationships: ero.Relationships,
		Meta:          ero.Meta,
	})
}

type EncodeResourceObjects []EncodeResourceObject

func (ero EncodeResourceObjects) MarshalJSON() ([]byte, error) {
	marshaledObjects := []json.RawMessage{}

	for _, resourceObject := range ero {
		marshaledObject, err := json.Marshal(resourceObject)
		if err != nil {
			return nil, err
		}

		marshaledObjects = append(marshaledObjects, marshaledObject)
	}

	return json.Marshal(marshaledObjects)
}

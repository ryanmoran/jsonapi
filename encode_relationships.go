package jsonapi

import (
	"encoding/json"
)

type EncodeRelationships map[string]RelationshipDocument

func NewEncodeRelationships(m interface{}) EncodeRelationships {
	relatable, ok := m.(Relatable)
	if !ok {
		return nil
	}

	relationships := EncodeRelationships{}
	multiRelationships := map[string][]Encodable{}

	for _, relationship := range relatable.Relationships() {
		if relationship.Type == SingularRelationship {
			relationships[relationship.Name] = RelationshipDocument{
				Data: NewEncodeResourceLinkage(relationship.Resource),
			}
		} else {
			multiRelationships[relationship.Name] = append(multiRelationships[relationship.Name], relationship.Resource)
		}
	}

	for name, encodables := range multiRelationships {
		var linkages EncodeResourceLinkages
		for _, encodable := range encodables {
			linkages = append(linkages, NewEncodeResourceLinkage(encodable))
		}

		relationships[name] = RelationshipDocument{Data: linkages}
	}

	return relationships
}

type RelationshipDocument struct {
	Data json.Marshaler `json:"data"`
}

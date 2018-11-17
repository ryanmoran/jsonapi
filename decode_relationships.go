package jsonapi

import (
	"encoding/json"
	"fmt"
)

type DecodeRelationships struct {
	d Decodable
}

func NewDecodeRelationships(d Decodable) DecodeRelationships {
	return DecodeRelationships{d}
}

func (dr DecodeRelationships) UnmarshalJSON(data []byte) error {
	assignable, ok := dr.d.(RelationshipsAssignable)
	if !ok {
		return nil
	}

	var relationshipsMap map[string]interface{}

	err := json.Unmarshal(data, &relationshipsMap)
	if err != nil {
		return err
	}

	var relationships []Relationship
	for name, resource := range relationshipsMap {
		switch r := resource.(type) {
		case map[string]interface{}:
			rType, _ := r["type"].(string)
			rID, _ := r["id"].(string)

			relationships = append(relationships, Relationship{
				Name: name,
				Type: SingularRelationship,
				Resource: DecodeResourceLinkage{
					ResourceType: rType,
					ID:           rID,
				},
			})
		case []interface{}:
			for _, item := range r {
				s, ok := item.(map[string]interface{})
				if !ok {
					return NewDecodeError(dr.d, fmt.Sprintf("relationship %s is not an array of objects", name))
				}

				sType, _ := s["type"].(string)
				sID, _ := s["id"].(string)

				relationships = append(relationships, Relationship{
					Name: name,
					Type: MultiRelationship,
					Resource: DecodeResourceLinkage{
						ResourceType: sType,
						ID:           sID,
					},
				})
			}
		}
	}

	assignable.AssignRelationships(relationships)

	return nil
}

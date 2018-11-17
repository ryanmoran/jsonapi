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
		r, ok := resource.(map[string]interface{})
		if !ok {
			return NewDecodeError(dr.d, fmt.Sprintf("relationship %s is not an object", name))
		}

		switch s := r["data"].(type) {
		case map[string]interface{}:
			sType, _ := s["type"].(string)
			sID, _ := s["id"].(string)

			relationships = append(relationships, Relationship{
				Name: name,
				Type: SingularRelationship,
				Resource: DecodeResourceLinkage{
					ResourceType: sType,
					ID:           sID,
				},
			})
		case []interface{}:
			for _, item := range s {
				t, ok := item.(map[string]interface{})
				if !ok {
					return NewDecodeError(dr.d, fmt.Sprintf("relationship %s data is not an array of objects", name))
				}

				tType, _ := t["type"].(string)
				tID, _ := t["id"].(string)

				relationships = append(relationships, Relationship{
					Name: name,
					Type: MultiRelationship,
					Resource: DecodeResourceLinkage{
						ResourceType: tType,
						ID:           tID,
					},
				})
			}
		}
	}

	assignable.AssignRelationships(relationships)

	return nil
}

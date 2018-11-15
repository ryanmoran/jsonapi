package jsonapi_test

import (
	"testing"

	"github.com/ryanmoran/jsonapi"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestJsonapi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "jsonapi")
}

type NotAResourcePayload struct{}

type SimplePayload struct {
	ID string
}

func (sp SimplePayload) Type() string {
	return "simple-payload"
}

func (sp SimplePayload) Primary() string {
	return sp.ID
}

func (sp *SimplePayload) SetPrimary(id string) {
	sp.ID = id
}

type AttributesPayload struct {
	ID       string
	SomeAttr string `jsonapi:"some-attr"`
}

func (ap AttributesPayload) Type() string {
	return "attributes-payload"
}

func (ap AttributesPayload) Primary() string {
	return ap.ID
}

func (ap *AttributesPayload) SetPrimary(id string) {
	ap.ID = id
}

type LinksPayload struct {
	ID string
}

func (lp LinksPayload) Type() string {
	return "links-payload"
}

func (lp LinksPayload) Primary() string {
	return lp.ID
}

func (lp LinksPayload) Links() []jsonapi.Link {
	return []jsonapi.Link{
		{Name: "some-link", Href: "some-href"},
	}
}

type UnmarshalablePayload struct {
	ID       string
	SomeFunc func() `jsonapi:"some-func"`
}

func (up UnmarshalablePayload) Type() string {
	return "unmarshalable-payload"
}

func (up UnmarshalablePayload) Primary() string {
	return up.ID
}

type RelationshipsPayload struct {
	ID               string
	SingleRelationID string
	MultiRelationIDs []string
}

func (rp RelationshipsPayload) Type() string {
	return "relationships-payload"
}

func (rp RelationshipsPayload) Primary() string {
	return rp.ID
}

func (rp *RelationshipsPayload) SetPrimary(id string) {
	rp.ID = id
}

func (rp RelationshipsPayload) Relationships() []jsonapi.Relationship {
	relationships := []jsonapi.Relationship{
		{
			Name:     "single-relation",
			Type:     jsonapi.SingularRelationship,
			Resource: SimplePayload{ID: rp.SingleRelationID},
		},
	}

	for _, id := range rp.MultiRelationIDs {
		relationships = append(relationships, jsonapi.Relationship{
			Name:     "multi-relation",
			Type:     jsonapi.MultiRelationship,
			Resource: SimplePayload{ID: id},
		})
	}

	return relationships
}

func (rp *RelationshipsPayload) AssignRelationships(relationships []jsonapi.Relationship) {
	for _, relationship := range relationships {
		if relationship.Name == "single-relation" {
			rp.SingleRelationID = relationship.Resource.Primary()
		}

		if relationship.Name == "multi-relation" {
			rp.MultiRelationIDs = append(rp.MultiRelationIDs, relationship.Resource.Primary())
		}
	}
}

type NotARelationshipPayload struct {
	ID string
}

func (nrp NotARelationshipPayload) Type() string {
	return "not-a-relationships-payload"
}

func (nrp NotARelationshipPayload) Primary() string {
	return nrp.ID
}

func (nrp NotARelationshipPayload) Relationships() []jsonapi.Relationship {
	return []jsonapi.Relationship{{Name: "not-a-relation"}}
}
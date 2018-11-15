package jsonapi

type RelationshipType int

const (
	Unknown RelationshipType = iota
	SingularRelationship
	MultiRelationship
)

type Relatable interface {
	Relationships() []Relationship
}

type RelationshipsAssignable interface {
	AssignRelationships([]Relationship)
}

type Relationship struct {
	Name     string
	Type     RelationshipType
	Resource Encodable
}

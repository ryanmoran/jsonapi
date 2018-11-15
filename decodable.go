package jsonapi

type Decodable interface {
	Type() string
	SetPrimary(id string)
}

package jsonapi

type Meta map[string]interface{}

type Metable interface {
	Meta() Meta
}

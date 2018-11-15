package jsonapi

type Encodable interface {
	Type() string
	Primary() string
}

func ToEncodable(v interface{}) (Encodable, error) {
	m, ok := v.(Encodable)
	if !ok {
		panic("cannot encode")
	}

	return m, nil
}

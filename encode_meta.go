package jsonapi

type EncodeMeta map[string]interface{}

func NewEncodeMeta(m interface{}) EncodeMeta {
	metable, ok := m.(Metable)
	if !ok {
		return nil
	}

	encodeMeta := EncodeMeta{}
	for key, value := range metable.Meta() {
		encodeMeta[key] = value
	}

	return encodeMeta
}

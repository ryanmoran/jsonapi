package jsonapi

type DecodeResourceLinkage struct {
	ResourceType string
	ID           string
}

func (drl DecodeResourceLinkage) Type() string {
	return drl.ResourceType
}

func (drl DecodeResourceLinkage) Primary() string {
	return drl.ID
}

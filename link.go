package jsonapi

type Linkable interface {
	Links() []Link
}

type Link struct {
	Name string `json:"-"`
	Href string `json:"href"`
}

package jsonapi

type EncodeLinks map[string]Link

func NewEncodeLinks(m interface{}) EncodeLinks {
	linkable, ok := m.(Linkable)
	if !ok {
		return nil
	}

	links := EncodeLinks{}
	for _, link := range linkable.Links() {
		links[link.Name] = link
	}

	return links
}

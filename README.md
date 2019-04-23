# JSONAPI Document Serialization for Go
This library surfaces a simple API for serializing/deserializing JSONAPI documents.

## Examples
### Serialize a simple resource object
```go
package main

import (
	"fmt"
	"os"

	"github.com/ryanmoran/jsonapi"
)

type Article struct {
	ID    string
	Title string `jsonapi:"title"`
	Body  string `jsonapi:"body"`
}

func (a Article) Type() string {
	return "article"
}

func (a Article) Primary() string {
	return a.ID
}

func main() {
	content, err := jsonapi.Marshal(Article{
		ID:    "1",
		Title: "JSON:API paints my bikeshed!",
		Body:  "The shortest article. Ever.",
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(content))
}
```

```bash
go run main.go
{
  "data": {
    "type": "article",
    "id": "1",
    "attributes": {
      "body": "The shortest article. Ever.",
      "title": "JSON:API paints my bikeshed!"
    }
  }
}
```

### Deserialize a simple resource object
```go
package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ryanmoran/jsonapi"
)

type Article struct {
	ID       string
	Title    string `jsonapi:"title"`
	Body     string `jsonapi:"body"`
	AuthorID string
}

func (a Article) Type() string {
	return "article"
}

func (a Article) Primary() string {
	return a.ID
}

func (a *Article) SetPrimary(id string) {
	a.ID = id
}

func main() {
	content, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var article Article
	err = jsonapi.Unmarshal(content, &article)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%#v\n", article)
}
```

```bash
echo '{
  "data": {
    "type": "article",
    "id": "1",
    "attributes": {
      "body": "The shortest article. Ever.",
      "title": "JSON:API paints my bikeshed!"
    }
  }
}' | go run main.go
main.Article{ID:"1", Title:"JSON:API paints my bikeshed!", Body:"The shortest article. Ever.", AuthorID:""}
```

### Serialize a list of simple resource objects
```go
package main

import (
	"fmt"
	"os"

	"github.com/ryanmoran/jsonapi"
)

type Article struct {
	ID    string
	Title string `jsonapi:"title"`
	Body  string `jsonapi:"body"`
}

func (a Article) Type() string {
	return "article"
}

func (a Article) Primary() string {
	return a.ID
}

func main() {
	content, err := jsonapi.Marshal([]Article{
		{
			ID:    "1",
			Title: "JSON:API paints my bikeshed!",
			Body:  "The shortest article. Ever.",
		},
		{
			ID:    "2",
			Title: "Following up on that short article...",
			Body:  "This article is also pretty short.",
		},
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(content))
}
```

```bash
go run main.go
{
  "data": [
    {
      "type": "article",
      "id": "1",
      "attributes": {
        "body": "The shortest article. Ever.",
        "title": "JSON:API paints my bikeshed!"
      }
    },
    {
      "type": "article",
      "id": "2",
      "attributes": {
        "body": "This article is also pretty short.",
        "title": "Following up on that short article..."
      }
    }
  ]
}
```

### Serialize a resource object with links, relationships, and meta
```go
package main

import (
	"fmt"
	"os"

	"github.com/ryanmoran/jsonapi"
)

type Article struct {
	ID       string
	Title    string `jsonapi:"title"`
	Body     string `jsonapi:"body"`
	AuthorID string
}

func (a Article) Type() string {
	return "article"
}

func (a Article) Primary() string {
	return a.ID
}

func (a Article) Links() []jsonapi.Link {
	return []jsonapi.Link{
		{
			Name: "self",
			Href: "/articles/1",
		},
	}
}

func (a Article) Relationships() []jsonapi.Relationship {
	return []jsonapi.Relationship{
		{
			Name: "author",
			Type: jsonapi.SingularRelationship,
			Resource: Person{
				ID: "42",
			},
		},
	}
}

func (a Article) Meta() jsonapi.Meta {
  return jsonapi.Meta{
    "views": 100
  }
}

type Person struct {
	ID string
}

func (p Person) Type() string {
	return "person"
}

func (p Person) Primary() string {
	return p.ID
}

func main() {
	content, err := jsonapi.Marshal(Article{
		ID:       "1",
		Title:    "JSON:API paints my bikeshed!",
		Body:     "The shortest article. Ever.",
		AuthorID: "42",
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(content))
}
```

```bash
go run main.go
{
  "data": {
    "type": "article",
    "id": "1",
    "attributes": {
      "body": "The shortest article. Ever.",
      "title": "JSON:API paints my bikeshed!"
    },
    "links": {
      "self": {
        "href": "/articles/1"
      }
    },
    "relationships": {
      "author": {
        "data": {
          "type": "person",
          "id": "42"
        }
      }
    },
    "meta": {
      "views": 100
    }
  }
}
```

### Deserialize a resource object with relationships

```go
package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ryanmoran/jsonapi"
)

type Article struct {
	ID       string
	Title    string `jsonapi:"title"`
	Body     string `jsonapi:"body"`
	AuthorID string
}

func (a Article) Type() string {
	return "article"
}

func (a Article) Primary() string {
	return a.ID
}

func (a *Article) SetPrimary(id string) {
	a.ID = id
}

func (a *Article) AssignRelationships(relationships []jsonapi.Relationship) {
	for _, relationship := range relationships {
		if relationship.Name == "author" {
			a.AuthorID = relationship.Resource.Primary()
		}
	}
}

func main() {
	content, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var article Article
	err = jsonapi.Unmarshal(content, &article)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%#v\n", article)
}
```

```bash
echo '{
  "data": {
    "type": "article",
    "id": "1",
    "attributes": {
      "body": "The shortest article. Ever.",
      "title": "JSON:API paints my bikeshed!"
    },
    "links": {
      "self": {
        "href": "/articles/1"
      }
    },
    "relationships": {
      "author": {
        "data": {
          "type": "person",
          "id": "42"
        }
      }
    }
  }
}
' | go run main.go
main.Article{ID:"1", Title:"JSON:API paints my bikeshed!", Body:"The shortest article. Ever.", AuthorID:"42"}
```

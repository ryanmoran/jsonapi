package jsonapi_test

import (
	"github.com/ryanmoran/jsonapi"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Constants", func() {
	It("defines a content type", func() {
		Expect(jsonapi.ContentType).To(Equal("application/vnd.api+json"))
	})
})

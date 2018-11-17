package jsonapi_test

import (
	"github.com/ryanmoran/jsonapi"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DecodeError", func() {
	Describe("Error", func() {
		It("returns a string describing the error", func() {
			err := jsonapi.NewDecodeError(123, "some-message")
			Expect(err.Error()).To(Equal("cannot decode int 123: some-message"))
		})
	})
})

var _ = Describe("NotDecodableError", func() {
	Describe("Error", func() {
		It("returns a string describing the error", func() {
			err := jsonapi.NewNotDecodableError(123)
			Expect(err.Error()).To(Equal("cannot decode int 123: does not implement jsonapi.Decodable"))
		})
	})
})

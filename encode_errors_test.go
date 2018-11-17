package jsonapi_test

import (
	"github.com/ryanmoran/jsonapi"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("EncodeError", func() {
	Describe("Error", func() {
		It("returns a string describing the error", func() {
			err := jsonapi.NewEncodeError(123, "some-message")
			Expect(err.Error()).To(Equal("cannot encode int 123: some-message"))
		})
	})
})

var _ = Describe("NotEncodableError", func() {
	Describe("Error", func() {
		It("returns a string describing the error", func() {
			err := jsonapi.NewNotEncodableError(123)
			Expect(err.Error()).To(Equal("cannot encode int 123: does not implement jsonapi.Encodable"))
		})
	})
})

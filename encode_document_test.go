package jsonapi_test

import (
	"github.com/ryanmoran/jsonapi"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("EncodeDocument", func() {
	Describe("NewSingularEncodeDocument", func() {
		Context("failure cases", func() {
			Context("when the payload is not encodable", func() {
				It("returns an error", func() {
					_, err := jsonapi.NewSingularEncodeDocument("some-payload")
					Expect(err).To(MatchError("cannot encode string some-payload: does not implement jsonapi.Encodable"))
				})
			})
		})
	})

	Describe("NewMultipleEncodeDocument", func() {
		Context("failure cases", func() {
			Context("when the payload is not encodable", func() {
				It("returns an error", func() {
					_, err := jsonapi.NewMultipleEncodeDocument([]string{"some-payload"})
					Expect(err).To(MatchError("cannot encode string some-payload: does not implement jsonapi.Encodable"))
				})
			})
		})
	})
})

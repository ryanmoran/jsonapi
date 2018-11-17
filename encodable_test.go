package jsonapi_test

import (
	"github.com/ryanmoran/jsonapi"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Encodable", func() {
	Describe("ToEncodable", func() {
		Context("failure case", func() {
			Context("when passed an non-encodable type", func() {
				It("returns an error", func() {
					_, err := jsonapi.ToEncodable("not-encodable")
					Expect(err).To(MatchError("cannot encode string not-encodable: does not implement jsonapi.Encodable"))
				})
			})
		})
	})
})

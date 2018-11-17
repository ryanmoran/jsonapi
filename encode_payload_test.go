package jsonapi_test

import (
	"github.com/ryanmoran/jsonapi"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("EncodePayload", func() {
	Describe("MarshalJSON", func() {
		Context("failure cases", func() {
			Context("when the payload is a non-encodable type", func() {
				It("returns an error", func() {
					payload := jsonapi.NewEncodePayload("not-encodable")
					_, err := payload.MarshalJSON()
					Expect(err).To(MatchError("cannot encode string not-encodable: payloads must be of type struct or slice"))
				})
			})
		})
	})
})

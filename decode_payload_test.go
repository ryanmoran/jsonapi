package jsonapi_test

import (
	"github.com/ryanmoran/jsonapi"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DecodePayload", func() {
	Describe("UnmarshalJSON", func() {
		Context("failure cases", func() {
			Context("the the payload is not decodable", func() {
				It("returns a NotDecodableError", func() {
					payload := jsonapi.NewDecodePayload("not-decodable")
					err := payload.UnmarshalJSON([]byte{})
					Expect(err).To(MatchError(jsonapi.NewNotDecodableError("not-decodable")))
				})
			})

			Context("when the JSON cannot be unmarshaled", func() {
				It("returns an error", func() {
					var simplePayload SimplePayload
					payload := jsonapi.NewDecodePayload(&simplePayload)
					err := payload.UnmarshalJSON([]byte("%%%"))
					Expect(err).To(MatchError(ContainSubstring("invalid character '%'")))
				})
			})
		})
	})
})

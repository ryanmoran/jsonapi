package jsonapi_test

import (
	"github.com/ryanmoran/jsonapi"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DecodeAttributes", func() {
	Describe("UnmarshalJSON", func() {
		Context("failure cases", func() {
			Context("when the JSON cannot be unmarshaled", func() {
				It("returns an error", func() {
					var payload AttributesPayload
					attributes := jsonapi.NewDecodeAttributes(&payload)
					err := attributes.UnmarshalJSON([]byte("%%%"))
					Expect(err).To(MatchError(ContainSubstring("invalid character '%'")))
				})
			})

			Context("when the attribute field types do not match", func() {
				It("returns an error", func() {
					var payload AttributesPayload
					attributes := jsonapi.NewDecodeAttributes(&payload)
					err := attributes.UnmarshalJSON([]byte(`{ "some-attr": 123 }`))
					Expect(err).To(MatchError("json: cannot unmarshal number into Go value of type string"))
				})
			})
		})
	})
})

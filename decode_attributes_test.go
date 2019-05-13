package jsonapi_test

import (
	"encoding/json"

	"github.com/ryanmoran/jsonapi"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DecodeAttributes", func() {
	Describe("UnmarshalJSON", func() {
		Context("when an attribute is json.RawMessage", func() {
			It("decodes the attribute correctly", func() {
				var payload RawMessageAttributesPayload

				attributes := jsonapi.NewDecodeAttributes(&payload)
				err := attributes.UnmarshalJSON([]byte(`{
					"some-attr": {
						"some-key": "some-value"
					}
				}`))
				Expect(err).NotTo(HaveOccurred())
				Expect(attributes).To(Equal(jsonapi.NewDecodeAttributes(&RawMessageAttributesPayload{
					SomeAttr: json.RawMessage(`{"some-key":"some-value"}`),
				})))
			})
		})

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
					err := attributes.UnmarshalJSON([]byte(`{
						"some-attr": 123
					}`))
					Expect(err).To(MatchError("cannot decode *jsonapi_test.AttributesPayload &{ }: field \"some-attr\" types \"string\" and \"float64\" do not match"))
				})
			})
		})
	})
})

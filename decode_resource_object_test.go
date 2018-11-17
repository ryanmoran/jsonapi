package jsonapi_test

import (
	"github.com/ryanmoran/jsonapi"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DecodeResourceObject", func() {
	Describe("UnmarshalJSON", func() {
		Context("failure cases", func() {
			Context("when the JSON cannot be unmarshaled", func() {
				It("returns an error", func() {
					var payload SimplePayload
					object := jsonapi.NewDecodeResourceObject(&payload)
					err := object.UnmarshalJSON([]byte("%%%"))
					Expect(err).To(MatchError(ContainSubstring("invalid character '%'")))
				})
			})
		})

		Context("when the type field and decodeable type don't match", func() {
			It("returns an error", func() {
				var payload SimplePayload
				object := jsonapi.NewDecodeResourceObject(&payload)
				err := object.UnmarshalJSON([]byte(`{
					"type": "invalid-type",
					"id": "some-id"
				}`))
				Expect(err).To(MatchError("cannot decode *jsonapi_test.SimplePayload &{}: types \"simple-payload\" and \"invalid-type\" do not match"))
			})
		})
	})
})

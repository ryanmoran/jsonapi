package jsonapi_test

import (
	"github.com/ryanmoran/jsonapi"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DecodeRelationships", func() {
	Describe("UnmarshalJSON", func() {
		Context("failure cases", func() {
			Context("when the JSON cannot be unmarshaled", func() {
				It("returns an error", func() {
					var payload RelationshipsPayload
					relationships := jsonapi.NewDecodeRelationships(&payload)
					err := relationships.UnmarshalJSON([]byte("%%%"))
					Expect(err).To(MatchError(ContainSubstring("invalid character '%'")))
				})
			})

			Context("when the JSON has a slice that is not of type []map[string]interface{}", func() {
				It("returns an error", func() {
					var payload RelationshipsPayload
					relationships := jsonapi.NewDecodeRelationships(&payload)
					err := relationships.UnmarshalJSON([]byte(`{
						"some-relationship": ["not-a-map"]
					}`))
					Expect(err).To(MatchError("cannot decode *jsonapi_test.RelationshipsPayload &{  []}: relationship some-relationship is not an array of objects"))
				})
			})
		})
	})
})

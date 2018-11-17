package jsonapi_test

import (
	"github.com/ryanmoran/jsonapi"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("DecodeDocument", func() {
	Describe("UnmarshalJSON", func() {
		Context("failure cases", func() {
			Context("when the JSON cannot be unmarshaled", func() {
				It("returns an error", func() {
					var payload SimplePayload
					document := jsonapi.NewDecodeDocument(&payload)
					err := document.UnmarshalJSON([]byte("%%%"))
					Expect(err).To(MatchError(ContainSubstring("invalid character '%'")))
				})
			})
		})
	})
})

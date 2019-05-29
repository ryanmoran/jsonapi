package jsonapi_test

import (
	"encoding/json"

	"github.com/ryanmoran/jsonapi"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Unmarshal", func() {
	It("unmarshals a simple payload", func() {
		var payload SimplePayload
		err := jsonapi.Unmarshal([]byte(`{
			"data": {
				"type": "simple-payload",
				"id": "some-id"
			}
		}`), &payload)
		Expect(err).NotTo(HaveOccurred())

		Expect(payload).To(Equal(SimplePayload{ID: "some-id"}))
	})

	It("unmarshals a payload with attributes", func() {
		var payload AttributesPayload
		err := jsonapi.Unmarshal([]byte(`{
			"data": {
				"type": "attributes-payload",
				"id": "some-id",
				"attributes": {
					"some-attr": "some-value"
				}
			}
		}`), &payload)
		Expect(err).NotTo(HaveOccurred())

		Expect(payload).To(Equal(AttributesPayload{
			ID:       "some-id",
			SomeAttr: "some-value",
		}))
	})

	It("unmarshals a payload with complex attributes", func() {
		var payload ComplexAttributesPayload
		err := jsonapi.Unmarshal([]byte(`{
			"data": {
				"type": "complex-attributes-payload",
				"id": "some-id",
				"attributes": {
					"some-attr": {
						"name": "some-name"
					}
				}
			}
		}`), &payload)
		Expect(err).NotTo(HaveOccurred())

		Expect(payload).To(Equal(ComplexAttributesPayload{
			ID: "some-id",
			SomeAttr: ComplexAttributesPayloadAttribute{
				Name: "some-name",
			},
		}))
	})

	It("unmarshals a payload with raw message attributes", func() {
		var payload RawMessageAttributesPayload
		err := jsonapi.Unmarshal([]byte(`{
			"data": {
				"type": "raw-message-attributes-payload",
				"id": "some-id",
				"attributes": {
					"some-attr": { "some-key": "some-value" }
				}
			}
		}`), &payload)
		Expect(err).NotTo(HaveOccurred())

		Expect(payload).To(Equal(RawMessageAttributesPayload{
			ID:       "some-id",
			SomeAttr: json.RawMessage(`{ "some-key": "some-value" }`),
		}))
	})

	It("unmarshals a payload with relationships", func() {
		var payload RelationshipsPayload
		err := jsonapi.Unmarshal([]byte(`{
			"data": {
				"type": "relationships-payload",
				"id": "some-id",
				"relationships": {
					"single-relation": {
						"data": {
							"type": "simple-payload",
							"id": "single-relation-id"
						}
					},
					"multi-relation": {
						"data": [
							{
								"type": "simple-payload",
								"id": "multi-relation-id"
							}
						]
					}
				}
			}
		}`), &payload)
		Expect(err).NotTo(HaveOccurred())

		Expect(payload).To(Equal(RelationshipsPayload{
			ID:               "some-id",
			SingleRelationID: "single-relation-id",
			MultiRelationIDs: []string{"multi-relation-id"},
		}))
	})
})

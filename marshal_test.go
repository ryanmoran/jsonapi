package jsonapi_test

import (
	"net/http"

	"github.com/ryanmoran/jsonapi"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Marshal", func() {
	It("marshals a simple payload", func() {
		document, err := jsonapi.Marshal(SimplePayload{ID: "some-id"})
		Expect(err).NotTo(HaveOccurred())

		Expect(document).To(MatchJSON(`{
			"data": {
				"type": "simple-payload",
				"id": "some-id"
			}
		}`))
	})

	It("marshals a list of simple payloads", func() {
		document, err := jsonapi.Marshal([]SimplePayload{
			{ID: "some-id"},
			{ID: "other-id"},
		})
		Expect(err).NotTo(HaveOccurred())

		Expect(document).To(MatchJSON(`{
			"data": [
				{
					"type": "simple-payload",
					"id": "some-id"
				},
				{
					"type": "simple-payload",
					"id": "other-id"
				}
			]
		}`))
	})

	Context("attributes", func() {
		It("marshals a payload with attributes", func() {
			document, err := jsonapi.Marshal(AttributesPayload{
				ID:       "some-id",
				SomeAttr: "some-value",
			})
			Expect(err).NotTo(HaveOccurred())

			Expect(document).To(MatchJSON(`{
				"data": {
					"type": "attributes-payload",
					"id": "some-id",
					"attributes": {
						"some-attr": "some-value"
					}
				}
			}`))
		})
	})

	Context("links", func() {
		It("marshals a payload with links", func() {
			document, err := jsonapi.Marshal(LinksPayload{ID: "some-id"})
			Expect(err).NotTo(HaveOccurred())

			Expect(document).To(MatchJSON(`{
				"data": {
					"type": "links-payload",
					"id": "some-id",
					"links": {
						"some-link": {
							"href": "some-href"
						}
					}
				}
			}`))
		})
	})

	Context("relationships", func() {
		It("marshals a payload with relationships", func() {
			document, err := jsonapi.Marshal(RelationshipsPayload{
				ID:               "some-id",
				SingleRelationID: "single-relation-id",
				MultiRelationIDs: []string{"multi-relation-id"},
			})
			Expect(err).NotTo(HaveOccurred())

			Expect(document).To(MatchJSON(`{
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
			}`))
		})
	})

	Context("error objects", func() {
		It("marshals an error payload", func() {
			document, err := jsonapi.Marshal(jsonapi.Errors{
				{
					ID:     "some-id",
					Status: http.StatusTeapot,
					Code:   "some-code",
					Title:  "some-title",
					Detail: "some-detail",
				},
			})
			Expect(err).NotTo(HaveOccurred())
			Expect(document).To(MatchJSON(`{
				"errors": [
					{
						"id": "some-id",
						"status": "418",
						"code": "some-code",
						"title": "some-title",
						"detail": "some-detail"
					}
				]
			}`))
		})

		It("does not require any error fields", func() {
			document, err := jsonapi.Marshal(jsonapi.Errors{{}})
			Expect(err).NotTo(HaveOccurred())
			Expect(document).To(MatchJSON(`{ "errors": [{}] }`))
		})
	})
})

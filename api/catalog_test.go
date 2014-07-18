package api_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/pivotal-cf/go-service-broker/api/matchers"

	"github.com/pivotal-cf/go-service-broker/api"
)

var _ = Describe("Catalog", func() {
	Describe("Service", func() {
		Describe("JSON encoding", func() {
			It("uses the correct keys", func() {
				service := api.Service{
					ID:          "ID-1",
					Name:        "Cassandra",
					Description: "A Cassandra Plan",
					Bindable:    true,
					Plans:       []api.ServicePlan{},
					Metadata:    api.ServiceMetadata{},
					Tags:        []string{},
				}
				json := `{"id":"ID-1","name":"Cassandra","description":"A Cassandra Plan","bindable":true,"plans":[],"metadata":{"displayName":"","longDescription":"","documentationUrl":"","supportUrl":"","listing":{"blurb":"","imageUrl":""},"provider":{"name":""}},"tags":[]}`

				Expect(service).To(MarshalToJSON(json))
			})
		})
	})

	Describe("ServicePlan", func() {
		Describe("JSON encoding", func() {
			It("uses the correct keys", func() {
				plan := api.ServicePlan{
					ID:          "ID-1",
					Name:        "Cassandra",
					Description: "A Cassandra Plan",
					Metadata: api.ServicePlanMetadata{
						Bullets: []string{},
						Costs:   []api.ServicePlanMetadataCost{},
					},
				}
				json := `{"id":"ID-1","name":"Cassandra","description":"A Cassandra Plan","metadata":{"bullets":[],"displayName":"","costs":[]}}`

				Expect(plan).To(MarshalToJSON(json))
			})
		})
	})

	Describe("ServicePlanMetadata", func() {
		Describe("JSON encoding", func() {
			It("uses the correct keys", func() {
				metadata := api.ServicePlanMetadata{
					Bullets:     []string{},
					DisplayName: "Some display name",
					Costs:       []api.ServicePlanMetadataCost{},
				}
				json := `{"bullets":[],"displayName":"Some display name","costs":[]}`

				Expect(metadata).To(MarshalToJSON(json))
			})
		})
	})

	Describe("ServicePlanMetadataCost", func() {
		Describe("JSON encoding", func() {
			It("uses the correct keys", func() {
				cost := api.ServicePlanMetadataCost{
					Unit: "Free",
				}
				json := `{"unit":"Free"}`

				Expect(cost).To(MarshalToJSON(json))
			})
		})
	})

	Describe("ServiceMetadata", func() {
		Describe("JSON encoding", func() {
			It("uses the correct keys", func() {
				metadata := api.ServiceMetadata{
					DisplayName:      "Cassandra",
					LongDescription:  "A long description of Cassandra",
					DocumentationUrl: "",
					SupportUrl:       "",
					Listing:          api.ServiceMetadataListing{},
					Provider:         api.ServiceMetadataProvider{},
				}
				json := `{"displayName":"Cassandra","longDescription":"A long description of Cassandra","documentationUrl":"","supportUrl":"","listing":{"blurb":"","imageUrl":""},"provider":{"name":""}}`

				Expect(metadata).To(MarshalToJSON(json))
			})
		})
	})

	Describe("ServiceMetadataListing", func() {
		Describe("JSON encoding", func() {
			It("uses the correct keys", func() {
				listing := api.ServiceMetadataListing{
					Blurb:    "Blurb",
					ImageUrl: "foo",
				}
				json := `{"blurb":"Blurb","imageUrl":"foo"}`

				Expect(listing).To(MarshalToJSON(json))
			})
		})
	})

	Describe("ServiceMetadataProvider", func() {
		Describe("JSON encoding", func() {
			It("uses the correct keys", func() {
				provider := api.ServiceMetadataProvider{
					Name: "Pivotal",
				}
				json := `{"name":"Pivotal"}`

				Expect(provider).To(MarshalToJSON(json))
			})
		})
	})
})

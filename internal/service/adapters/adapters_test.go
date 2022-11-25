package adapters

import (
	"testing"

	orgscraper "green-scraper/internal/model"

	"github.com/stretchr/testify/assert"
)

// Test_ToOrganization TODO: figure out how to assign to custom fields and sectors (nested json/structs is gross)
func Test_ToOrganization(t *testing.T) {
	type testCase struct {
		name        string
		inputJSON   *orgscraper.JSONResponse
		expectedOrg *orgscraper.Organization
	}
	input := &orgscraper.JSONResponse{}
	input.Props.PageProps.Company.ID = 123
	input.Props.PageProps.Company.Name = "Test"
	input.Props.PageProps.Company.Description = "\u003cp\u003eWe are developing a new class of cost-effective, multi-day energy storage systems that will enable a reliable and fully-renewable electric grid year-round.\u003c/p\u003e"
	cases := []testCase{
		{
			name:      "should convert an acquired JSON response into an organization",
			inputJSON: input,
			expectedOrg: &orgscraper.Organization{
				SiteID:       123,
				Name:         "Test",
				Description:  "We are developing a new class of cost-effective, multi-day energy storage systems that will enable a reliable and fully-renewable electric grid year-round.",
				Tags:         []string{},
				Size:         "",
				Location:     "",
				URL:          "",
				ReferenceURL: "",
				Logo:         "",
				PostingCount: 0,
				Outreach:     orgscraper.Contact{},
				Active:       false,
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			out := toOrganization(c.inputJSON)
			assert.Equal(t, c.expectedOrg, out)
		})
	}
}

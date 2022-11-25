package adapters

import (
	"fmt"
	"strconv"
	"strings"

	"green-scraper/internal/model"
)

const (
	ContactEmail = "What is your work email? (ex: Jane@company.com)"
	OrgSize      = "Organization Size"
)

func ToOrganization(r *model.JSONResponse) *model.Organization {
	var email string
	var size string
	var scrapedSource string // TODO: get this?
	for _, customFields := range r.Props.PageProps.Company.CustomFields {
		switch customFields.Name {
		case ContactEmail:
			email = fmt.Sprintf("%v", customFields.Value)
			break
		case OrgSize:
			size = fmt.Sprintf("%v", customFields.Value)
			break
		}
	}
	tags := make([]string, len(r.Props.PageProps.Company.Sectors))
	for i, v := range r.Props.PageProps.Company.Sectors {
		tags[i] = convertChars(v.Tags)
	}
	return &model.Organization{
		SiteID:       r.Props.PageProps.Company.ID,
		Name:         r.Props.PageProps.Company.Name,
		Description:  convertChars(stripHTML(r.Props.PageProps.Company.Description)),
		Tags:         tags,
		Size:         size,
		Location:     r.Props.PageProps.Company.Location,
		URL:          r.Props.PageProps.Company.Website,
		ReferenceURL: scrapedSource,
		Logo:         r.Props.PageProps.Company.Logo,
		PostingCount: r.Props.PageProps.Company.ActiveCount,
		Outreach: model.Contact{
			Name:  r.Props.PageProps.Company.ContactName,
			Email: email,
		},
		Active: r.Props.PageProps.Company.Active,
	}
}

func OrganizationsToCSV(orgs []*model.Organization) [][]string {
	csv := make([][]string, len(orgs)+1)

	csv[0] = []string{"ID", "Name", "Size", "Link", "Tags", "Location", "Description"}
	for i, v := range orgs {
		csvTags := strings.Join(v.Tags, ", ")
		csv[i+1] = []string{strconv.Itoa(v.SiteID), v.Name, v.Size, v.URL, csvTags, v.Location, v.Description}
	}
	return csv
}

func CSVToOrganizations(org []string) *model.Organization {
	id, _ := strconv.Atoi(org[0])
	return &model.Organization{
		SiteID:      id,
		Name:        org[1],
		Description: org[6],
		Tags:        strings.Split(org[4], ", "),
		Size:        org[2],
		Location:    org[5],
		URL:         org[3],
	}
}

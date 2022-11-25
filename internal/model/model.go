package model

type Contact struct {
	Name  string
	Email string
}

type Organization struct {
	SiteID       int // ID supplied from site (not striclty unique across multiple sites)
	Name         string
	Description  string
	Tags         []string // Tags on the org. e.g. Sectors: ["Energy", "Food & Agriculture"]
	Size         string   // Range e.g. "101-250"
	Location     string   // Location of the organization
	URL          string   // Link to the org.'s website
	ReferenceURL string   // Link to where the struct info was scraped from
	Logo         string   // Link to logo image
	PostingCount int      // How many job postings are listed for the org.
	Outreach     Contact  // Listed contact information
	Active       bool     // If the listing/company/org. is active
}

type JSONResponse struct {
	Props struct {
		PageProps struct {
			Company struct {
				ID          int    `json:"id"`
				Name        string `json:"company_name"`
				Description string `json:"company_description"`
				Location    string `json:"location"`
				Logo        string `json:"logo"`
				Website     string `json:"website"`
				ContactName string `json:"full_name"`
				Active      bool   `json:"active"`
				ActiveCount int    `json:"active_jobs_count"`
				Sectors     []struct {
					Tags string `json:"name_value"`
				} `json:"sectors"`
				CustomFields []struct {
					Name  string      `json:"name"`
					Value interface{} `json:"value"`
				} `json:"custom_fields"`
			} `json:"company"`
		} `json:"pageProps"`
	} `json:"props"`
}

package orgscraper

import (
	"green-scraper/internal/model"
)

// Service TODO: make an actual service if required
type Service interface {
	CreateCSV(orgs []*model.Organization) error
}

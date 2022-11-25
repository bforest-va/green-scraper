package orgscraper

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"green-scraper/internal/model"
	"green-scraper/internal/service/adapters"
)

func CreateCSV(filename string, orgs []*model.Organization) error {
	fmt.Println("Updating CSV...")
	csvFile, err := os.Create(filename)
	if err != nil {
		return err
	}

	csvContent := adapters.OrganizationsToCSV(orgs)

	csvWriter := csv.NewWriter(csvFile)
	for _, c := range csvContent {
		_ = csvWriter.Write(c)
	}
	csvWriter.Flush()
	_ = csvFile.Close()
	fmt.Printf("CSV updated with %d records.", len(csvContent)-1)
	return nil
}

func LoadCSV(filename string) ([]*model.Organization, error) {
	csvFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	r := csv.NewReader(csvFile)
	company, err := r.Read()
	var orgs []*model.Organization
	if err == nil {
		for {
			company, err = r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				return nil, err
			}
			orgs = append(orgs, adapters.CSVToOrganizations(company))
		}
		fmt.Println("CSV Loaded")
	}
	return orgs, nil
}

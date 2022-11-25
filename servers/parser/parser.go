package main

import (
	"log"
	"os"
	"regexp"

	"green-scraper/internal/model"
	orgscraper "green-scraper/internal/service"
)

func main() {
	organizations, err := orgscraper.LoadCSV("./output/companies.csv")
	if err != nil {
		log.Println("Failed to read CSV with:", err.Error())
		os.Exit(0)
	}

	re := regexp.MustCompile(`consulting|brand|public relations|ads |PR |advertising`)
	var filtered []*model.Organization
	for _, org := range organizations {
		if re.MatchString(org.Description) {
			filtered = append(filtered, org)
		}
	}

	if len(filtered) > 0 {
		if err := orgscraper.CreateCSV("./output/filtered_companies.csv", filtered); err != nil {
			log.Println("Failed to create a file with:", err.Error())
			os.Exit(0)
		}
	}
}

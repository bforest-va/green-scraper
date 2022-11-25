package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"green-scraper/internal/model"
	orgscraper "green-scraper/internal/service"
	"green-scraper/internal/service/adapters"

	"github.com/gocolly/colly/v2"
)

const (
	TargetDomain = "climatebase.org"
	XMLURL       = "https://climatebase.org/organizations-sitemap.xml"
)

func main() {
	var organizations []*model.Organization

	// Instantiate default collector
	c := colly.NewCollector(colly.AllowedDomains(TargetDomain), colly.Async(true), colly.CacheDir("../cache"))

	// Create another collector to scrape company details
	companyCollector := c.Clone()

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	companyCollector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting Company", r.URL.String())
	})

	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "nError:", err)
	})

	companyCollector.OnHTML("#__NEXT_DATA__", func(e *colly.HTMLElement) {
		var scrape *model.JSONResponse
		if err := json.Unmarshal([]byte(e.Text), &scrape); err != nil {
			log.Println("Failed to Unmarshal with:", err.Error())
			os.Exit(0)
		}
		organizations = append(organizations, adapters.ToOrganization(scrape))
	})

	c.OnXML("//urlset/url/loc", func(e *colly.XMLElement) {
		companyCollector.Visit(e.Text)
	})

	if err := c.Visit(XMLURL); err != nil {
		log.Println("Failed to visit", XMLURL)
	}

	if len(organizations) > 0 {
		if err := orgscraper.CreateCSV("./output/companies.csv", organizations); err != nil {
			log.Println("Failed to create a file with:", err.Error())
			os.Exit(0)
		}
	}
}

//// TODO: delete this
//// div[data-insights-index] > a
//c.OnHTML("a[href]", func(e *colly.HTMLElement) {
//	//fmt.Println(e)
//	link := e.Attr("href")
//	if strings.Contains(link, "/company/") {
//		e.Request.Visit(link)
//		companyCollector.Visit("https://climatebase.org/company/1130954/form-energy?source=climatebase_orgs")
//		companyCollector.Visit("https://climatebase.org/company/1131202/electric-hydrogen?source=climatebase_orgs")
//	}
//})

//// TODO: have this selector actually hit...
//c.OnHTML(`div[data-insights-index] > a`, func(e *colly.HTMLElement) {
//	companyURL := e.Attr("href")
//	fmt.Println(companyURL)
//	if strings.Contains(companyURL, "/company/") {
//		companyCollector.Visit(companyURL)
//	}
//})

//if err := c.Visit("https://8psnffqtxq-dsn.algolia.net/1/indexes/*/queries?x-algolia-agent=Algolia%20for%20JavaScript%20(4.12.2)%3B%20Browser%3B%20JS%20Helper%20(3.7.0)%3B%20react%20(17.0.2)%3B%20react-instantsearch%20(6.19.0)"); err != nil {
//	fmt.Println(err)
//}

//for i := 0; i < 1; i++ {
//	err := c.Visit(VisitingURL + QueryParams + strconv.Itoa(i))
//	if err != nil {
//		fmt.Printf("Failed to visit %s with error: %s\n", VisitingURL, err.Error())
//	}
//}

//	POST /1/indexes/*/queries?x-algolia-agent=Algolia%20for%20JavaScript%20(4.12.2)%3B%20Browser%3B%20JS%20Helper%20(3.7.0)%3B%20react%20(17.0.2)%3B%20react-instantsearch%20(6.19.0) HTTP/1.1
//	Accept: */*
//		Accept-Encoding: gzip, deflate, br
//	Accept-Language: en-US,en;q=0.9
//  Connection: keep-alive
//	Content-Length: 192
//  Host: 8psnffqtxq-dsn.algolia.net
//  Origin: https://climatebase.org
//  Referer: https://climatebase.org/
//	Sec-Fetch-Dest: empty
//	Sec-Fetch-Mode: cors
//	Sec-Fetch-Site: cross-site
//	User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36
//	content-type: application/x-www-form-urlencoded
//	sec-ch-ua: "Google Chrome";v="107", "Chromium";v="107", "Not=A?Brand";v="24"
//	sec-ch-ua-mobile: ?0
//	sec-ch-ua-platform: "macOS"
//	x-algolia-api-key: d2ebe27d3cc3d35fea04da7b1b0718a8
//	x-algolia-application-id: 8PSNFFQTXQ
//alg := colly.NewCollector()
//alg.Post("https://8psnffqtxq-dsn.algolia.net/1/indexes/*/queries?x-algolia-agent=Algolia%20for%20JavaScript%20(4.12.2)%3B%20Browser%3B%20JS%20Helper%20(3.7.0)%3B%20react%20(17.0.2)%3B%20react-instantsearch%20(6.19.0)", map[string]string{"x-algolia-api-key": "d2ebe27d3cc3d35fea04da7b1b0718a8", "x-algolia-application-id": "8PSNFFQTXQ"})
//alg.OnResponse(func(r *colly.Response) {
//	log.Println("response received", r.StatusCode)
//})
//alg.OnRequest(func(r *colly.Request) {
//	fmt.Println("Visiting", r.URL.String())
//})
//alg.Visit("https://8psnffqtxq-dsn.algolia.net/1/indexes/*/queries?x-algolia-agent=Algolia%20for%20JavaScript%20(4.12.2)%3B%20Browser%3B%20JS%20Helper%20(3.7.0)%3B%20react%20(17.0.2)%3B%20react-instantsearch%20(6.19.0)")
//alg.OnHTML("", func(e *colly.HTMLElement) {
//	fmt.Println(e)
//})

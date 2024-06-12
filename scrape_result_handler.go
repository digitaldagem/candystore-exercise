package main

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"io"
	"log"
	"net/http"
	"sort"
	"strconv"
)

// scrapeResultHandler handles the /top_customer_favorite_snacks endpoint
func scrapeResultHandler(w http.ResponseWriter, r *http.Request) {
	url := "https://candystore.zimpler.net/#candystore-customers"
	rowMap := make(map[string]CustomerData)

	doc := fetchHTMLDocument(url)
	if doc == nil {
		http.Error(w, "failed to fetch and parse HTML document", http.StatusInternalServerError)
		return
	}

	extractData(doc, rowMap)
	results := processResults(rowMap)

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(results)
	if err != nil {
		log.Println("error encoding results json", err)
	}
}

// fetchHTMLDocument fetches the HTML document from the given URL and parses it
func fetchHTMLDocument(url string) *goquery.Document {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("error closing reader", err)
		}
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	return doc
}

// extractData extracts the relevant data from the HTML document and updates the map
func extractData(doc *goquery.Document, rowMap map[string]CustomerData) {
	doc.Find("#top\\.customers tbody tr").Each(func(i int, s *goquery.Selection) {
		row := make([]string, 0, 3)
		s.Find("td").Each(func(i int, s *goquery.Selection) {
			row = append(row, s.Text())
		})
		processCustomerData(row, rowMap)
	})
}

// processCustomerData processes a single row of customer data and updates the map
func processCustomerData(customerData []string, customerDataMap map[string]CustomerData) {
	key := customerData[0] + customerData[1]
	num, err := strconv.Atoi(customerData[2])
	if err != nil {
		log.Fatal(err)
	}

	if existing, exists := customerDataMap[key]; exists {
		customerDataMap[key] = CustomerData{
			Name:  customerData[0],
			Snack: customerData[1],
			Total: num + existing.Total,
		}
	} else {
		customerDataMap[key] = CustomerData{
			Name:  customerData[0],
			Snack: customerData[1],
			Total: num,
		}
	}
}

// processResults processes the final results and returns a slice of CustomerData
func processResults(customerDataMap map[string]CustomerData) []CustomerData {
	resultsMap := make(map[string]CustomerData)
	for _, value := range customerDataMap {
		if existing, exists := resultsMap[value.Name]; exists {
			if value.Total > existing.Total {
				resultsMap[value.Name] = value
			}
		} else {
			resultsMap[value.Name] = value
		}
	}

	results := make([]CustomerData, 0, len(resultsMap))
	for _, value := range resultsMap {
		results = append(results, value)
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].Total > results[j].Total
	})
	return results
}

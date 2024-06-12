package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/top_customer_favorite_snacks", scrapeResultHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

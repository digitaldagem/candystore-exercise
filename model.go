package main

// CustomerData represents a customer data row in the table
type CustomerData struct {
	Name  string `json:"name"`
	Snack string `json:"favoriteSnack"`
	Total int    `json:"totalSnacks"`
}

// Package models contains the data structures for the application.
package models

// Cart represents a shopping cart with products and total price.
type Cart struct {
	ID         int       `json:"id"`
	Products   []Product `json:"products"`
	TotalPrice float64   `json:"total_price"`
	Status     string    `json:"status"`
}

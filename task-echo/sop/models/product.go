// Package models contains the data structures for the application.
package models

// Product represents an item available for purchase.
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

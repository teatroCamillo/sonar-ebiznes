// Package models contains the data structures for the application.
package models

// Payment represents a payment transaction.
type Payment struct {
	ID     int     `json:"id"`
	Amount float64 `json:"amount"`
	Method string  `json:"method"`
	CartID int     `json:"cart_id"`
}

package models

type Cart struct {
	ID         int       `json:"id"`
	Products   []Product `json:"products"`
	TotalPrice float64   `json:"total_price"`
	Status     string    `json:"status"`
}
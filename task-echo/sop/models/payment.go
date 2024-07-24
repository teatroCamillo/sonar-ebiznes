package models

type Payment struct {
	ID     int     `json:"id"`
	Amount float64 `json:"amount"`
	Method string  `json:"method"`
	CartID int     `json:"cart_id"`
}
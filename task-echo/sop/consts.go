// Package sop is a source for controllers, models and consts
package sop

// Endpoints for the API
const (
	Localhost		   = "http://localhost:3000"
    ProductsEndpoint   = "/products"
    CartsEndpoint      = "/carts"
    PaymentsEndpoint   = "/payments"
    ProductByID        = "/products/:id"
    CartByID           = "/carts/:id"
    CartProducts       = "/carts/:id/products/:productId"
    PaymentByCartID    = "/payments/:cartID"
    PaymentByID        = "/payments/:id"
)
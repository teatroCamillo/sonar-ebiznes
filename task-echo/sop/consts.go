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

	InvalidCartID	   = "invalid cart ID"
	CartNotFound	   = "cart not found"

	CreateCartErrorv   = "CreateCart error: %v"
	ProductContains    = "Products contains:"

	InvalidPaymentID   = "invalid payment ID"
	PaymentNotFound	   = "payment not found"

	CreatePaymentErrorv   = "CreatePayment error: %v"
)
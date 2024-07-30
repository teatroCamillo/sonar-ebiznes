// Package main is the entry point for the application.
package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"task/sop/controllers"
	"task/sop"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{sop.Localhost},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	pc := controllers.NewProductController()
	cartC := controllers.NewCartController(pc)
	paymentC := controllers.NewPaymentController(cartC)

	e.POST(sop.ProductsEndpoint, pc.CreateProduct)
	e.GET(sop.ProductsEndpoint, pc.GetProducts)
	e.PUT(sop.ProductByID, pc.UpdateProduct)
	e.DELETE(sop.ProductByID, pc.DeleteProduct)

	e.POST(sop.CartsEndpoint, cartC.CreateCart)
	e.GET(sop.CartByID, cartC.GetCart)
	e.GET(sop.CartsEndpoint, cartC.GetCarts)
	e.POST(sop.CartProducts, cartC.AddProductToCart)
	e.DELETE(sop.CartByID, cartC.DeleteCart)
	e.DELETE(sop.CartProducts, cartC.RemoveProductFromCart)

	e.POST(sop.PaymentByCartID, paymentC.CreatePayment)
	e.GET(sop.PaymentsEndpoint, paymentC.GetPayments)
	e.GET(sop.PaymentByID, paymentC.GetPayment)
	e.PUT(sop.PaymentByID, paymentC.UpdatePayment)
	e.DELETE(sop.PaymentByID, paymentC.DeletePayment)

	e.Logger.Fatal(e.Start(":8080"))
}

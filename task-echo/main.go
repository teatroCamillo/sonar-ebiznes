package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"task/sop/controllers"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
	}))

	pc := controllers.NewProductController()
	cartC := controllers.NewCartController(pc)
	paymentC := controllers.NewPaymentController(cartC)

	e.POST("/products", pc.CreateProduct)
	e.GET("/products", pc.GetProducts)
	e.PUT("/products/:id", pc.UpdateProduct)
	e.DELETE("/products/:id", pc.DeleteProduct)

	e.POST("/carts", cartC.CreateCart)
	e.GET("/carts/:id", cartC.GetCart)
	e.GET("/carts", cartC.GetCarts)
	e.POST("/carts/:id/products/:productId", cartC.AddProductToCart)
	e.DELETE("/carts/:id", cartC.DeleteCart)
	e.DELETE("/carts/:id/products/:productId", cartC.RemoveProductFromCart)

	e.POST("/payments/:cartID", paymentC.CreatePayment)
	e.GET("/payments", paymentC.GetPayments)
	e.GET("/payments/:id", paymentC.GetPayment)
	e.PUT("/payments/:id", paymentC.UpdatePayment)
	e.DELETE("/payments/:id", paymentC.DeletePayment)

	e.Logger.Fatal(e.Start(":8080"))
}

package controllers

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"task/sop/models"
	"errors"
	"strconv"
)

type PaymentController struct {
	payments []models.Payment
	nextID   int
	cc       *CartController
}

func NewPaymentController(cc *CartController) *PaymentController {
	return &PaymentController{
		payments: []models.Payment{},
		nextID:   1,
		cc:       cc,
	}
}

func (pc *PaymentController) CreatePayment(c echo.Context) error {
	cartID, err := strconv.Atoi(c.Param("cartID"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.New("invalid cart ID"))
	}

	var cart *models.Cart
	for i := range pc.cc.carts {
		if pc.cc.carts[i].ID == cartID {
			cart = &pc.cc.carts[i]
			break
		}
	}

	if cart == nil {
		return c.JSON(http.StatusNotFound, errors.New("cart not found"))
	}

	if cart.Status == "paid" {
		return c.JSON(http.StatusForbidden, errors.New("cart is already paid"))
	}

	payment := models.Payment{
		ID:     pc.nextID,
		Amount: cart.TotalPrice,
		Method: "card", // Można zmienić na dynamiczne pobieranie metody płatności z żądania
		CartID: cart.ID,
	}

	pc.nextID++
	pc.payments = append(pc.payments, payment)
	cart.Status = "paid"

	return c.JSON(http.StatusCreated, payment)
}
// func (pc *PaymentController) CreatePayment(c echo.Context) error {
// 	cartID, err := strconv.Atoi(c.Param("cart_id"))
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, errors.New("invalid cart ID"))
// 	}

// 	var cart *models.Cart
// 	for i := range pc.cc.carts {
// 		if pc.cc.carts[i].ID == cartID {
// 			cart = &pc.cc.carts[i]
// 			break
// 		}
// 	}

// 	if cart == nil {
// 		return c.JSON(http.StatusNotFound, errors.New("cart not found"))
// 	}

// 	if cart.Status == "paid" {
// 		return c.JSON(http.StatusForbidden, errors.New("cart is already paid"))
// 	}

// 	payment := models.Payment{
// 		ID:     pc.nextID,
// 		Amount: cart.TotalPrice,
// 		Method: "card", // Można zmienić na dynamiczne pobieranie metody płatności z żądania
// 		CartID: cart.ID,
// 	}

// 	pc.nextID++
// 	pc.payments = append(pc.payments, payment)
// 	cart.Status = "paid"

// 	return c.JSON(http.StatusCreated, payment)
// }

func (pc *PaymentController) GetPayments(c echo.Context) error {
	return c.JSON(http.StatusOK, pc.payments)
}

func (pc *PaymentController) GetPayment(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 0 {
		return c.JSON(http.StatusBadRequest, errors.New("invalid payment ID"))
	}
	for _, payment := range pc.payments {
		if payment.ID == id {
			return c.JSON(http.StatusOK, payment)
		}
	}
	return c.JSON(http.StatusNotFound, errors.New("payment not found"))
}

func (pc *PaymentController) UpdatePayment(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 0 {
		return c.JSON(http.StatusBadRequest, errors.New("invalid payment ID"))
	}

	var updatedPayment models.Payment
	if err := c.Bind(&updatedPayment); err != nil {
		return err
	}

	for i := range pc.payments {
		if pc.payments[i].ID == id {
			pc.payments[i] = updatedPayment
			return c.JSON(http.StatusOK, updatedPayment)
		}
	}

	return c.JSON(http.StatusNotFound, errors.New("payment not found"))
}

func (pc *PaymentController) DeletePayment(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 0 {
		return c.JSON(http.StatusBadRequest, errors.New("invalid payment ID"))
	}

	for i, payment := range pc.payments {
		if payment.ID == id {
			pc.payments = append(pc.payments[:i], pc.payments[i+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}

	return c.JSON(http.StatusNotFound, errors.New("payment not found"))
}

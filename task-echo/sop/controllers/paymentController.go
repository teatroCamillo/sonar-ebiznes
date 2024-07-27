// Package controllers contains the controllers for handling HTTP requests.
package controllers

import (
	"net/http"
	"strconv"
	"errors"
	"task/sop/models"

	"github.com/labstack/echo/v4"
)

// PaymentController handles operations on payments.
type PaymentController struct {
	payments []models.Payment
	nextID   int
	cc       *CartController
}

// NewPaymentController creates a new PaymentController.
func NewPaymentController(cc *CartController) *PaymentController {
	return &PaymentController{
		payments: []models.Payment{},
		nextID:   1,
		cc:       cc,
	}
}

// CreatePayment handles the creation of a new payment.
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
		Method: "card",
		CartID: cart.ID,
	}

	pc.nextID++
	pc.payments = append(pc.payments, payment)
	cart.Status = "paid"

	return c.JSON(http.StatusCreated, payment)
}

// GetPayments handles retrieving all payments.
func (pc *PaymentController) GetPayments(c echo.Context) error {
	return c.JSON(http.StatusOK, pc.payments)
}

// GetPayment handles retrieving a single payment by ID.
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

// UpdatePayment handles updating an existing payment.
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

// DeletePayment handles deleting a payment by ID.
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

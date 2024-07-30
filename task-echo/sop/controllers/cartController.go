// Package controllers contains the controllers for handling HTTP requests.
package controllers

import (
	"net/http"
	"strconv"
	"errors"
	"task/sop/models"
	"github.com/labstack/echo/v4"
	"task/sop"
)

// CartController handles operations on carts.
type CartController struct {
	carts  []models.Cart
	nextID int
	pc     *ProductController
}

// NewCartController creates a new CartController.
func NewCartController(pc *ProductController) *CartController {
	return &CartController{
		carts: []models.Cart{
			{ID: 0, Products: []models.Product{}, TotalPrice: 0, Status: "new"},
		},
		nextID: 1,
		pc:     pc,
	}
}

// CreateCart handles the creation of a new cart.
func (cc *CartController) CreateCart(c echo.Context) error {
	var cart models.Cart
	if err := c.Bind(&cart); err != nil {
		return err
	}
	cart.ID = cc.nextID
	cart.Status = "new"
	cc.nextID++
	cc.carts = append(cc.carts, cart)
	return c.JSON(http.StatusCreated, cart)
}

// GetCarts handles retrieving all carts.
func (cc *CartController) GetCarts(c echo.Context) error {
	return c.JSON(http.StatusOK, cc.carts)
}

// GetCart handles retrieving a single cart by ID.
func (cc *CartController) GetCart(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id < 0 {
		return c.JSON(http.StatusBadRequest, errors.New(sop.InvalidCartID))
	}
	for _, cart := range cc.carts {
		if cart.ID == id {
			return c.JSON(http.StatusOK, cart)
		}
	}
	return c.JSON(http.StatusNotFound, errors.New(sop.CartNotFound))
}

// AddProductToCart handles adding a product to a cart.
func (cc *CartController) AddProductToCart(c echo.Context) error {
	cartID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.New(sop.InvalidCartID))
	}

	productID, err := strconv.Atoi(c.Param("productId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.New("invalid product ID"))
	}

	var cart *models.Cart
	for i := range cc.carts {
		if cc.carts[i].ID == cartID {
			cart = &cc.carts[i]
			break
		}
	}

	if cart == nil || cart.Status == "paid" {
		cart = &models.Cart{
			ID:         cc.nextID,
			Products:   []models.Product{},
			TotalPrice: 0.0,
			Status:     "new",
		}
		cc.nextID++
		cc.carts = append(cc.carts, *cart)
	}

	var product *models.Product
	for _, p := range cc.pc.products {
		if p.ID == productID {
			product = &p
			break
		}
	}

	if product == nil {
		return c.JSON(http.StatusNotFound, errors.New("product not found"))
	}

	cart.Products = append(cart.Products, *product)
	cart.TotalPrice = calculateTotalPrice(cart.Products)

	return c.JSON(http.StatusOK, cart)
}

// RemoveProductFromCart handles removing a product from a cart.
func (cc *CartController) RemoveProductFromCart(c echo.Context) error {
	cartID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.New(sop.InvalidCartID))
	}

	productID, err := strconv.Atoi(c.Param("productId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.New("invalid product ID"))
	}

	var cart *models.Cart
	for i := range cc.carts {
		if cc.carts[i].ID == cartID {
			cart = &cc.carts[i]
			break
		}
	}

	if cart == nil {
		return c.JSON(http.StatusNotFound, errors.New(sop.CartNotFound))
	}

	productFound := false
	for i, p := range cart.Products {
		if p.ID == productID {
			cart.Products = append(cart.Products[:i], cart.Products[i+1:]...)
			cart.TotalPrice = calculateTotalPrice(cart.Products)
			productFound = true
			break
		}
	}

	if !productFound {
		return c.JSON(http.StatusNotFound, errors.New("product not found in cart"))
	}

	return c.NoContent(http.StatusNoContent)
}

// DeleteCart handles deleting a cart by ID.
func (cc *CartController) DeleteCart(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.New(sop.InvalidCartID))
	}

	for i, cart := range cc.carts {
		if cart.ID == id {
			cc.carts = append(cc.carts[:i], cc.carts[i+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}

	return c.JSON(http.StatusNotFound, errors.New(sop.CartNotFound))
}

func calculateTotalPrice(products []models.Product) float64 {
	total := 0.0
	for _, product := range products {
		total += product.Price
	}
	return total
}

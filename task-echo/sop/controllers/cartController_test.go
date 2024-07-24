package controllers

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateCart(t *testing.T) {
	e := echo.New()
	pc := NewProductController()
	cc := NewCartController(pc)

	cartJSON := `{"products":[]}`
	req := httptest.NewRequest(http.MethodPost, "/carts", strings.NewReader(cartJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, cc.CreateCart(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}

func TestGetCarts(t *testing.T) {
	e := echo.New()
	pc := NewProductController()
	cc := NewCartController(pc)

	cc.CreateCart(e.NewContext(
		httptest.NewRequest(http.MethodPost, "/carts", strings.NewReader(`{"products":[]}`)),
		httptest.NewRecorder(),
	))

	req := httptest.NewRequest(http.MethodGet, "/carts", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, cc.GetCarts(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetCart(t *testing.T) {
	e := echo.New()
	pc := NewProductController()
	cc := NewCartController(pc)

	cc.CreateCart(e.NewContext(
		httptest.NewRequest(http.MethodPost, "/carts", strings.NewReader(`{"products":[]}`)),
		httptest.NewRecorder(),
	))

	req := httptest.NewRequest(http.MethodGet, "/carts/0", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("0")

	if assert.NoError(t, cc.GetCart(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestAddProductToCart(t *testing.T) {
	e := echo.New()
	pc := NewProductController()
	cc := NewCartController(pc)

	log.Println("Products contains:")
	log.Println(pc.products)
	pc.CreateProduct(e.NewContext(
		httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(`{"name":"Smartphone","price":699.99}`)),
		httptest.NewRecorder(),
	))
	log.Println("Products contains:")
	log.Println(pc.products)
	
	cc.CreateCart(e.NewContext(
		httptest.NewRequest(http.MethodPost, "/carts", strings.NewReader(`{"products":[]}`)),
		httptest.NewRecorder(),
	))

	req := httptest.NewRequest(http.MethodPost, "/carts/0/products/0", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id", "productId")
	c.SetParamValues("0", "0")

	if assert.NoError(t, cc.AddProductToCart(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "Smartphone")
	}
}

func TestRemoveProductFromCart(t *testing.T) {
	e := echo.New()
	pc := NewProductController()
	cc := NewCartController(pc)

	log.Println("Products contains:")
	log.Println(pc.products)

	// Dodajemy nowy produkt (ID 10) do listy produktów
	pc.CreateProduct(e.NewContext(
		httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(`{"name":"New Product","price":99.99}`)),
		httptest.NewRecorder(),
	))
	log.Println("Products contains 2:")
	log.Println(pc.products)

	cc.CreateCart(e.NewContext(
		httptest.NewRequest(http.MethodPost, "/carts", strings.NewReader(`{"products":[]}`)),
		httptest.NewRecorder(),
	))

	log.Println("Cart BEFORE population contains:")
	log.Println(cc.carts)

	cc.AddProductToCart(e.NewContext(
		httptest.NewRequest(http.MethodPost, "/carts/0/products/10", nil),
		httptest.NewRecorder(),
	))

	log.Println("Cart after population contains:")
	log.Println(cc.carts)

	log.Println("Preparing to remove product from cart")
	req := httptest.NewRequest(http.MethodDelete, "/carts/0/products/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id", "productId")
	c.SetParamValues("0", "1")

	if assert.NoError(t, cc.RemoveProductFromCart(c)) {
		assert.Equal(t, http.StatusNoContent, rec.Code)
	} else {
		log.Printf("Failed to remove product from cart: %v\n", rec.Body.String())
	}
	log.Println("END rec: ", rec)
}

func TestDeleteCart(t *testing.T) {
	e := echo.New()
	pc := NewProductController()
	cc := NewCartController(pc)

	cc.CreateCart(e.NewContext(
		httptest.NewRequest(http.MethodPost, "/carts", strings.NewReader(`{"products":[]}`)),
		httptest.NewRecorder(),
	))

	req := httptest.NewRequest(http.MethodDelete, "/carts/0", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("0")

	if assert.NoError(t, cc.DeleteCart(c)) {
		assert.Equal(t, http.StatusNoContent, rec.Code)
	}
}

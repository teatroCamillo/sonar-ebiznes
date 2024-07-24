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

func TestCreateProduct(t *testing.T) {
	e := echo.New()
	pc := NewProductController()

	log.Println("Products contains:")
	log.Println(pc.products)
	productJSON := `{"name":"Termostat","price":699.99}`
	req := httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(productJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	log.Println("Products contains:")
	log.Println(pc.products)

	if assert.NoError(t, pc.CreateProduct(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Contains(t, rec.Body.String(), "Termostat")
	}
}

func TestGetProducts(t *testing.T) {
	e := echo.New()
	pc := NewProductController()

	req := httptest.NewRequest(http.MethodGet, "/products", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, pc.GetProducts(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "Smartphone")
	}
}

func TestUpdateProduct(t *testing.T) {
	e := echo.New()
	pc := NewProductController()

	productJSON := `{"name":"Updated Smartphone","price":799.99}`
	req := httptest.NewRequest(http.MethodPut, "/products/0", strings.NewReader(productJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("0")

	if assert.NoError(t, pc.UpdateProduct(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "Updated Smartphone")
	}
}

func TestDeleteProduct(t *testing.T) {
	e := echo.New()
	pc := NewProductController()

	req := httptest.NewRequest(http.MethodDelete, "/products/0", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("0")

	if assert.NoError(t, pc.DeleteProduct(c)) {
		assert.Equal(t, http.StatusNoContent, rec.Code)
	}
}

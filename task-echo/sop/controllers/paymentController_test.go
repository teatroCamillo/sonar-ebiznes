package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreatePayment(t *testing.T) {
	e := echo.New()
	pc := NewProductController()
	cc := NewCartController(pc)
	paymentC := NewPaymentController(cc)

	// Tworzenie koszyka
	cc.CreateCart(e.NewContext(
		httptest.NewRequest(http.MethodPost, "/carts", strings.NewReader(`{"products":[]}`)),
		httptest.NewRecorder(),
	))

	req := httptest.NewRequest(http.MethodPost, "/payments/0", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("cart_id")
	c.SetParamValues("0")

	if assert.NoError(t, paymentC.CreatePayment(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Contains(t, rec.Body.String(), `"cart_id":0`)
		assert.Contains(t, rec.Body.String(), `"amount":0`)
	}

	// Próba ponownej płatności
	req = httptest.NewRequest(http.MethodPost, "/payments/0", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)
	c.SetParamNames("cart_id")
	c.SetParamValues("0")

	if assert.Error(t, paymentC.CreatePayment(c)) {
		assert.Equal(t, http.StatusForbidden, rec.Code)
	}
}

func TestGetPayments(t *testing.T) {
	e := echo.New()
	pc := NewProductController()
	cc := NewCartController(pc)
	paymentC := NewPaymentController(cc)

	paymentC.CreatePayment(e.NewContext(
		httptest.NewRequest(http.MethodPost, "/payments/0", nil),
		httptest.NewRecorder(),
	))

	req := httptest.NewRequest(http.MethodGet, "/payments", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	if assert.NoError(t, paymentC.GetPayments(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestGetPayment(t *testing.T) {
	e := echo.New()
	pc := NewProductController()
	cc := NewCartController(pc)
	paymentC := NewPaymentController(cc)

	paymentC.CreatePayment(e.NewContext(
		httptest.NewRequest(http.MethodPost, "/payments/0", nil),
		httptest.NewRecorder(),
	))

	req := httptest.NewRequest(http.MethodGet, "/payments/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, paymentC.GetPayment(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), `"id":1`)
	}
}

func TestUpdatePayment(t *testing.T) {
	e := echo.New()
	pc := NewProductController()
	cc := NewCartController(pc)
	paymentC := NewPaymentController(cc)

	paymentC.CreatePayment(e.NewContext(
		httptest.NewRequest(http.MethodPost, "/payments/0", nil),
		httptest.NewRecorder(),
	))

	paymentJSON := `{"amount":150.0,"method":"PayPal","cart_id":0}`
	req := httptest.NewRequest(http.MethodPut, "/payments/1", strings.NewReader(paymentJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, paymentC.UpdatePayment(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "PayPal")
	}
}

func TestDeletePayment(t *testing.T) {
	e := echo.New()
	pc := NewProductController()
	cc := NewCartController(pc)
	paymentC := NewPaymentController(cc)

	paymentC.CreatePayment(e.NewContext(
		httptest.NewRequest(http.MethodPost, "/payments/0", nil),
		httptest.NewRecorder(),
	))

	req := httptest.NewRequest(http.MethodDelete, "/payments/1", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, paymentC.DeletePayment(c)) {
		assert.Equal(t, http.StatusNoContent, rec.Code)
	}
}

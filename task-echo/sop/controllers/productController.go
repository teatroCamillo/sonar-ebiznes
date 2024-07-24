package controllers

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"task/sop/models"
	"errors"
	"strconv"
)

type ProductController struct {
	products []models.Product
	nextID   int
}

func NewProductController() *ProductController {
	return &ProductController{
		products: []models.Product{
			{ID: 0, Name: "Smartphone", Price: 699.99},
			{ID: 1, Name: "Laptop", Price: 999.99},
			{ID: 2, Name: "Tablet", Price: 399.99},
			{ID: 3, Name: "Smartwatch", Price: 199.99},
			{ID: 4, Name: "Wireless Earbuds", Price: 149.99},
			{ID: 5, Name: "Bluetooth Speaker", Price: 89.99},
			{ID: 6, Name: "Gaming Console", Price: 499.99},
			{ID: 7, Name: "Digital Camera", Price: 549.99},
			{ID: 8, Name: "4K TV", Price: 799.99},
			{ID: 9, Name: "VR Headset", Price: 299.99},
		},
		nextID: 10,
	}
}

func (pc *ProductController) CreateProduct(c echo.Context) error {
	var product models.Product
	if err := c.Bind(&product); err != nil {
		return err
	}
	product.ID = pc.nextID
	pc.nextID++
	pc.products = append(pc.products, product)
	return c.JSON(http.StatusCreated, product)
}

func (pc *ProductController) UpdateProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.New("invalid product ID"))
	}

	var updatedProduct models.Product
	if err := c.Bind(&updatedProduct); err != nil {
		return err
	}

	for i, p := range pc.products {
		if p.ID == id {
			updatedProduct.ID = id
			pc.products[i] = updatedProduct
			return c.JSON(http.StatusOK, updatedProduct)
		}
	}

	return c.JSON(http.StatusNotFound, errors.New("product not found"))
}

func (pc *ProductController) GetProducts(c echo.Context) error {
	return c.JSON(http.StatusOK, pc.products)
}

func (pc *ProductController) DeleteProduct(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, errors.New("invalid product ID"))
	}

	for i, p := range pc.products {
		if p.ID == id {
			pc.products = append(pc.products[:i], pc.products[i+1:]...)
			return c.NoContent(http.StatusNoContent)
		}
	}

	return c.JSON(http.StatusNotFound, errors.New("product not found"))
}

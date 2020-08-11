package api

import (
	//"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	//"github.com/labstack/echo/middleware"
)

//Product is
type Product struct {
	Name   string `json:"name"`
	Colour string `json:"colour"`
	Price  int    `json:"price"`
	Id     int    `json:"id"`
}

//Products is
type Products []Product

var products = Products{Product{"Ball", "Black", 400, 15}, Product{"Bag", "Blue", 600, 16}}

//GetProducts is
func GetProducts(c echo.Context) error {
	return c.JSON(http.StatusOK, products)
}

//PostProduct is
func PostProduct(c echo.Context) error {
	product := Product{}
	err := c.Bind(&product)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}
	//product.Id = 50
	products = append(products, product)
	return c.JSON(http.StatusCreated, products)
}

//GetProduct is
func GetProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, product := range products {
		if product.Id == id {
			return c.JSON(http.StatusOK, product)
		}
	}
	return c.JSON(http.StatusBadRequest, nil)
}

//PutProduct is
func PutProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, _ := range products {
		if products[i].Id == id {
			products[i].Price = 1000
			return c.JSON(http.StatusOK, products)
		}
	}
	return c.JSON(http.StatusBadRequest, nil)
}

//DeleteProduct is
func DeleteProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, _ := range products {
		if products[i].Id == id {
			products = append(products[:i], products[i+1:]...)
			return c.JSON(http.StatusOK, products)
		}
	}
	return c.JSON(http.StatusBadRequest, nil)
}

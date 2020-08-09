package api

import (
	//"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	//"github.com/labstack/echo/middleware"
)

//Merchant is

type Merchant struct {
	Id int `json:"id"`
	//Nickname string `json:"nickname"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	Type string `json:"type"`
	//Online   bool   `json:"online"`
}

//Product is
type Product struct {
	Name   string `json:"name"`
	Colour string `json:"colour"`
	Price  int    `json:"price"`
	Id     int    `json:"id"`
	//Nickname string `json:"nickname"`
	//Online   bool   `json:"online"`
}

type Merchants []Merchant

type Products []Product

/*type Merchants []All.Merchant
type Products []all.Product*/
//var mp = all{{11, "Reem", 23, "Female"}, {"Bag", Blue, 600.13}}

//type Gamers []Gamer
var merchants = Merchants{Merchant{11, "Reem", 23, "Female"}, Merchant{12, "Silvana", 23, "Female"}}

var products = Products{Product{"Ball", "Black", 400, 15}, Product{"Bag", "Blue", 600, 16}}

//var players = Players{Player{generateId(), "DarthVader", true}, Player{generateId(), "Batman", true}}
//var players = Players{Player{11, "DarthVader", true}, Player{11, "Batman", true}}

//var Gamers = Gamers{Gamer{11, "DarthVader", true}, Gamer{11, "Batman", true}}

func getMerchants(c echo.Context) error {
	return c.JSON(http.StatusOK, merchants)
}

func postMerchant(c echo.Context) error {
	merchant := Merchant{}
	err := c.Bind(&merchant)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}
	//player.Id = 11
	merchants = append(merchants, merchant)
	return c.JSON(http.StatusCreated, merchants)
}

func getMerchant(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, merchant := range merchants {
		if merchant.Id == id {
			return c.JSON(http.StatusOK, merchant)
		}
	}
	return c.JSON(http.StatusBadRequest, nil)
}

func putMerchant(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, _ := range merchants {
		if merchants[i].Id == id {
			//merchants[i].Online = !merchants[i].Online
			return c.JSON(http.StatusOK, merchants)
		}
	}
	return c.JSON(http.StatusBadRequest, nil)
}

func deleteMerchant(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, _ := range merchants {
		if merchants[i].Id == id {
			merchants = append(merchants[:i], merchants[i+1:]...)
			return c.JSON(http.StatusOK, merchants)
		}
	}
	return c.JSON(http.StatusBadRequest, nil)
}
func getProducts(c echo.Context) error {
	return c.JSON(http.StatusOK, products)
}

func postProduct(c echo.Context) error {
	product := Product{}
	err := c.Bind(&product)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}

	products = append(products, product)
	return c.JSON(http.StatusCreated, products)
}

func getProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, product := range products {
		if product.Id == id {
			return c.JSON(http.StatusOK, product)
		}
	}
	return c.JSON(http.StatusBadRequest, nil)
}

func putProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, _ := range products {
		if products[i].Id == id {
			return c.JSON(http.StatusOK, products)
		}
	}
	return c.JSON(http.StatusBadRequest, nil)
}

func deleteProduct(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, _ := range products {
		if products[i].Id == id {
			products = append(products[:i], products[i+1:]...)
			return c.JSON(http.StatusOK, products)
		}
	}
	return c.JSON(http.StatusBadRequest, nil)
}

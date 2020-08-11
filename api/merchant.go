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
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
	Type string `json:"type"`
}

//Merchants is
type Merchants []Merchant

var merchants = Merchants{Merchant{11, "Reem", 23, "Female"}, Merchant{12, "Silvana", 23, "Female"}}

//GetMerchants is
func GetMerchants(c echo.Context) error {
	return c.JSON(http.StatusOK, merchants)
}

//PostMerchant is
func PostMerchant(c echo.Context) error {
	merchant := Merchant{}
	err := c.Bind(&merchant)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}
	//merchant.Id = 55
	merchants = append(merchants, merchant)
	return c.JSON(http.StatusCreated, merchants)
}

//GetMerchant is
func GetMerchant(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for _, merchant := range merchants {
		if merchant.Id == id {
			return c.JSON(http.StatusOK, merchant)
		}
	}
	return c.JSON(http.StatusBadRequest, nil)
}

//PutMerchant is
func PutMerchant(c echo.Context) error {

	id, _ := strconv.Atoi(c.Param("id"))
	for i, _ := range merchants {
		if merchants[i].Id == id {

			merchants[i].Type = "others"
			return c.JSON(http.StatusOK, merchants)
		}
	}
	return c.JSON(http.StatusBadRequest, nil)
}

//DeleteMerchant is
func DeleteMerchant(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	for i, _ := range merchants {
		if merchants[i].Id == id {
			merchants = append(merchants[:i], merchants[i+1:]...)
			return c.JSON(http.StatusOK, merchants)
		}
	}
	return c.JSON(http.StatusBadRequest, nil)
}

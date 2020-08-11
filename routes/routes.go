package routes

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/restapi/api"
)

var e *echo.Echo

func init() {
	e = echo.New()

	// Middleware
	e.Use(middleware.Logger())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

}

// Router list all the routes used by this app
func Router() *echo.Echo {

	//API groups
	v1 := e.Group("/api/v1")
	merchants := v1.Group("/merchants")
	products := v1.Group("/products")

	merchants.GET("/", api.GetMerchants)
	merchants.POST("/", api.PostMerchant)
	merchants.GET("/:id", api.GetMerchant)
	merchants.PUT("/:id", api.PutMerchant)
	merchants.DELETE("/:id", api.DeleteMerchant)

	products.GET("/", api.GetProducts)
	products.POST("/", api.PostProduct)
	products.GET("/:id", api.GetProduct)
	products.PUT("/:id", api.PutProduct)
	products.DELETE("/:id", api.DeleteProduct)

	return e
}

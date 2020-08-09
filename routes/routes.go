package routes

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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
	products := e.Group("/api/products")
	merchants := e.Group("api/merchants")

	merchants.GET("/", api.getMerchants)
	merchants.POST("/", api.postMerchant)
	merchants.GET("/:id", api.getMerchant)
	merchants.PUT("/:id", api.putMerchant)
	merchants.DELETE("/:id", deleteMerchant)

	products.GET("/", api.getProducts)
	cats.POST("/", api.postProduct)
	cats.GET("/:id", api.getProduct)
	cats.PUT("/:id", api.putProduct)
	cats.DELETE("/:id", api.deleteProduct)

	return e
}

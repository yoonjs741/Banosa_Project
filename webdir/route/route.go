package route

import (
	"Banosa_Project/webdir/v1/app/controller"

	"github.com/labstack/echo"
)

// Route :
func Route() *echo.Echo {
	e := echo.New()
	e.GET("/", controller.HelloworldCTRL)
	e.GET("/asson", controller.HelloAssonCTRL)
	e.GET("/donjon", controller.HelloDonjonCTRL)
	e.GET("/query", controller.HelloQueryCTRL)
	return e

}
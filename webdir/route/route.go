package route

import (
	"Banosa_Project/webdir/v1/app/controller"

	"github.com/labstack/echo"
)

// Route : Route controller to path
func Route() *echo.Echo {
	e := echo.New()
	e.GET("/", controller.HelloworldCTRL)
	e.GET("/asson", controller.HelloAssonCTRL)
	e.GET("/donjon", controller.HelloDonjonCTRL)
	e.GET("/query", controller.HelloQueryCTRL)
	e.GET("/param", controller.HelloParamCTRL)
	e.GET("/facto", controller.HelloQueryFacto)
	return e

}

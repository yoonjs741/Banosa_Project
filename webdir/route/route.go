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
	e.GET("/factojson", controller.HelloQueryFactoJSON)
	e.GET("/price", controller.HelloPricingJSON)
	e.GET("/name", controller.HelloYourNameCTRL)
	e.POST("/userinfo", controller.HelloUserInfoCTRL)
	e.GET("/strtest", controller.HelloStrcutCTRL)
	//	e.GET("/gql", controller.HelloGQLJSON)
	return e
}

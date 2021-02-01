package controller

import (
	"Banosa_Project/webdir/v1/app/factorial"
	"net/http"

	"github.com/labstack/echo"
)

// HelloworldCTRL : Return string as 200 to CTRL
func HelloworldCTRL(c echo.Context) error {
	//return c.String(http.StatusOK, "Hello World!")
	return c.HTML(200, "<h2> Welcome to Banosa </h2>")
}

func HelloAssonCTRL(c echo.Context) error {
	return c.String(http.StatusOK, "Hello Asson! Hm...")
}

func HelloDonjonCTRL(c echo.Context) error {
	return c.String(http.StatusOK, "Hello DonJon! What should I do next :<")
}

func HelloQueryCTRL(c echo.Context) error {
	return c.String(http.StatusOK, c.QueryParams().Encode())
}

func HelloParamCTRL(c echo.Context) error {
	return c.String(http.StatusOK, c.QueryString())
}

func HelloQueryFacto(c echo.Context) error {
	//id := c.QueryString()
	id := c.QueryParam("id") //id의 값만 보게 변경
	return c.JSON(200, factorial.GetFacto(id))

}

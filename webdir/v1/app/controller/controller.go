package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

// HelloworldCTRL : Return string as 200 to CTRL
func HelloworldCTRL(c echo.Context) error {
	//return c.String(http.StatusOK, "Hello World!")
	return c.HTML(200, "<h2> Welcome to Banosa </h2>")
}

func HelloAssonCTRL(c echo.Context) error {
	return c.String(http.StatusOK, "Hello Asson!")
}

func HelloDonjonCTRL(c echo.Context) error {
	return c.String(http.StatusOK, "Hello DonJon!")
}

func HelloQueryCTRL(c echo.Context) error {
	return c.String(http.StatusOK, c.QueryParams().Encode())
}

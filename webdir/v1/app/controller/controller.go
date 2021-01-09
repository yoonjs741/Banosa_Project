package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

// HelloworldCTRL :
func HelloworldCTRL(c echo.Context) error {
	//return c.String(http.StatusOK, "Hello World!")
	return c.HTML(200, "<h2> Fuck you </h2>")
}

func HelloAssonCTRL(c echo.Context) error {
	return c.String(http.StatusOK, "Hello Asson Fuck you!")
}

func HelloDonjonCTRL(c echo.Context) error {
	return c.String(http.StatusOK, "Hello DonJon Shit up!")
}

func HelloQueryCTRL(c echo.Context) error {
	return c.String(http.StatusOK, c.QueryParams().Encode())
}

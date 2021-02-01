package controller

import (
	"Banosa_Project/webdir/v1/app/factorial"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// HelloworldCTRL : Return string as 200 to CTRL
func HelloworldCTRL(c echo.Context) error {
	//return c.String(http.StatusOK, "Hello World!")
	return c.HTML(200, "<h2> Welcome to Banosa </h2>")
}

// HelloAssonCTRL : Return Asson string as http ok to CTRL
func HelloAssonCTRL(c echo.Context) error {
	return c.String(http.StatusOK, "Hello Asson! Hm...")
}

// HelloDonjonCTRL : Return Donjon string as http ok to CTRL
func HelloDonjonCTRL(c echo.Context) error {
	return c.String(http.StatusOK, "Hello DonJon! What should I do next :<")
}

// HelloQueryCTRL : Return Emdcoded QueryParams  as http ok to CTRL
func HelloQueryCTRL(c echo.Context) error {
	return c.String(http.StatusOK, c.QueryParams().Encode())
}

// HelloParamCTRL : Return QueryString as http ok to CTRL
func HelloParamCTRL(c echo.Context) error {
	return c.String(http.StatusOK, c.QueryString())
}

// HelloQueryFacto : Return QueryFacto string as http ok to CTRL
func HelloQueryFacto(c echo.Context) error {
	var StrconvErr error
	var IDint int

	id := c.QueryParam("id") //id의 값만 보게 변경

	IDint, StrconvErr = strconv.Atoi(id)

	if StrconvErr != nil {
		log.Println(StrconvErr)
		return c.String(http.StatusOK, "Please input only number.")
	}

	return c.String(http.StatusOK, factorial.GetFacto(IDint))
}

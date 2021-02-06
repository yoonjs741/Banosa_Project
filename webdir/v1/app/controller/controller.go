package controller

import (
	"Banosa_Project/webdir/v1/app/factorial"
	"Banosa_Project/webdir/v1/app/pricing"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// ResponseJSON : Reponse data sturct
type ResponseJSON struct {
	Input string
	Data  string `json:"Result"`
}

type ResponsePriceJSON struct {
	Value  string
	Coupon string
	Price  string
}

// HelloworldCTRL : Return string as 200 to CTRL
func HelloworldCTRL(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!")
	//return c.HTML(200, html/index.html)

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

	id := c.QueryParam("id") //id의 값만 보게 변경

	//strconv.Atoi idInt 변수에 입력
	idInt, strConverr := strconv.Atoi(id)

	// strConverr가 nil이면 factorial.GetFacto 응답하기로 변경 else 문 제거,
	if strConverr != nil {
		log.Println(strConverr)
		return c.String(http.StatusOK, "Please input only number.")
	}

	// GetFacto 입력값 int로 변경
	// strconv 대신 fmt.Sprintf 사용. -> 더 빠르다고 함
	// 다시 strconv 사용. -> fmt.Sprintf 는 타입 변환과 출력을 동시에 할때 빠르고, 여기처럼 함수에 넣어서 바로 사용할땐 strconv가 더 빠르다고 함
	result := strconv.Itoa(factorial.GetFacto(idInt))
	return c.String(http.StatusOK, result)
}

// HelloQueryFactoJSON : Return QueryFacto string as JSON to CTRLe
func HelloQueryFactoJSON(c echo.Context) error {

	id := c.QueryParam("id")
	idInt, err := strconv.Atoi(id)

	if err != nil {
		log.Println(err)
		return c.String(http.StatusOK, "Please input only number.")
	}

	factoResultJSON := new(ResponseJSON)

	factoResultJSON.Input = c.QueryParam("id")
	factoResultJSON.Data = fmt.Sprintf("%d", factorial.GetFacto(idInt)) //Same with test:= ResponseJSON{}

	return c.JSON(http.StatusOK, factoResultJSON)
}

func HelloPricingJSON(c echo.Context) error {
	value := c.QueryParam("value")
	discount := c.QueryParam("coupon")

	valueInt, err1 := strconv.Atoi(value)
	discountInt, err2 := strconv.Atoi(discount)

	if err1 != nil {
		log.Println(err1)
		return c.String(http.StatusOK, "Please check Value")
	}

	if err2 != nil {
		log.Println(err2)
		return c.String(http.StatusOK, "Please check coupon")
	}
	priceResultJSON := new(ResponsePriceJSON)

	priceResultJSON.Value = c.QueryParam("value")
	priceResultJSON.Coupon = c.QueryParam("coupon")
	priceResultJSON.Price = fmt.Sprintf("%d", pricing.GetPrice(valueInt, discountInt))

	return c.JSON(http.StatusOK, priceResultJSON)
}

//func HelloGQLJSON(c echo.Context) error {
//	return c.String(http.StatusOK, graphql.GetPeopleData())
//}

package controller

import (
	"Banosa_Project/webdir/v1/app/factorial"
	"Banosa_Project/webdir/v1/app/pricing"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/labstack/echo"
)

// ResponseJSON : Reponse data sturct
type ResponseJSON struct {
	Input string
	Data  string `json:"Result"`
}

// ResponsePriceJSON : ResponsePrice JSON struct
type ResponsePriceJSON struct {
	Value  string
	Coupon string
	Price  string
}

// User : Request, Response User struct
type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	CreateAt  time.Time
}

type Dict struct {
	Data datablock `json:"data"`
}

type datablock struct {
	data map[int]string
}

type Template struct {
	templates *template.Template
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

// HelloYourNameCTRL : Return id query param
func HelloYourNameCTRL(c echo.Context) error {
	name := c.QueryParam("id")
	if name == "" {
		name = "World"
	}
	return c.String(http.StatusOK, fmt.Sprintf("Hello %s!", name))
}

// HelloUserInfoCTRL : Input JSON str, Reponse JSON str
func HelloUserInfoCTRL(c echo.Context) error {
	user := new(User)
	err := json.NewDecoder(c.Request().Body).Decode(user)
	if err != nil {
		c.Response().Writer.WriteHeader(http.StatusBadRequest)
		return err
	}
	user.CreateAt = time.Now()

	// c.Response().Header().Set(echo.HeaderContentType, echo.MIMEApplicationJSONCharsetUTF8)
	// c.Response().WriteHeader(http.StatusOK)

	return c.JSON(http.StatusOK, user)
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

// HelloPricingJSON : Return price result by json
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
	priceResultJSON := ResponsePriceJSON{
		Value:  c.QueryParam("value"),
		Coupon: c.QueryParam("coupon"),
		Price:  fmt.Sprintf("%d", pricing.GetPrice(valueInt, discountInt)),
	}

	return c.JSON(http.StatusOK, priceResultJSON)
}

// ParseHTML : TODO
func ParseHTML(c echo.Context) error {
	// return c.Render(http.StatusOK, "banosa.html", data)
	w := c.Response().Writer
	t, err := template.New("").Parse("template/banosa.html")
	if err != nil {
		log.Fatal(err)
	}
	t.ExecuteTemplate(w, "banosa.html", nil)
	return nil
}

func HelloStrcutCTRL(c echo.Context) error {
	dic := newDict()
	dic.data[1] = "A"
	dic.data[2] = "B"

	dictdata := Dict{
		Data: *dic,
	}

	byte, err := json.Marshal(dictdata)
	if err != nil {
		fmt.Println(err)
	}

	return c.String(http.StatusOK, string(byte))
}

func newDict() *datablock {
	d := datablock{}
	d.data = map[int]string{}
	return &d
}

// func New() error {

// 	// var test string , test := "" 는 언뜻 같아 보일 수 있지만, 엄연히 다르다
// 	// var test string은 test를 string 타입으로 선언만 했지, 값을 초기화 하진 않았다.
// 	// test = "" 은 test를 string 타입으로 선언함과 동시에 값을 초기화 했다.
// 	var test string
// 	test = ""

// 	// 구조체의 선언과 초기화도 마찬가지.
// 	// var hosts *Hostzones의 경우, Hostzones의 포인터로 hosts를 선언만 한 것이고,
// 	// hosts1 := new(Hostzones)의 경우, Hostzones의 포인터로 hosts1를 선언함과 동시에 값을 초기화 한 것이다.
// 	var hosts *Hostzones

// 	hosts1 := new(Hostzones)
// 	hosts2 := new(Hostzones)
// 	if true {
// 		hosts.Hostzones = hosts1
// 		}  false {
// 		hosts= hosts2

// 	}

// 	//
// 	for _, v := range respData.Record.Get {
// 		log.Println(v.Active.Types)
// 		if len(v.Active.Types.NS) != 0 {
// 			hostlist := Hostlist{}
// 			hostlist.DomainName = v.Name
// 			recordscnt := CountRecord(v.Name)
// 			hostlist.RecordsCount = recordscnt
// 			hosts.Hostzones = append(hosts.Hostzones, hostlist)
// 		}
// 	}

// 	// 아래 코드에서는 panic이 발생함.
// 	// hostlist를 Hostlist 구조체 타입으로 선언만 했지, 초기화하지 않은 상태에서 if 문을 태워서 그런 것 같음.
// 	// hostlist를 for문 밖에서 선언한것도 NG
// 	var hostlist Hostlist
// 	var hosts *Hostzones

// 	for _, v := range respData.Record.Get {
// 		if len(v.Active.Types.NS) != 0 {
// 			hostlist.DomainName = v.Name
// 			recordscnt := CountRecord(v.Name)
// 			hostlist.RecordsCount = recordscnt
// 			hosts.Hostzones = append(hosts.Hostzones, hostlist)
// 		}
// 	}test

// 	return nil
// }

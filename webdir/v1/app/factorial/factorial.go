package factorial

import (
	"strconv"
)

// GetFacto : Get QueryID from controller and calculate factorial and return string
// 입력값 int로 변경
func GetFacto(QueryID int) string {

	var QueryIDFactorial int = 1

	for i := 1; i <= QueryID; i++ {
		QueryIDFactorial *= i
	}

	QueryIDFactorialResult := strconv.Itoa(QueryIDFactorial)
	return QueryIDFactorialResult
}

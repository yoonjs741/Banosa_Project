package factorial

import (
	"strconv"
)

// GetFacto : Get QueryID from controller and calculate factorial and return it
func GetFacto(QueryID string) string {

	var QueryIDint int
	QueryIDFactorial := 1

	QueryIDint, _ = strconv.Atoi(QueryID)

	for i := 1; i <= QueryIDint; i++ {
		QueryIDFactorial *= i
	}

	QueryIDFactorialResult := strconv.Itoa(QueryIDFactorial)

	return QueryIDFactorialResult
}

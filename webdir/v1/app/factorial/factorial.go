package factorial

// GetFacto : Get QueryID from controller and calculate factorial and return string
// 입력값 int로 변경
// 응답값도 int로 변경
func GetFacto(QueryID int) int {

	var QueryIDFactorial int = 1

	for i := 1; i <= QueryID; i++ {
		QueryIDFactorial *= i
	}

	return QueryIDFactorial
}

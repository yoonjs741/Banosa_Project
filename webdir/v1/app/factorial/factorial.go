package factorial

// GetFacto : Get queryID from controller and calculate factorial and return string
// 입력값 int로 변경
// 응답값도 int로 변경
func GetFacto(queryID int) int {

	var queryIDFactorial int = 1

	for i := 1; i <= queryID; i++ {
		queryIDFactorial *= i
	}

	return queryIDFactorial
}

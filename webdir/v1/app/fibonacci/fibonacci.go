package fibonacci

// GetFacto : Get queryID from controller and calculate factorial and return string
// 입력값 int로 변경
// 응답값도 int로 변경
func GetFibonacci(queryID int) int {

	var queryIDFibonacci int

	if queryIDFibonacci < 2 {
		return queryIDFibonacci
	} else {
		var idx, tmp, fn, fn_1 int
		fn = 1
		for idx = 2; idx <= queryIDFibonacci; idx++ {
			tmp = fn
			fn += fn_1
			fn_1 = tmp
		}
		return fn
	}
}

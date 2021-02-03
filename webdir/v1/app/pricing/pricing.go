package pricing

// GetPrice : Input value and sale rate, Output Price.
func GetPrice(orgValue, discountRate int) int {

	discountRateRemainder := float32(discountRate) / 100
	return int(float32(orgValue) - (float32(orgValue) * float32(discountRateRemainder)))
} // 원가, 할인율 받아서 현재가 반환하는 함수 만들 예정

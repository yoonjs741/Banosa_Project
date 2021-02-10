package pricing

// GetPrice : Input value and sale rate, Output Price.
func GetPrice(orgValue, discountRate int) int {

	discountRateRemainder := float32(discountRate) / 100
	return int(float32(orgValue) - (float32(orgValue) * float32(discountRateRemainder)))
}

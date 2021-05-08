package _890

import (
	"fmt"
	"math"
)

func plate(consonants, digits int) int {
	amountPlates := 0.0
	if consonants == 0 && digits == 0 {
		amountPlates = 0
	} else {
		amountPlates = math.Pow(26, float64(consonants)) * math.Pow(10, float64(digits))
	}
	return int(amountPlates)
}
func main() {
	var instances int
	fmt.Scan(&instances)
	for i := 0; i < instances; i++ {
		var consonants, digits int
		fmt.Scan(&consonants, &digits)
		fmt.Printf("%d\n", plate(consonants, digits))
	}
}
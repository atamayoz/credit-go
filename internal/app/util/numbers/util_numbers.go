package numbers

import (
	"fmt"
	"math"
	"strconv"
)

func RoundToFourDecimals(number float64) (float64, error) {
	roundedStr := fmt.Sprintf("%.4f", number)
	return strconv.ParseFloat(roundedStr, 64)
}

func RoundToTwoDecimal(num float64) float64 {
	return math.Round(num*100) / 100
}

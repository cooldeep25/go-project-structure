package utils

import (
	"math"
	// "github.com/betacraft/num2words"
)

// Round shortens a float32 value to a specified precision (number of digits after the decimal point)
// with "round half up" tie-braking rule. Half-way values (23.5) are always rounded up (24).
func Round(v float64, precision int) float64 {
	shift := math.Pow(10, float64(precision))
	t := float64(v) * shift
	if t > 0 {
		return math.Floor(t+0.5) / shift
	}
	return math.Ceil(t-0.5) / shift
}

func Truncate(v float64, precision int) float64 {
	shift := math.Pow(10, float64(precision))
	return math.Floor(shift*v) / shift
}

// func ConvertToWords(no float64) string {
// 	rs := int(no)
// 	paisa := int(math.Mod(no*100, 100))
// 	result := ""
// 	if rs != 0 {
// 		result = num2words.Convert(rs) + " Rupees "
// 	}

// 	if paisa != 0 {
// 		result = result + num2words.Convert(paisa) + " Paisa"
// 	}

// 	return result
// }

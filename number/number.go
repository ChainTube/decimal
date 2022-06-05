package number

import "math"

// ShiftFloatToInt shifts decimal point to the right (multiply) to store a float as int.
func ShiftFloatToInt(nr float64, decimals int) int64 {
	return int64(math.Round(float64(nr) * math.Pow(10, float64(decimals))))
}

// ShiftIntToFloat shifts decimal places to the left (divide) to get a float value of an int created with ShiftFloatToInt().
func ShiftIntToFloat(nr int64, decimals int) float64 {
	return float64(nr) / math.Pow(10, float64(decimals))
}

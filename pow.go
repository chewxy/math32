package math32

import "math"

func Pow(x, y float32) float32 {
	return float32(math.Pow(float64(x), float64(y)))
}

func Pow10(e int) float32 {
	return float32(math.Pow10(e))
}

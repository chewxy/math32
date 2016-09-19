package math32

import "math"

func Remainder(x, y float32) float32 {
	return float32(math.Remainder(float64(x), float64(y)))
}

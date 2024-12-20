package math32

import "math"

// Erfc returns the complementary error function of x.
//
// Special cases are:
//
//	Erfc(+Inf) = 0
//	Erfc(-Inf) = 2
//	Erfc(NaN) = NaN
func Erfc(x float32) float32 {
	return float32(math.Erfc(float64(x)))
}

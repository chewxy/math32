package math32

import "math"

// Erf returns the error function of x.
//
// Special cases are:
//
//	Erf(+Inf) = 1
//	Erf(-Inf) = -1
//	Erf(NaN) = NaN
func Erf(x float32) float32 {
	return float32(math.Erf(float64(x)))
}

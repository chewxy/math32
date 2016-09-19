package math32

import "math"

func Tanh(x float32) float32 {
	return float32(math.Tanh(float64(x)))
}

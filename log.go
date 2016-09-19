package math32

import "math"

func Log(x float32) float32 {
	return float32(math.Log(float64(x)))
}

package math32

import "math"

func Round(value float32) float32 {
	return float32(math.Round(float64(value)))
}

package math32

func Log10(x float32) float32 {
	return Log(x) * (1 / Ln10)
}

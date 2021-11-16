package math32

func Log2(x float32) float32 {
	return log2(x)
}

func log2(x float32) float32 {
	frac, exp := Frexp(x)
	// Make sure exact powers of two give an exact answer.
	// Don't depend on Log(0.5)*(1/Ln2)+exp being exactly exp-1.
	if frac == 0.5 {
		return float32(exp - 1)
	}
	return Log(frac)*(1/Ln2) + float32(exp)
}

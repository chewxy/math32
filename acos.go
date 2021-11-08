package math32

func Acos(x float32) float32 {
	const (
		pio2_hi = Pi / 2
		pio2_lo = -4.37113900018624283e-8

		pS0 = 1.6666667163e-01  // 0x3e2aaaab
		pS1 = -3.2556581497e-01 // 0xbea6b090
		pS2 = 2.0121252537e-01  // 0x3e4e0aa8
		pS3 = -4.0055535734e-02 // 0xbd241146
		pS4 = 7.9153501429e-04  // 0x3a4f7f04
		pS5 = 3.4793309169e-05  // 0x3811ef08
		qS1 = -2.4033949375e+00 // 0xc019d139
		qS2 = 2.0209457874e+00  // 0x4001572d
		qS3 = -6.8828397989e-01 // 0xbf303361
		qS4 = 7.7038154006e-02  // 0x3d9dc62e
	)
	var z, p, q, r, w, s, c, df float32
	hx := float32ibits(x)
	ix := hx & 0x7fffffff

	if ix == 0x3f800000 { // |x|==1
		if hx > 0 {
			return 0.0 // acos(1) = 0
		} else {
			return Pi + 2*pio2_lo // acos(-1)= pi
		}
	} else if ix > 0x3f800000 { //|x| >= 1
		return (x - x) / (x - x) // acos(|x|>1) is NaN
	}

	if ix < 0x3f000000 { // |x| < 0.5
		if ix <= 0x32800000 {
			return pio2_hi + pio2_lo //if|x|<=2**-26
		}
		z = x * x
		p = z * (pS0 + z*(pS1+z*(pS2+z*(pS3+z*(pS4+z*pS5)))))
		q = 1 + z*(qS1+z*(qS2+z*(qS3+z*qS4)))
		r = p / q
		return pio2_hi - (x - (pio2_lo - x*r))
	} else if hx < 0 { // x < -0.5
		z = (1 + x) * 0.5
		p = z * (pS0 + z*(pS1+z*(pS2+z*(pS3+z*(pS4+z*pS5)))))
		q = 1 + z*(qS1+z*(qS2+z*(qS3+z*qS4)))
		s = Sqrt(z)
		r = p / q
		w = r*s - pio2_lo
		return Pi - 2*(s+w)
	} else { // x > 0.5
		z = (1 - x) * 0.5
		s = Sqrt(z)
		df = s
		idf := float32ibits(df)
		msk := 0xfffff000
		df = float32fromibits(idf & int32(msk))
		c = (z - df*df) / (s + df)
		p = z * (pS0 + z*(pS1+z*(pS2+z*(pS3+z*(pS4+z*pS5)))))
		q = 1 + z*(qS1+z*(qS2+z*(qS3+z*qS4)))
		r = p / q
		w = r*s + c
		return 2 * (df + w)
	}
}

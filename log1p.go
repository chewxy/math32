package math32

import "math"

func Log1p(x float32) float32 {
	return float32(math.Log1p(float64(x)))
}

func log1p(x float32) float32 {
	const (
		Sqrt2M1     = 4.142135623730950488017e-01  // Sqrt(2)-1 = 0x3fda827999fcef34
		Sqrt2HalfM1 = -2.928932188134524755992e-01 // Sqrt(2)/2-1 = 0xbfd2bec333018866
		Small       = 1.0 / (1 << (12))            // 2**-29 = 0x3e20000000000000
		Tiny        = 1.0 / (1 << (shift + 1))     // 2**-54
		Two53       = 1 << (shift + 1)             // 2**53
		Ln2Hi       = 6.93147180369123816490e-01   // 3fe62e42fee00000
		Ln2Lo       = 1.90821492927058770002e-10   // 3dea39ef35793c76
		Lp1         = 6.666666666666735130e-01     // 3FE5555555555593
		Lp2         = 3.999999999940941908e-01     // 3FD999999997FA04
		Lp3         = 2.857142874366239149e-01     // 3FD2492494229359
		Lp4         = 2.222219843214978396e-01     // 3FCC71C51D8E78AF
		Lp5         = 1.818357216161805012e-01     // 3FC7466496CB03DE
		Lp6         = 1.531383769920937332e-01     // 3FC39A09D078C69F
		Lp7         = 1.479819860511658591e-01     // 3FC2F112DF3E5244
		divbit      = 1 << shift
		uvone       = 0x3f800000
	)

	// special cases
	switch {
	case x < -1 || IsNaN(x): // includes -Inf
		return NaN()
	case x == -1:
		return Inf(-1)
	case IsInf(x, 1):
		return Inf(1)
	}

	absx := Abs(x)

	var f float32
	var iu uint32
	k := 1
	if absx < Sqrt2M1 { //  |x| < Sqrt(2)-1
		if absx < Small { // |x| < 2**-29
			if absx < Tiny { // |x| < 2**-54
				return x
			}
			return x - x*x*0.5
		}
		if x > Sqrt2HalfM1 { // Sqrt(2)/2-1 < x
			// (Sqrt(2)/2-1) < x < (Sqrt(2)-1)
			k = 0
			f = x
			iu = 1
		}
	}
	var c float32
	if k != 0 {
		var u float32
		if absx < Two53 { // 1<<53
			u = 1.0 + x
			iu = Float32bits(u)
			k = int((iu >> shift) - bias)
			// correction term
			if k > 0 {
				c = 1.0 - (u - x)
			} else {
				c = x - (u - 1.0)
			}
			c /= u
		} else {
			u = x
			iu = Float32bits(u)
			k = int((iu >> shift) - bias)
			c = 0
		}
		iu &= notmask

		if iu < Float32bits(math.Sqrt2)&^mask { // mantissa of Sqrt(2)
			u = Float32frombits(iu | uvone) // normalize u
		} else {
			k++
			u = Float32frombits(iu | uvone&^divbit) // normalize u/2
			iu = (divbit - iu) >> 2
		}
		f = u - 1.0 // Sqrt(2)/2 < u < Sqrt(2)
	}
	hfsq := 0.5 * f * f
	var s, R, z float32
	if iu == 0 { // |f| < 2**-20
		if f == 0 {
			if k == 0 {
				return 0
			}
			c += float32(k) * Ln2Lo
			return float32(k)*Ln2Hi + c
		}
		R = hfsq * (1.0 - 0.66666666666666666*f) // avoid division
		if k == 0 {
			return f - R
		}
		return float32(k)*Ln2Hi - ((R - (float32(k)*Ln2Lo + c)) - f)
	}
	s = f / (2.0 + f)
	z = s * s
	R = z * (Lp1 + z*(Lp2+z*(Lp3+z*(Lp4+z*(Lp5+z*(Lp6+z*Lp7))))))
	if k == 0 {
		return f - (hfsq - s*(hfsq+R))
	}
	return float32(k)*Ln2Hi - ((hfsq - (s*(hfsq+R) + (float32(k)*Ln2Lo + c))) - f)
}

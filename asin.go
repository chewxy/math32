package math32

func Asin(x float32) float32 {
	const (
		huge    = 1e30
		pio2_hi = Pi / 2
		pio2_lo = -4.37113900018624283e-8
		pio4_hi = Pi / 4
		p0      = 1.666675248e-1
		p1      = 7.495297643e-2
		p2      = 4.547037598e-2
		p3      = 2.417951451e-2
		p4      = 4.216630880e-2
	)

	hx := float32ibits(x)
	ix := hx & 0x7fffffff
	if ix == 0x3f800000 {
		// asin(1)=+-pi/2 with inexact
		return x*pio2_hi + x*pio2_lo
	} else if ix > 0x3f800000 { // |x|>= 1
		// asin(|x|>1) is NaN
		return (x - x) / (x - x)
	} else if ix < 0x3f000000 { //|x|<0.5
		if ix < 0x32000000 { // if |x| < 2**-27
			// math_check_force_underflow(x)
			if huge+x > 1 {
				// return x with inexact if x!=0
				return x
			}
		} else {
			t := x * x
			w := t * (p0 + t*(p1+t*(p2+t*(p3+t*p4))))
			return x + x*w
		}
	}

	// 1> |x|>= 0.5
	w := 1 - Abs(x)
	t := w * 0.5
	p := t * (p0 + t*(p1+t*(p2+t*(p3+t*p4))))
	s := Sqrt(t)
	if ix >= 0x3F79999A { /* if |x| > 0.975 */
		t = pio2_hi - (2.0*(s+s*p) - pio2_lo)
	} else {
		w = s
		iw := float32ibits(w)
		msk := 0xfffff000
		w = float32fromibits(iw & int32(msk))
		c := (t - w*w) / (s + w)
		r := p
		p = 2*s*r - (pio2_lo - 2*c)
		q := pio4_hi - 2*w
		t = pio4_hi - (p - q)
	}
	if hx > 0 {
		return t
	}
	return -t
}

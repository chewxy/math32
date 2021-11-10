package math32

func Sinh(x float32) float32 {
	// The coefficients are #2029 from Hart & Cheney. (20.36D)
	const (
		P0 = -0.6307673640497716991184787251e+6
		P1 = -0.8991272022039509355398013511e+5
		P2 = -0.2894211355989563807284660366e+4
		P3 = -0.2630563213397497062819489e+2
		Q0 = -0.6307673640497716991212077277e+6
		Q1 = 0.1521517378790019070696485176e+5
		Q2 = -0.173678953558233699533450911e+3
	)

	sign := false
	if x < 0 {
		x = -x
		sign = true
	}

	var temp float32
	switch {
	case x > 21:
		temp = Exp(x) * 0.5

	case x > 0.5:
		ex := Exp(x)
		temp = (ex - 1/ex) * 0.5

	default:
		sq := x * x
		temp = (((P3*sq+P2)*sq+P1)*sq + P0) * x
		temp = temp / (((sq+Q2)*sq+Q1)*sq + Q0)
	}

	if sign {
		temp = -temp
	}
	return temp
}

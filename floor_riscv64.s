//go:build !tinygo && !noasm

#include "textflag.h"

// The base RISC-V F extension has no round-to-integral instruction, so these
// round through an integer: FCVTWS (which Go always assembles with the RTZ
// rounding mode, i.e. truncation) followed by FCVTSW back to float32.
//
// Inputs with |x| >= 2^23, ±Inf and NaN are already integral and are returned
// unchanged. Every other input fits in an int32, so the conversion is exact.
// Floor and Ceil then adjust the truncated value by one when it lies on the
// wrong side of x. Finally FSGNJS copies the sign of x onto the result, so that
// results that round to zero keep the sign of x, e.g. Ceil(-0.5) = -0. No
// nonzero result can have a sign different from x, so this is always correct.

// func archFloor(x float32) float32
TEXT ·archFloor(SB),NOSPLIT,$0-12
	MOVF	x+0(FP), F0
	FMVXW	F0, X10
	SRL	$23, X10, X11
	AND	$0xFF, X11, X11	// X11 = biased exponent of x
	MOV	$149, X12	// 149 = 127 + 22
	BLTU	X12, X11, done	// |x| >= 2^23, ±Inf or NaN
	FCVTWS	F0, X10		// X10 = trunc(x)
	FCVTSW	X10, F1
	FLTS	F1, F0, X11	// X11 = 1 if x < trunc(x)
	SUB	X11, X10, X10
	FCVTSW	X10, F1
	FSGNJS	F0, F1, F0	// F0 = |F1| with the sign of x
done:
	MOVF	F0, ret+8(FP)
	RET

// func archCeil(x float32) float32
TEXT ·archCeil(SB),NOSPLIT,$0-12
	MOVF	x+0(FP), F0
	FMVXW	F0, X10
	SRL	$23, X10, X11
	AND	$0xFF, X11, X11	// X11 = biased exponent of x
	MOV	$149, X12	// 149 = 127 + 22
	BLTU	X12, X11, done	// |x| >= 2^23, ±Inf or NaN
	FCVTWS	F0, X10		// X10 = trunc(x)
	FCVTSW	X10, F1
	FLTS	F0, F1, X11	// X11 = 1 if trunc(x) < x
	ADD	X11, X10, X10
	FCVTSW	X10, F1
	FSGNJS	F0, F1, F0	// F0 = |F1| with the sign of x
done:
	MOVF	F0, ret+8(FP)
	RET

// func archTrunc(x float32) float32
TEXT ·archTrunc(SB),NOSPLIT,$0-12
	MOVF	x+0(FP), F0
	FMVXW	F0, X10
	SRL	$23, X10, X11
	AND	$0xFF, X11, X11	// X11 = biased exponent of x
	MOV	$149, X12	// 149 = 127 + 22
	BLTU	X12, X11, done	// |x| >= 2^23, ±Inf or NaN
	FCVTWS	F0, X10		// X10 = trunc(x)
	FCVTSW	X10, F1
	FSGNJS	F0, F1, F0	// F0 = |F1| with the sign of x
done:
	MOVF	F0, ret+8(FP)
	RET

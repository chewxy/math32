//go:build !tinygo && !noasm

#include "textflag.h"

// FMOVS widens the float32 argument to float64 exactly, and the rounded result
// is an integer representable as a float32, so storing it back is exact too.

// func archFloor(x float32) float32
TEXT ·archFloor(SB),NOSPLIT,$0-12
	FMOVS	x+0(FP), F0
	FRIM	F0, F0
	FMOVS	F0, ret+8(FP)
	RET

// func archCeil(x float32) float32
TEXT ·archCeil(SB),NOSPLIT,$0-12
	FMOVS	x+0(FP), F0
	FRIP	F0, F0
	FMOVS	F0, ret+8(FP)
	RET

// func archTrunc(x float32) float32
TEXT ·archTrunc(SB),NOSPLIT,$0-12
	FMOVS	x+0(FP), F0
	FRIZ	F0, F0
	FMOVS	F0, ret+8(FP)
	RET

//go:build !tinygo && !noasm

#include "textflag.h"

// func archSqrt(x float32) float32
TEXT ·archSqrt(SB),NOSPLIT,$0-12
	FMOVS	x+0(FP), F0
	FSQRTS	F0, F0
	FMOVS	F0, ret+8(FP)
	RET

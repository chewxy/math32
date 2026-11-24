//go:build !tinygo && !noasm

#include "textflag.h"

// func archSqrt(x float32) float32
TEXT ·archSqrt(SB),NOSPLIT,$0-12
	MOVF	x+0(FP), F0
	SQRTF	F0, F0
	MOVF	F0, ret+8(FP)
	RET

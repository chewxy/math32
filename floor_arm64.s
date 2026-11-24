//go:build !tinygo && !noasm

#include "textflag.h"

// func archFloor(x float32) float32
TEXT ·archFloor(SB),NOSPLIT,$0-12
	FMOVS	x+0(FP), F0
	FRINTMS	F0, F0
	FMOVS	F0, ret+8(FP)
	RET

// func archCeil(x float32) float32
TEXT ·archCeil(SB),NOSPLIT,$0-12
	FMOVS	x+0(FP), F0
	FRINTPS	F0, F0
	FMOVS	F0, ret+8(FP)
	RET

// func archTrunc(x float32) float32
TEXT ·archTrunc(SB),NOSPLIT,$0-12
	FMOVS	x+0(FP), F0
	FRINTZS	F0, F0
	FMOVS	F0, ret+8(FP)
	RET

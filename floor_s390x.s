//go:build !tinygo && !noasm

#include "textflag.h"

// FIEBR rounds to an integral value using the rounding mode given by its
// first operand: 5 is toward zero, 6 is toward +Inf and 7 is toward -Inf.

// func archFloor(x float32) float32
TEXT ·archFloor(SB),NOSPLIT,$0-12
	FMOVS	x+0(FP), F0
	FIEBR	$7, F0, F0
	FMOVS	F0, ret+8(FP)
	RET

// func archCeil(x float32) float32
TEXT ·archCeil(SB),NOSPLIT,$0-12
	FMOVS	x+0(FP), F0
	FIEBR	$6, F0, F0
	FMOVS	F0, ret+8(FP)
	RET

// func archTrunc(x float32) float32
TEXT ·archTrunc(SB),NOSPLIT,$0-12
	FMOVS	x+0(FP), F0
	FIEBR	$5, F0, F0
	FMOVS	F0, ret+8(FP)
	RET

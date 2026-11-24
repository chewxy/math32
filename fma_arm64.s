//go:build !tinygo && !noasm

#include "textflag.h"

// func archFMA(x, y, z float32) float32
TEXT ·archFMA(SB),NOSPLIT,$0-20
	FMOVS	x+0(FP), F0
	FMOVS	y+4(FP), F1
	FMOVS	z+8(FP), F2
	FMADDS	F0, F2, F1, F0	// F0 = F2 + F1*F0
	FMOVS	F0, ret+16(FP)
	RET

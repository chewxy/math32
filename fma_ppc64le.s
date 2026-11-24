//go:build !tinygo && !noasm

#include "textflag.h"

// func archFMA(x, y, z float32) float32
TEXT ·archFMA(SB),NOSPLIT,$0-20
	FMOVS	x+0(FP), F1
	FMOVS	y+4(FP), F2
	FMOVS	z+8(FP), F3
	FMADDS	F1, F3, F2, F1	// F1 = F1*F2 + F3
	FMOVS	F1, ret+16(FP)
	RET

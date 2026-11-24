//go:build !tinygo && !noasm

#include "textflag.h"

// func archFMA(x, y, z float32) float32
TEXT ·archFMA(SB),NOSPLIT,$0-20
	MOVF	x+0(FP), F0
	MOVF	y+4(FP), F1
	MOVF	z+8(FP), F2
	FMADDS	F0, F1, F2, F0	// F0 = F0*F1 + F2
	MOVF	F0, ret+16(FP)
	RET

//go:build !tinygo && !noasm && go1.24

#include "textflag.h"

// The Go assembler supports FMADDF on loong64 since Go 1.24.

// func archFMA(x, y, z float32) float32
TEXT ·archFMA(SB),NOSPLIT,$0-20
	MOVF	x+0(FP), F0
	MOVF	y+4(FP), F1
	MOVF	z+8(FP), F2
	FMADDF	F2, F1, F0, F0	// F0 = F0*F1 + F2
	MOVF	F0, ret+16(FP)
	RET

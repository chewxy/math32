// +build wasm

#include "textflag.h"

// func archExp(x float32) float32
TEXT ·archExp(SB),NOSPLIT,$0
	JMP ·exp(SB)

// func archExp2(x float32) float32
TEXT ·archExp2(SB),NOSPLIT,$0
	JMP ·exp2(SB)

// func Log(x float32) float32
TEXT ·Log(SB),NOSPLIT,$0
	JMP ·log(SB)

// func Remainder(x, y float32) float32
TEXT ·Remainder(SB),NOSPLIT,$0
	JMP ·remainder(SB)

// func archSqrt(x float32) float32
TEXT ·archSqrt(SB),NOSPLIT,$0
	JMP ·sqrt(SB)

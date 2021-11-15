// +build wasm

#include "textflag.h"

// func Exp(x float32) float32
TEXT ·Exp(SB),NOSPLIT,$0
	JMP ·exp(SB)

// func Exp2(x float32) float32
TEXT ·Exp2(SB),NOSPLIT,$0
	JMP ·exp2(SB)

// func Log(x float32) float32
TEXT ·Log(SB),NOSPLIT,$0
	JMP ·log(SB)

// func archRemainder(x, y float32) float32
TEXT ·archRemainder(SB),NOSPLIT,$0
	JMP ·remainder(SB)

// func Sqrt(x float32) float32
TEXT ·Sqrt(SB),NOSPLIT,$0
	JMP ·sqrt(SB)

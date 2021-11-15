#include "textflag.h"

// func Exp(x float32) float32
TEXT ·Exp(SB),NOSPLIT,$0
	BR ·exp(SB)

// func Log(x float32) float32
TEXT ·Log(SB),NOSPLIT,$0
	BR ·log(SB)

// func archRemainder(x, y float32) float32
TEXT ·archRemainder(SB),NOSPLIT,$0
	BR ·remainder(SB)

// func Sqrt(x float32) float32
TEXT ·Sqrt(SB),NOSPLIT,$0
	BR ·sqrt(SB)

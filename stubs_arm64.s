#include "textflag.h"

// func archLog(x float64) float64
TEXT ·Log(SB),NOSPLIT,$0
	B ·log(SB)

TEXT ·archRemainder(SB),NOSPLIT,$0
	B ·remainder(SB)
